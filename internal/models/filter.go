package models

type PersonFilter struct {
	Name       *string
	Surname    *string
	Patronymic *string
	MinAge     *int
	MaxAge     *int
	Gender     *string
	National   *string
	Limit      int
	Offset     int
}
