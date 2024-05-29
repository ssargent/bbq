package collector

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	collectorv1 "github.com/ssargent/apis/pkg/bbq/collector/v1"
	"github.com/ssargent/apis/pkg/bbq/collector/v1/collectorv1connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CollectorConfig struct {
}

func NewCollectorServiceHandler() (string, http.Handler, error) {
	path, handler := collectorv1connect.NewCollectorServiceHandler(&collectorServiceServer{})

	return path, handler, nil
}

type collectorServiceServer struct {
	collectorv1connect.UnimplementedCollectorServiceHandler
}

func (c *collectorServiceServer) Record(
	ctx context.Context,
	req *connect.Request[collectorv1.RecordRequest],
) (*connect.Response[collectorv1.RecordResponse], error) {
	//	sessionID := uuid.New()
	res := connect.NewResponse[collectorv1.RecordResponse](
		&collectorv1.RecordResponse{
			SessionId:  "session-1",
			RecordedAt: timestamppb.Now(),
		},
	)

	return res, nil
}

func (c *collectorServiceServer) Session(ctx context.Context, req *connect.Request[collectorv1.SessionRequest]) (*connect.Response[collectorv1.SessionResponse], error) {
	sessionID := uuid.New()
	res := connect.NewResponse[collectorv1.SessionResponse](
		&collectorv1.SessionResponse{
			Session: &collectorv1.Session{
				Id:   sessionID.String(),
				Name: req.Msg.Name,
			},
		},
	)

	return res, nil
}
