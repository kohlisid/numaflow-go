package flatmapper

import (
	"context"
	"time"
)

// Datum contains methods to get the payload information.
type Datum interface {
	// Value returns the payload of the message.
	Value() []byte
	// EventTime returns the event time of the message.
	EventTime() time.Time
	// Watermark returns the watermark of the message.
	Watermark() time.Time
	// Headers returns the headers of the message.
	Headers() map[string]string

	Uuid() string
}

// FlatMapper is the interface of map function implementation.
type FlatMapper interface {
	// Map is the function to process each coming message.
	Map(ctx context.Context, keys []string, datum Datum) Messages
}

// MapperFunc is a utility type used to convert a map function to a Mapper.
type MapperFunc func(ctx context.Context, keys []string, datum Datum) Messages

// Map implements the function of map function.
func (mf MapperFunc) Map(ctx context.Context, keys []string, datum Datum) Messages {
	return mf(ctx, keys, datum)
}

// FlatmapCreator is the interface which is used to create a FlatMapper.
type FlatmapCreator interface {
	// Create creates a FlatMapper, will be invoked by each
	Create() FlatMapper
}

// simpleFlatmapCreator is an implementation of FlatmapCreator, which creates a FlatM for the given function.
type simpleFlatmapCreator struct {
	f func(ctx context.Context, keys []string, datum Datum) Messages
}

// Create creates a Reducer for the given function.
func (s *simpleFlatmapCreator) Create() FlatMapper {
	return MapperFunc(s.f)
}

// SimpleCreatorWithFlatmapFn creates a simple FlatMapCreator for the given reduce function.
func SimpleCreatorWithFlatmapFn(f func(ctx context.Context, keys []string, datum Datum) Messages) FlatmapCreator {
	return &simpleFlatmapCreator{f: f}
}
