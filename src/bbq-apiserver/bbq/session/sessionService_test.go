package session

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

func TestNewSessionService(t *testing.T) {
	type args struct {
		cache      infrastructure.CacheService
		repository bbq.SessionRepository
	}
	tests := []struct {
		name string
		args args
		want bbq.SessionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSessionService(tt.args.cache, tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSessionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_GetSessions(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			got, err := s.GetSessions(tt.args.tenantID)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionService.GetSessions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessionService.GetSessions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_GetSessionByID(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
		id       uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			got, err := s.GetSessionByID(tt.args.tenantID, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionService.GetSessionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessionService.GetSessionByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_GetSessionByMonitorAddress(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
		address  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			got, err := s.GetSessionByMonitorAddress(tt.args.tenantID, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionService.GetSessionByMonitorAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessionService.GetSessionByMonitorAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_CreateSession(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
		entity   bbq.Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			got, err := s.CreateSession(tt.args.tenantID, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionService.CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessionService.CreateSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_UpdateSession(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
		entity   bbq.Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bbq.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			got, err := s.UpdateSession(tt.args.tenantID, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionService.UpdateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessionService.UpdateSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sessionService_DeleteSession(t *testing.T) {
	type fields struct {
		repository bbq.SessionRepository
		cache      infrastructure.CacheService
	}
	type args struct {
		tenantID uuid.UUID
		entity   bbq.Session
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
			s := &sessionService{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
			}
			if err := s.DeleteSession(tt.args.tenantID, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("sessionService.DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
