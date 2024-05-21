package flatmapper

import (
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	flatmappb "github.com/numaproj/numaflow-go/pkg/apis/proto/flatmap/v1"
)

const (
	uds                   = "unix"
	defaultMaxMessageSize = 1024 * 1024 * 64
	address               = "/var/run/numaflow/flatmap.sock"
	serverInfoFilePath    = "/var/run/numaflow/flatmap-server-info"
)

// Service implements the proto gen server interface and contains the map
// streaming function.
type Service struct {
	flatmappb.UnimplementedFlatmapServer

	FlatMapper FlatMapper
}

// IsReady returns true to indicate the gRPC connection is ready.
func (fs *Service) IsReady(context.Context, *emptypb.Empty) (*flatmappb.ReadyResponse, error) {
	return &flatmappb.ReadyResponse{Ready: true}, nil
}

func (fs *Service) MapFn(stream flatmappb.Flatmap_MapFnServer) error {
	var (
		//ctx           = stream.Context()
		//g             errgroup.Group
		datumStreamCh = make(chan Datum)
	)

	// Make call to kick off stream handling
	messageCh := make(chan Message)
	done := make(chan bool)
	go func() {
		for msg := range datumStreamCh {
			log.Print("MYDEBUG: Got a message on the stream ", msg.Uuid())
			messageCh <- Message{
				value: msg.Value(),
				tags:  nil,
				uuid:  msg.Uuid(),
			}
		}

	}()

	// Read messages and push to read channel
	go func() {
		for {
			d, err := stream.Recv()
			log.Print("MYDEBUG: I'm here", d.GetUuid())
			if err == io.EOF {
				close(datumStreamCh)
				return
			}
			if err != nil {
				close(datumStreamCh)
				// TODO: research on gRPC errors and revisit the error handler
				return
			}
			var hd = &handlerDatum{
				value:     d.GetValue(),
				eventTime: d.GetEventTime().AsTime(),
				watermark: d.GetWatermark().AsTime(),
				headers:   d.GetHeaders(),
				uuid:      d.GetUuid(),
			}
			log.Print("MYDEBUG: sending to stream", d.GetUuid())
			datumStreamCh <- hd
		}
	}()

	// Now listen to collect messages
	finished := false
	for {
		select {
		case <-done:
			finished = true
		case message, ok := <-messageCh:
			log.Print("MYDEBUG: a message on the stream ", message.Uuid())
			if !ok {
				// Channel already closed, not closing again.
				//return nil
			}
			element := &flatmappb.MapResponse{
				Result: &flatmappb.MapResponse_Result{
					Keys:  message.Keys(),
					Value: message.Value(),
					Tags:  message.Tags(),
					Uuid:  message.Uuid(),
				},
			}
			err := stream.Send(element)
			// the error here is returned by stream.Send() which is already a gRPC error
			if err != nil {
				// Channel may or may not be closed, as we are not sure leave it to GC.
				return err
			}
		default:
			if finished {
				close(messageCh)
				return nil
			}
		}
	}
	//taskManager := newReduceTaskManager(fs.reducerCreatorHandle)

}
