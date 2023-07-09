package intake

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	pb "github.com/ssargent/apis/pkg/bbq/intake/v1"
	"github.com/ssargent/bbq/cmd/bbq/internal/config"
	"github.com/ssargent/bbq/internal/services"
)

type intakeServer struct {
	pb.UnsafeIntakeServiceServer

	cfg    *config.Config
	cache  *cache.Cache
	db     *sqlx.DB
	logger *zap.Logger

	intake *services.IntakeService
}

// todo(scott): this should take in the service directly, not construct it.
func RegisterIntake(server *grpc.Server, cfg *config.Config, cache *cache.Cache, db *sqlx.DB, logger *zap.Logger) {
	pb.RegisterIntakeServiceServer(server, newIntakeServer(cfg, cache, db, logger))
}

func newIntakeServer(cfg *config.Config, cache *cache.Cache, db *sqlx.DB, logger *zap.Logger) *intakeServer {
	intake := services.NewIntakeService(cache, db, logger)
	return &intakeServer{
		intake: intake,
		cfg:    cfg,
		cache:  cache,
		db:     db,
		logger: logger,
	}
}

func (s *intakeServer) Record(ctx context.Context, in *pb.RecordRequest) (*pb.RecordResponse, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("proto.Marshal: %w", err)
	}

	jsonData, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	fmt.Printf("size of proto: %d, size of json: %d\nproto data := %v\n json version := %s\n", len(data), len(jsonData), data, string(jsonData))

	return &pb.RecordResponse{
		SessionId: "test123",
	}, nil
}

func (s *intakeServer) Session(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	sessionID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("NewUUID: %w", err)
	}

	// s.intake.CreateSession(ctx)

	resp := pb.SessionResponse{
		Session: &pb.Session{
			Id:   sessionID.String(),
			Name: in.Name,
			DataRate: &pb.SessionDataRate{
				Sensors:           8,
				MaxReadingsMinute: 120,
			},
		},
	}

	return &resp, nil
}
