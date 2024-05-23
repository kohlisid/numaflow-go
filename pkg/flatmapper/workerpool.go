package flatmapper

import (
	"context"

	flatmappb "github.com/numaproj/numaflow-go/pkg/apis/proto/flatmap/v1"
)

type WorkRequest struct {
	keys  []string
	datum Datum
	uid   string
}

type WorkResponse struct {
}
type PoolManager struct {
	jobQueue    chan *WorkRequest
	maxSize     int
	resultQueue chan *flatmappb.MapResponse
	flatMapper  FlatMapper
	tracker     map[int]chan struct{}
}

func (pm *PoolManager) NewWorker(ctx context.Context) chan struct{} {
	doneChan := make(chan struct{})

	go func() {
		defer close(doneChan)
		for msg := range pm.jobQueue {
			mapResp := pm.flatMapper.Map(ctx, msg.keys, msg.datum)
			uid := msg.uid
			workResp := pm.parseMapResponse(mapResp, uid)
			for _, resp := range workResp {
				pm.resultQueue <- resp
			}
		}
	}()

	return doneChan
}

func NewPoolManager(maxWorkers int, flatMapper FlatMapper) *PoolManager {
	// Buffered channel on maxWorkers
	jobQueue := make(chan *WorkRequest, maxWorkers)
	resultQueue := make(chan *flatmappb.MapResponse, maxWorkers)
	return &PoolManager{
		jobQueue:    jobQueue,
		maxSize:     maxWorkers,
		resultQueue: resultQueue,
		flatMapper:  flatMapper,
		tracker:     make(map[int]chan struct{}),
	}
}

func (pm *PoolManager) DeployPool(ctx context.Context) {
	for i := 0; i < pm.maxSize; i++ {
		doneChan := pm.NewWorker(ctx)
		pm.tracker[i] = doneChan
	}
}

func (pm *PoolManager) parseMapResponse(msgs Messages, uid string) []*flatmappb.MapResponse {
	var elements []*flatmappb.MapResponse
	for _, m := range msgs.Items() {
		elements = append(elements, &flatmappb.MapResponse{
			Result: &flatmappb.MapResponse_Result{
				Keys:  m.Keys(),
				Value: m.Value(),
				Tags:  m.Tags(),
				EOR:   false,
				Uuid:  uid,
			},
		})
	}

	// Append the EOR to indicate that the processing for the given request has completed
	elements = append(elements, &flatmappb.MapResponse{
		Result: &flatmappb.MapResponse_Result{
			EOR:  true,
			Uuid: uid,
		},
	})
	return elements
}

func (pm *PoolManager) WaitAll() {
	for _, doneCh := range pm.tracker {
		<-doneCh
	}
}

func (pm *PoolManager) CloseAll() {
	close(pm.jobQueue)
	// TODO(stream): should we give a way to explicitly kill all goroutines as well
}
