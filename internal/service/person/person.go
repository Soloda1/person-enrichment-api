package person

import (
	"context"
	"person-enrichment-api/internal/repository/person"
	"person-enrichment-api/internal/utils/logger"
)

type PersonService interface {
	CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error)
	GetPersonByID(ctx context.Context, person_id int) (*person.Person, error)
	GetAllPersons(ctx context.Context) ([]*person.Person, error)
	UpdatePerson(ctx context.Context, person *person.Person) error
	DeletePerson(ctx context.Context, person_id int) error
}

type Service struct {
	repo person.PersonRepository
	log  *logger.Logger
}

func NewService(repo person.PersonRepository, log *logger.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error) {
	return s.repo.Create(ctx, person)
}

func (s *Service) GetPersonByID(ctx context.Context, person_id int) (*person.Person, error) {
	return s.repo.GetByID(ctx, person_id)
}

func (s *Service) GetAllPersons(ctx context.Context) ([]*person.Person, error) {
	return s.repo.GetALl(ctx)
}

func (s *Service) UpdatePerson(ctx context.Context, person *person.Person) error {
	return s.repo.Update(ctx, person)
}

func (s *Service) DeletePerson(ctx context.Context, person_id int) error {
	return s.repo.Delete(ctx, person_id)
}
