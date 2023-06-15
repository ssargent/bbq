package collector

import (
	"context"
	"net/http"

	collectorv1 "github.com/ssargent/apis/pkg/bbq/collector/v1"
	"github.com/ssargent/apis/pkg/bbq/collector/v1/collectorv1connect"

	"github.com/bufbuild/connect-go"
)

func RunCollector() error {
	mux := http.NewServeMux()
	path, handler := collectorv1connect.NewCollectorServiceHandler(&collectorServiceServer{})
	mux.Handle(path, handler)

	return nil
}

type collectorServiceServer struct {
	collectorv1connect.UnimplementedCollectorServiceHandler
}

func (c *collectorServiceServer) Record(
	ctx context.Context,
	req *connect.Request[collectorv1.RecordRequest],
) (*connect.Response[collectorv1.RecordResponse], error) {
	panic("not implemented")
}
