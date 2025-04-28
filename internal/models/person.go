package models

import "github.com/jackc/pgx/v5/pgtype"

type Person struct {
	ID         int              `json:"id,omitempty"`
	Name       string           `json:"name,omitempty"`
	Surname    string           `json:"surname,omitempty"`
	Patronymic *string          `json:"patronymic,omitempty"`
	Age        int              `json:"age,omitempty"`
	Gender     string           `json:"gender,omitempty"`
	National   []string         `json:"national,omitempty"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}
