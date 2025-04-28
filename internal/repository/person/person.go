package person

import (
	"context"
	"log/slog"
	"person-enrichment-api/internal/models"
	"person-enrichment-api/internal/repository"
	"person-enrichment-api/internal/utils/logger"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type PersonRepository interface {
	Create(ctx context.Context, person *models.Person) (*models.Person, error)
	GetByID(ctx context.Context, personId int) (*models.Person, error)
	GetALl(ctx context.Context, filter models.PersonFilter) ([]*models.Person, error)
	Update(ctx context.Context, person *models.Person) (*models.Person, error)
	Delete(ctx context.Context, personId int) error
}

type Repository struct {
	storage *repository.Storage
	log     *logger.Logger
}

func NewRepository(storage *repository.Storage, log *logger.Logger) *Repository {
	return &Repository{storage, log}
}

func (r *Repository) Create(ctx context.Context, person *models.Person) (*models.Person, error) {
	createdAt := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}

	args := pgx.NamedArgs{
		"name":       person.Name,
		"surname":    person.Surname,
		"patronymic": person.Patronymic,
		"age":        person.Age,
		"gender":     person.Gender,
		"national":   person.National,
		"created_at": createdAt,
		"updated_at": createdAt}

	query := `INSERT INTO Person (name, surname, patronymic, age, gender, national, created_at, updated_at)
	VALUES (@name, @surname, @patronymic, @age, @gender, @national, @created_at, @updated_at)
	RETURNING *`

	var createdPerson models.Person
	err := r.storage.Pool.QueryRow(ctx, query, args).Scan(
		&createdPerson.ID,
		&createdPerson.Name,
		&createdPerson.Surname,
		&createdPerson.Patronymic,
		&createdPerson.Age,
		&createdPerson.Gender,
		&createdPerson.National,
		&createdPerson.CreatedAt,
		&createdPerson.UpdatedAt,
	)

	if err != nil {
		r.log.Debug("Failed to insert person into database", slog.String("error", err.Error()))
		return nil, err
	}

	return &createdPerson, nil
}

func (r *Repository) Delete(ctx context.Context, personId int) error {
	args := pgx.NamedArgs{
		"person_id": personId,
	}

	query := `DELETE FROM Person WHERE person_id = @person_id`

	_, err := r.storage.Pool.Exec(ctx, query, args)

	if err != nil {
		r.log.Debug("Error deleting person", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, personId int) (*models.Person, error) {
	args := pgx.NamedArgs{"person_id": personId}
	query := `SELECT * FROM Person WHERE person_id = @person_id`
	var person models.Person
	err := r.storage.Pool.QueryRow(ctx, query, args).Scan(
		&person.ID,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Age,
		&person.Gender,
		&person.National,
		&person.CreatedAt,
		&person.UpdatedAt)
	if err != nil {
		r.log.Debug("Error getting person", slog.String("error", err.Error()))
		return nil, err
	}

	return &person, nil
}

func (r *Repository) GetALl(ctx context.Context, filter models.PersonFilter) ([]*models.Person, error) {
	args := pgx.NamedArgs{
		"limit":  filter.Limit,
		"offset": filter.Offset,
	}

	query := `SELECT * FROM Person WHERE true`

	if filter.Name != nil {
		query += " AND name ILIKE @name"
		args["name"] = "%" + *filter.Name + "%"
	}

	if filter.Surname != nil {
		query += " AND surname ILIKE @surname"
		args["surname"] = "%" + *filter.Surname + "%"
	}

	if filter.Patronymic != nil {
		query += " AND patronymic ILIKE @patronymic"
		args["patronymic"] = "%" + *filter.Patronymic + "%"
	}

	if filter.Gender != nil {
		query += " AND gender ILIKE @gender"
		args["gender"] = "%" + *filter.Gender + "%"
	}

	if filter.National != nil {
		query += " AND EXISTS (SELECT * FROM unnest(national) AS nat WHERE LOWER(nat) = @national) "
		args["national"] = *filter.National
	}

	if filter.MinAge != nil {
		query += " AND age >= @min_age"
		args["min_age"] = *filter.MinAge
	}

	if filter.MaxAge != nil {
		query += " AND age <= @max_age"
		args["max_age"] = *filter.MaxAge
	}

	query += " ORDER BY person_id LIMIT @limit OFFSET @offset"

	rows, err := r.storage.Pool.Query(ctx, query, args)
	if err != nil {
		r.log.Debug("Error getting persons", slog.String("error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var persons []*models.Person
	for rows.Next() {
		var person models.Person
		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.National,
			&person.CreatedAt,
			&person.UpdatedAt,
		)
		if err != nil {
			r.log.Debug("Error getting person", slog.String("error", err.Error()))
			return nil, err
		}

		persons = append(persons, &person)
	}

	return persons, nil
}

func (r *Repository) Update(ctx context.Context, person *models.Person) (*models.Person, error) {
	updatedAt := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}

	args := pgx.NamedArgs{
		"person_id":  person.ID,
		"updated_at": updatedAt,
	}

	query := "UPDATE Person SET updated_at = @updated_at"

	if person.Name != "" {
		query += ", name = @name"
		args["name"] = person.Name
	}

	if person.Surname != "" {
		query += ", surname = @surname"
		args["surname"] = person.Surname
	}

	if person.Patronymic != nil {
		query += ", patronymic = @patronymic"
		args["patronymic"] = person.Patronymic
	}

	if person.Age != 0 {
		query += ", age = @age"
		args["age"] = person.Age
	}

	if person.Gender != "" {
		query += ", gender = @gender"
		args["gender"] = person.Gender
	}

	if person.National != nil {
		query += ", national = @national"
		args["national"] = person.National
	}

	query += " WHERE person_id = @person_id RETURNING *"

	var updatedPerson models.Person
	err := r.storage.Pool.QueryRow(ctx, query, args).Scan(
		&updatedPerson.ID,
		&updatedPerson.Name,
		&updatedPerson.Surname,
		&updatedPerson.Patronymic,
		&updatedPerson.Age,
		&updatedPerson.Gender,
		&updatedPerson.National,
		&updatedPerson.CreatedAt,
		&updatedPerson.UpdatedAt,
	)
	if err != nil {
		r.log.Debug("Error updating person", slog.String("error", err.Error()))
		return nil, err
	}

	return &updatedPerson, nil
}
