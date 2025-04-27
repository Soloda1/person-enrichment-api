package enrichment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type EnrichedPerson struct {
	Name     string
	Age      int
	Gender   string
	National []string
}

type EnrichmentService interface {
	Enrich(ctx context.Context, name string) (*EnrichedPerson, error)
}

type Service struct {
	httpClient *http.Client
}

func NewEnrichmentService() *Service {
	return &Service{
		httpClient: &http.Client{},
	}
}

func (e *Service) Enrich(ctx context.Context, name string) (*EnrichedPerson, error) {
	var (
		age      int
		gender   string
		national []string
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		res, err := e.getAge(ctx, name)
		if err != nil {
			return err
		}
		age = res
		return nil
	})

	g.Go(func() error {
		res, err := e.getGender(ctx, name)
		if err != nil {
			return err
		}
		gender = res
		return nil
	})

	g.Go(func() error {
		res, err := e.getNationalities(ctx, name)
		if err != nil {
			return err
		}
		national = res
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &EnrichedPerson{
		Name:     name,
		Age:      age,
		Gender:   gender,
		National: national,
	}, nil
}

func (e *Service) getAge(ctx context.Context, name string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.agify.io/?name=%s", name), nil)
	if err != nil {
		return 0, err
	}
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Age int `json:"age"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Age, nil
}

func (e *Service) getGender(ctx context.Context, name string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.genderize.io/?name=%s", name), nil)
	if err != nil {
		return "", err
	}
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Gender string `json:"gender"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Gender, nil
}

func (e *Service) getNationalities(ctx context.Context, name string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.nationalize.io/?name=%s", name), nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var national []string
	for _, country := range data.Country {
		national = append(national, country.CountryID)
	}
	return national, nil
}
