package models

// PersonFilter фильтр для поиска людей
// @Description Фильтр для поиска людей по различным параметрам
type PersonFilter struct {
	// Name фильтр по имени
	// @Description Фильтр по имени
	Name *string `json:"name,omitempty"`

	// Surname фильтр по фамилии
	// @Description Фильтр по фамилии
	Surname *string `json:"surname,omitempty"`

	// Patronymic фильтр по отчеству
	// @Description Фильтр по отчеству
	Patronymic *string `json:"patronymic,omitempty"`

	// Gender фильтр по полу
	// @Description Фильтр по полу
	Gender *string `json:"gender,omitempty"`

	// National фильтр по национальности
	// @Description Фильтр по национальности
	National *string `json:"national,omitempty"`

	// MinAge минимальный возраст
	// @Description Минимальный возраст для фильтрации
	MinAge *int `json:"min_age,omitempty"`

	// MaxAge максимальный возраст
	// @Description Максимальный возраст для фильтрации
	MaxAge *int `json:"max_age,omitempty"`

	// Limit количество записей на странице
	// @Description Количество записей на странице
	Limit int `json:"limit"`

	// Offset смещение для пагинации
	// @Description Смещение для пагинации
	Offset int `json:"offset"`
}
