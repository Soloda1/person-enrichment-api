package models

import "time"

// Person представляет собой модель данных человека
// @Description Модель данных человека с обогащенной информацией
type Person struct {
	// ID уникальный идентификатор
	// @Description Уникальный идентификатор человека
	ID int `json:"id" db:"person_id"`

	// Name имя человека
	// @Description Имя человека
	// @Required
	Name string `json:"name" db:"name" binding:"required"`

	// Surname фамилия человека
	// @Description Фамилия человека
	// @Required
	Surname string `json:"surname" db:"surname" binding:"required"`

	// Patronymic отчество человека
	// @Description Отчество человека (опционально)
	Patronymic *string `json:"patronymic,omitempty" db:"patronymic"`

	// Age возраст человека
	// @Description Возраст человека (определяется автоматически)
	Age int `json:"age" db:"age"`

	// Gender пол человека
	// @Description Пол человека (определяется автоматически)
	Gender string `json:"gender" db:"gender"`

	// National национальности
	// @Description Список вероятных национальностей (определяется автоматически)
	National []string `json:"national" db:"national"`

	// CreatedAt время создания записи
	// @Description Время создания записи
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt время последнего обновления
	// @Description Время последнего обновления записи
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
