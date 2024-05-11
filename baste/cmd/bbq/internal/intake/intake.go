package intake

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	pb "github.com/ssargent/apis/pkg/bbq/intake/v1"
	"github.com/ssargent/bbq/cmd/bbq/internal/config"
	"github.com/ssargent/bbq/internal/repository"
	"github.com/ssargent/bbq/internal/services"
	"github.com/ssargent/bbq/pkg/bbq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	intake := services.NewIntakeService(cache, logger, db, &repository.Queries{})
	return &intakeServer{
		intake: intake,
		cfg:    cfg,
		cache:  cache,
		db:     db,
		logger: logger,
	}
}

func (s *intakeServer) Record(ctx context.Context, in *pb.RecordRequest) (*pb.RecordResponse, error) {
	var sensorReadingCount int
	var actualSensorReadingCount int

	// Getting the full count of readings helps us to not have to reallocate the array often.
	for _, r := range in.Reading {
		sensorReadingCount += len(r.Readings)
	}

	intakeSessionID := uuid.Nil
	if len(in.Reading) > 0 {
		sessionID, err := uuid.Parse(in.Reading[0].SessionId)
		if err != nil {
			return nil, fmt.Errorf("uuid.Parse(SessionId): %w", err)
		}
		intakeSessionID = sessionID
	}

	readings := make([]*bbq.SensorReading, sensorReadingCount)

	for _, rding := range in.Reading {
		sessionID, err := uuid.Parse(rding.SessionId)
		if err != nil {
			return nil, fmt.Errorf("uuid.Parse(SessionId): %w", err)
		}

		for _, r := range rding.Readings {
			reading := bbq.SensorReading{}
			reading.SessionID = sessionID
			reading.ProbeNumber = r.SensorNumber
			reading.Temperature = float64(r.Temperature)
			reading.ReadingOccurred = rding.RecordedAt

			readings[actualSensorReadingCount] = &reading
			actualSensorReadingCount++
		}
	}

	if actualSensorReadingCount != sensorReadingCount {
		s.logger.Warn("actual sensor reading count does not match expected",
			zap.Int("actualSensorReadingCount", actualSensorReadingCount),
			zap.Int("expectedSensorReadingCount", sensorReadingCount))
	}

	if err := s.intake.CreateReadings(ctx, readings); err != nil {
		return nil, fmt.Errorf("CreateReadings: %w", err)
	}

	// need to redo this.
	return &pb.RecordResponse{
		SessionId:  intakeSessionID.String(),
		RecordedAt: timestamppb.New(time.Now()),
	}, nil
}

func (s *intakeServer) Session(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	var session bbq.Session

	session.Description = in.Description
	session.DeviceName = in.DeviceName
	session.SensorName = in.SensorName

	if in.DesiredState == nil {
		session.DesiredState = uuid.Nil
	} else {
		desiredState, err := uuid.Parse(*in.DesiredState)
		if err != nil {
			return nil, fmt.Errorf("uuid.Parse(DesiredState) %w", err)
		}

		session.DesiredState = desiredState
	}

	if in.SubjectId == nil {
		session.SubjectID = uuid.Nil
	} else {
		subjectId, err := uuid.Parse(*in.SubjectId)
		if err != nil {
			return nil, fmt.Errorf("uuid.Parse(SubjectId): %w", err)
		}

		session.SubjectID = subjectId
	}

	created, err := s.intake.CreateSession(ctx, &session)
	if err != nil {
		return nil, fmt.Errorf("CreateSession: %w", err)
	}

	resp := pb.SessionResponse{
		Session: &pb.Session{
			Id:           created.ID.String(),
			DeviceId:     created.DeviceID.String(),
			DesiredState: created.DesiredState.String(),
			Description:  created.Description,
			StartTime:    created.StartTime,
			SensorId:     created.SensorID.String(),
			SessionType:  fmt.Sprintf("%d", created.SessionType),
			SubjectId:    created.SubjectID.String(),
		},
	}

	return &resp, nil
}
