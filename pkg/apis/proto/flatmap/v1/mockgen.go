package v1

//go:generate mockgen -destination flatmapmock/flatmapmock.go -package flatmapmock github.com/numaproj/numaflow-go/pkg/apis/proto/flatmap/v1 FlatmapClient,Flatmap_MapFnClient
