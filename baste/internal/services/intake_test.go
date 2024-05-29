package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/bbq/internal/bbq/repository"
	"github.com/ssargent/bbq/pkg/bbq"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestIntakeService_CreateSession(t *testing.T) {
	type args struct {
		db      *pgxpool.Pool
		cache   *cache.Cache
		logger  *zap.Logger
		queries repository.Querier
		session *bbq.Session
	}
	tests := map[string]struct {
		args args
		want func(t *testing.T, got *bbq.Session, err error)
	}{
		// TODO: Add test cases.
		"default": {
			args: args{},
			want: func(t *testing.T, got *bbq.Session, err error) {
				t.Helper()
				require.NoError(t, err)
				require.NotNil(t, got)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &IntakeService{
				db:      tt.args.db,
				cache:   tt.args.cache,
				logger:  tt.args.logger,
				queries: tt.args.queries,
			}
			got, err := s.CreateSession(context.Background(), tt.args.session)
			tt.want(t, got, err)
		})
	}
}

func TestIntakeService_CreateReadings(t *testing.T) {
	type fields struct {
		db      *pgxpool.Pool
		cache   *cache.Cache
		logger  *zap.Logger
		queries repository.Querier
	}
	type args struct {
		readings []*bbq.SensorReading
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntakeService{
				db:      tt.fields.db,
				cache:   tt.fields.cache,
				logger:  tt.fields.logger,
				queries: tt.fields.queries,
			}
			if err := s.CreateReadings(context.Background(), tt.args.readings); (err != nil) != tt.wantErr {
				t.Errorf("IntakeService.CreateReadings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntakeService_prepareSession(t *testing.T) {
	type fields struct {
		db      *pgxpool.Pool
		cache   *cache.Cache
		logger  *zap.Logger
		queries repository.Querier
	}
	type args struct {
		session *bbq.Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntakeService{
				db:      tt.fields.db,
				cache:   tt.fields.cache,
				logger:  tt.fields.logger,
				queries: tt.fields.queries,
			}
			got, err := s.prepareSession(context.Background(), tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntakeService.prepareSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntakeService.prepareSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
