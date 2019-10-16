package subject

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

type subjectService struct {
	repository bbq.SubjectRepository
	cache      infrastructure.CacheService
}

// NewSubjectService will create an MonitorService
func NewSubjectService(cache infrastructure.CacheService, repository bbq.SubjectRepository) bbq.SubjectService {
	return &subjectService{repository: repository, cache: cache}
}

func (s *subjectService) GetOrCreateSubject(tenantID uuid.UUID, name string, description string) (bbq.Subject, error) {
	subject, err := s.repository.GetByName(tenantID, name)

	if err != nil {
		if err == sql.ErrNoRows {
			subject, err := s.repository.Create(bbq.Subject{
				Name:        name,
				Description: description,
				TenantID:    tenantID,
			})

			return subject, err
		}
		return bbq.Subject{}, err
	}

	return subject, err
}
