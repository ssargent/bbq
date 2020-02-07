package subject

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
)

type subjectRepository struct {
	database *sql.DB
}

// NewSubjectRepository wtse-1
func NewSubjectRepository(database *sql.DB) bbq.SubjectRepository {
	return &subjectRepository{database: database}
}

/*

type SubjectRepository interface {
	GetByID(tenantID uuid.UUID, subjectID uuid.UUID) (Subject, error)
	GetByName(tenantID uuid.UUID, name string) (Subject, error)
	Create(entity Subject) (Subject, error)
	Update(entity Subject) (Subject, error)
	Delete(entity Subject) error
}

*/

func (s *subjectRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.Subject, error) {
	var subjects []bbq.Subject
	query := `select id, uid, name, description, tenantid from bbq.subjects
		      where (tenantid = $1 or tenantid is null)`

	rows, err := s.database.Query(query, tenantID) //.Scan(&sub.ID, sub.Uid, sub.Name, sub.Description, sub.TenantID)

	if err != nil {
		return []bbq.Subject{}, err
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var sub bbq.Subject
		if err := rows.Scan(&sub.ID, &sub.Uid, &sub.Name, &sub.Description, &sub.TenantID); err != nil {
			return nil, err
		}
		subjects = append(subjects, sub)
	}

	return subjects, nil
}

func (s *subjectRepository) GetByID(tenantID uuid.UUID, id uuid.UUID) (bbq.Subject, error) {
	var sub bbq.Subject
	query := `select id, uid, name, description, tenantid from bbq.subjects
		      where uid = $2 and (tenantid = $1 or tenantid is null)
				`

	err := s.database.QueryRow(query, tenantID, id).Scan(&sub.ID, sub.Uid, sub.Name, sub.Description, sub.TenantID)

	if err != nil {
		return bbq.Subject{}, err
	}

	return sub, nil
}

func (s *subjectRepository) GetByName(tenantID uuid.UUID, name string) (bbq.Subject, error) {
	var sub bbq.Subject
	query := `select id, uid, name, description, tenantid from bbq.subjects
		      where name = $1 and (tenantid = $2 or tenantid is null)
				`

	err := s.database.QueryRow(query, name, tenantID).Scan(&sub.ID, &sub.Uid, &sub.Name, &sub.Description, &sub.TenantID)

	if err != nil {
		return bbq.Subject{}, err
	}

	return sub, nil
}

func (s *subjectRepository) Create(entity bbq.Subject) (bbq.Subject, error) {
	insertStatement := "insert into bbq.subjects (name, description, tenantid) values ($1, $2, $3) returning *"

	var sub bbq.Subject
	err := s.database.QueryRow(insertStatement, entity.Name, entity.Description, entity.TenantID).Scan(&sub.ID, &sub.Uid, &sub.Name, &sub.Description, &sub.TenantID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.Subject{}, errors.New("a subject with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.Subject{}, err
	}

	return sub, nil
}

func (s *subjectRepository) Update(entity bbq.Subject) (bbq.Subject, error) {
	panic("not implemented")
}

func (s *subjectRepository) Delete(entity bbq.Subject) error {
	panic("not implemented")
}
