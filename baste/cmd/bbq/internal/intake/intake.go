package intake

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/bbq/internal/bbq/repository"
	"github.com/ssargent/bbq/internal/config"
	"github.com/ssargent/bbq/internal/services"
	"github.com/ssargent/bbq/pkg/bbq"
	intakev1 "github.com/ssargent/public-apis/pkg/bbq/intake/v1"
	"github.com/ssargent/public-apis/pkg/bbq/intake/v1/intakev1connect"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IntakeConfig struct {
}

func NewIntakeServiceHandler(cfg *config.Config, cache *cache.Cache, db *pgxpool.Pool, logger *zap.Logger) (string, http.Handler, error) {
	path, handler := intakev1connect.NewIntakeServiceHandler(newIntakeServer(cfg, cache, db, logger))

	return path, handler, nil
}

type intakeServiceServer struct {
	intakev1connect.UnimplementedIntakeServiceHandler

	cfg    *config.Config
	cache  *cache.Cache
	db     *pgxpool.Pool
	logger *zap.Logger

	intake *services.IntakeService
}

func newIntakeServer(cfg *config.Config, cache *cache.Cache, db *pgxpool.Pool, logger *zap.Logger) *intakeServiceServer {
	intake := services.NewIntakeService(cache, logger, db, &repository.Queries{})
	return &intakeServiceServer{
		intake: intake,
		cfg:    cfg,
		cache:  cache,
		db:     db,
		logger: logger,
	}
}

func (i *intakeServiceServer) Record(ctx context.Context, req *connect.Request[intakev1.RecordRequest]) (*connect.Response[intakev1.RecordResponse], error) {
	var sensorReadingCount int
	var actualSensorReadingCount int

	// Getting the full count of readings helps us to not have to reallocate the array often.
	for _, r := range req.Msg.Reading {
		sensorReadingCount += len(r.Readings)
	}

	intakeSessionID := uuid.Nil
	if len(req.Msg.Reading) > 0 {
		sessionID, err := uuid.Parse(req.Msg.Reading[0].SessionId)
		if err != nil {
			return nil, fmt.Errorf("uuid.Parse(SessionId): %w", err)
		}
		intakeSessionID = sessionID
	}

	readings := make([]*bbq.SensorReading, sensorReadingCount)

	for _, rding := range req.Msg.Reading {
		for _, r := range rding.Readings {
			occurredAt := rding.RecordedAt
			if occurredAt == 0 {
				occurredAt = time.Now().Unix()
			}
			reading := bbq.SensorReading{}
			reading.SessionID = intakeSessionID
			reading.ProbeNumber = r.SensorNumber
			reading.Temperature = float64(r.Temperature)
			reading.ReadingOccurred = occurredAt

			readings[actualSensorReadingCount] = &reading
			actualSensorReadingCount++
		}
	}

	if actualSensorReadingCount != sensorReadingCount {
		i.logger.Warn("actual sensor reading count does not match expected",
			zap.Int("actualSensorReadingCount", actualSensorReadingCount),
			zap.Int("expectedSensorReadingCount", sensorReadingCount))
	}

	if err := i.intake.CreateReadings(ctx, readings); err != nil {
		return nil, fmt.Errorf("CreateReadings: %w", err)
	}

	res := connect.NewResponse[intakev1.RecordResponse](
		&intakev1.RecordResponse{
			SessionId:  intakeSessionID.String(),
			RecordedAt: timestamppb.Now(),
		},
	)

	return res, nil
}

func (i *intakeServiceServer) Session(ctx context.Context, req *connect.Request[intakev1.SessionRequest]) (*connect.Response[intakev1.SessionResponse], error) {
	var session bbq.Session

	in := req.Msg

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

	created, err := i.intake.CreateSession(ctx, &session)
	if err != nil {
		return nil, fmt.Errorf("CreateSession: %w", err)
	}

	res := connect.NewResponse[intakev1.SessionResponse](
		&intakev1.SessionResponse{
			Session: &intakev1.Session{
				Id:           created.ID.String(),
				DeviceId:     created.DeviceID.String(),
				DesiredState: created.DesiredState.String(),
				Description:  created.Description,
				StartTime:    created.StartTime,
				SensorId:     created.SensorID.String(),
				SessionType:  fmt.Sprintf("%d", created.SessionType),
				SubjectId:    created.SubjectID.String(),
			},
		},
	)

	return res, nil
}
