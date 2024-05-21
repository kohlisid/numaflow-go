package main

import (
	"context"
	"log"
	"strings"

	"github.com/numaproj/numaflow-go/pkg/flatmapper"
)

func mapFn(_ context.Context, keys []string, d flatmapper.Datum) flatmapper.Messages {
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	// Split the msg into an array with comma.
	strs := strings.Split(string(msg), ",")
	results := flatmapper.MessagesBuilder()
	for _, s := range strs {
		results = results.Append(flatmapper.NewMessage([]byte(s), d.Uuid()))
	}
	return results
}

func main() {
	err := flatmapper.NewServer(flatmapper.MapperFunc(mapFn)).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ", err)
	}
}
