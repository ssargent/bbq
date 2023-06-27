package intake

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/google/uuid"
	pb "github.com/ssargent/apis/pkg/bbq/intake/v1"
)

type intakeServer struct {
	pb.UnsafeIntakeServiceServer
}

func RegisterIntake(server *grpc.Server) {
	pb.RegisterIntakeServiceServer(server, newIntakeServer())
}

func newIntakeServer() *intakeServer {
	return &intakeServer{}
}

func (s *intakeServer) Record(ctx context.Context, in *pb.RecordRequest) (*pb.RecordResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *intakeServer) Session(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	sessionID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("NewUUID: %w", err)
	}

	dt := time.Now()

	resp := pb.SessionResponse{
		Session: &pb.Session{
			Id:   sessionID.String(),
			Name: fmt.Sprintf("bbq%d-%d-%d", dt.Year(), int(dt.Month()), dt.Day()),
			DataRate: &pb.SessionDataRate{
				Sensors:           8,
				MaxReadingsMinute: 120,
			},
		},
	}

	return &resp, nil
}
