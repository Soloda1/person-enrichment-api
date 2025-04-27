package person

import (
	"context"
	"log/slog"
	"person-enrichment-api/external/enrichment"
	"person-enrichment-api/internal/repository/person"
	"person-enrichment-api/internal/utils/logger"
)

type PersonService interface {
	CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error)
	GetPersonByID(ctx context.Context, personId int) (*person.Person, error)
	GetAllPersons(ctx context.Context) ([]*person.Person, error)
	UpdatePerson(ctx context.Context, person *person.Person) (*person.Person, error)
	DeletePerson(ctx context.Context, personId int) error
}

type Service struct {
	repo   person.PersonRepository
	log    *logger.Logger
	enrich enrichment.EnrichmentService
}

type EnrichProvider interface {
	Enrich(ctx context.Context, name string) (*enrichment.EnrichedPerson, error)
}

func NewService(repo person.PersonRepository, log *logger.Logger, enrichmentService EnrichProvider) *Service {
	return &Service{
		repo:   repo,
		log:    log,
		enrich: enrichmentService,
	}
}

func (s *Service) CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error) {
	s.log.Info("Enriching person")
	enrichedPerson, err := s.enrich.Enrich(ctx, person.Name)
	if err != nil {
		s.log.Debug("Error enriching person", slog.String("error", err.Error()))
		return nil, err
	}

	person.Age = enrichedPerson.Age
	person.Gender = enrichedPerson.Gender
	person.National = enrichedPerson.National

	return s.repo.Create(ctx, person)
}

func (s *Service) GetPersonByID(ctx context.Context, personId int) (*person.Person, error) {
	return s.repo.GetByID(ctx, personId)
}

func (s *Service) GetAllPersons(ctx context.Context) ([]*person.Person, error) {
	return s.repo.GetALl(ctx)
}

func (s *Service) UpdatePerson(ctx context.Context, person *person.Person) (*person.Person, error) {
	return s.repo.Update(ctx, person)
}

func (s *Service) DeletePerson(ctx context.Context, personId int) error {
	return s.repo.Delete(ctx, personId)
}
