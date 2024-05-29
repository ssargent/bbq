package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/bbq/internal/bbq/repository"
	"github.com/ssargent/bbq/internal/caching"
	"github.com/ssargent/bbq/pkg/bbq"
	"go.uber.org/zap"
)

type IntakeService struct {
	db      *pgxpool.Pool
	cache   *cache.Cache
	logger  *zap.Logger
	queries repository.Querier
}

func (s *IntakeService) CreateSession(ctx context.Context, session *bbq.Session) (*bbq.Session, error) {
	session, err := s.prepareSession(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("prepareSession: %w", err)
	}

	param := repository.InsertSessionParams{
		DeviceID:     session.DeviceID,
		DesiredState: session.DesiredState,
		Description:  session.Description,
		StartTime:    time.Now().Unix(),
		SensorID:     session.SensorID,
		SessionType:  1,
	}

	createdSession, err := s.queries.InsertSession(ctx, s.db, &param)
	if err != nil {
		return nil, fmt.Errorf("InsertSession: %w", err)
	}

	// its odd to error here, but we want to know a problem happened.
	if err := s.cache.Add(
		caching.SessionCacheKey(createdSession.ID),
		createdSession,
		caching.SessionCacheDuration()); err != nil {
		return nil, fmt.Errorf("cache.Add: %w", err)
	}

	return &bbq.Session{BbqSession: *createdSession}, nil
}

func (s *IntakeService) CreateReadings(ctx context.Context, readings []*bbq.SensorReading) error {
	// we can get readings in bulk, stopping the madness of a ton of grpc overhead.
	// create the parameters for insert here.
	txn, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("db.BeginTxx: %w", err)
	}

	for _, r := range readings {
		if r == nil {
			continue
		}
		reading := &repository.InsertSensorReadingParams{
			SessionID:   r.SessionID,
			ProbeNumber: r.ProbeNumber,
			Temperature: r.Temperature,
			/*
			 let the reading tell us when it happened.  The device on the other end can
			 batch via whatever mechanism it wants and that will not affect the temporal granularity
			*/
			ReadingOccurred: r.ReadingOccurred,
		}

		if err := s.queries.InsertSensorReading(ctx, txn, reading); err != nil {
			txerr := txn.Rollback(ctx)
			if txerr != nil {
				return fmt.Errorf("txn.Rollback: %w, Original: InsertSensorReading: %w", txerr, err)
			}

			return fmt.Errorf("InsertSensorReading: %w", err)
		}
	}

	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("txn.Commit: %w", err)
	}

	return nil
}

func (s *IntakeService) prepareSession(
	ctx context.Context,
	session *bbq.Session) (*bbq.Session, error) {
	if session.DeviceName == nil {
		dev, err := s.queries.GetDefaultDevice(ctx, s.db)
		if err != nil {
			return nil, fmt.Errorf("GetDefaultDevice: %w", err)
		}

		session.DeviceID = dev.ID
	}

	if session.DeviceName != nil {
		dev, err := s.queries.GetDeviceByName(ctx, s.db, *session.DeviceName)
		if err != nil {
			return nil, fmt.Errorf("GetDeviceByName: %w", err)
		}

		session.DeviceID = dev.ID
	}

	if session.SensorName == nil {
		sensor, err := s.queries.GetDefaultSensor(ctx, s.db)
		if err != nil {
			return nil, fmt.Errorf("GetDefaultSensor: %w", err)
		}

		session.SensorID = sensor.ID
	}

	if session.DeviceName != nil {
		sensor, err := s.queries.GetSensorByName(ctx, s.db, *session.SensorName)
		if err != nil {
			return nil, fmt.Errorf("GetSensorByName: %w", err)
		}

		session.SensorID = sensor.ID
	}
	return session, nil
}

func NewIntakeService(cache *cache.Cache, logger *zap.Logger, db *pgxpool.Pool, q repository.Querier) *IntakeService {
	return &IntakeService{
		db:      db,
		cache:   cache,
		logger:  logger,
		queries: q,
	}
}
