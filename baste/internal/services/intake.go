package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/bbq/internal/repository"
	"github.com/ssargent/bbq/pkg/bbq"
	"go.uber.org/zap"
)

type IntakeService struct {
	db     *sqlx.DB
	cache  *cache.Cache
	logger *zap.Logger
}

func (s *IntakeService) CreateSession(ctx context.Context, session *bbq.Session) (*bbq.Session, error) {
	q := repository.Queries{}

	param := repository.InsertSessionParams{
		DeviceID:     session.DeviceID,
		DesiredState: session.DesiredState,
		Description:  session.Description,
		StartTime:    session.StartTime,
	}

	createdSession, err := q.InsertSession(ctx, s.db, &param)
	if err != nil {
		return nil, fmt.Errorf("InsertSession: %w", err)
	}

	return &bbq.Session{BbqSession: *createdSession}, nil
}

func NewIntakeService(cache *cache.Cache, db *sqlx.DB, logger *zap.Logger) *IntakeService {
	return &IntakeService{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}
