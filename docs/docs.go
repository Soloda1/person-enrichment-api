// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create": {
            "post": {
                "description": "Создает нового человека с указанными данными и обогащает их дополнительной информацией",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Создать нового человека",
                "parameters": [
                    {
                        "description": "Данные для создания человека",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/createPerson.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/createPerson.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/delete/{id}": {
            "delete": {
                "description": "Удаляет человека по его ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Удалить человека",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID человека",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deletePerson.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/person/{id}": {
            "get": {
                "description": "Получает информацию о человеке по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Получить человека по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID человека",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/getPerson.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/persons": {
            "get": {
                "description": "Получает список людей с возможностью фильтрации и пагинации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Получить список людей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Фильтр по имени",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фильтр по фамилии",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фильтр по отчеству",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фильтр по полу",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фильтр по национальности",
                        "name": "national",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Минимальный возраст",
                        "name": "min_age",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Максимальный возраст",
                        "name": "max_age",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество записей на странице",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/getPersons.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Возвращает статус работоспособности API и текущее время",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Проверка работоспособности API",
                "responses": {
                    "200": {
                        "description": "Успешный ответ с сообщением pong и временем",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/update": {
            "put": {
                "description": "Обновляет информацию о человеке по его ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Обновить информацию о человеке",
                "parameters": [
                    {
                        "description": "Данные для обновления",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/updatePerson.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/updatePerson.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "createPerson.Request": {
            "description": "Структура запроса для создания нового человека",
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "name": {
                    "description": "Name имя человека\n@Description Имя человека\n@Required",
                    "type": "string"
                },
                "patronymic": {
                    "description": "Patronymic отчество человека\n@Description Отчество человека (опционально)",
                    "type": "string"
                },
                "surname": {
                    "description": "Surname фамилия человека\n@Description Фамилия человека\n@Required",
                    "type": "string"
                }
            }
        },
        "createPerson.Response": {
            "description": "Структура ответа при создании нового человека",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Сообщение об ошибке (если есть)",
                    "type": "string"
                },
                "person": {
                    "description": "Person данные созданного человека\n@Description Данные созданного человека",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Person"
                        }
                    ]
                },
                "status": {
                    "description": "Status статус операции\n@Description HTTP статус операции",
                    "type": "string"
                }
            }
        },
        "deletePerson.Response": {
            "description": "Структура ответа при удалении человека",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Сообщение об ошибке (если есть)",
                    "type": "string"
                },
                "status": {
                    "description": "Status статус операции\n@Description HTTP статус операции",
                    "type": "string"
                }
            }
        },
        "getPerson.Response": {
            "description": "Структура ответа с информацией о человеке",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Сообщение об ошибке (если есть)",
                    "type": "string"
                },
                "person": {
                    "description": "Person данные человека\n@Description Данные найденного человека",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Person"
                        }
                    ]
                },
                "status": {
                    "description": "Status статус операции\n@Description HTTP статус операции",
                    "type": "string"
                }
            }
        },
        "getPersons.Response": {
            "description": "Структура ответа со списком людей и дополнительной информацией",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Сообщение об ошибке (если есть)",
                    "type": "string"
                },
                "person": {
                    "description": "Persons список людей\n@Description Список найденных людей",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Person"
                    }
                },
                "status": {
                    "description": "Status статус операции\n@Description HTTP статус операции",
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "description": "Модель данных человека с обогащенной информацией",
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "age": {
                    "description": "Age возраст человека\n@Description Возраст человека (определяется автоматически)",
                    "type": "integer"
                },
                "created_at": {
                    "description": "CreatedAt время создания записи\n@Description Время создания записи",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender пол человека\n@Description Пол человека (определяется автоматически)",
                    "type": "string"
                },
                "id": {
                    "description": "ID уникальный идентификатор\n@Description Уникальный идентификатор человека",
                    "type": "integer"
                },
                "name": {
                    "description": "Name имя человека\n@Description Имя человека\n@Required",
                    "type": "string"
                },
                "national": {
                    "description": "National национальности\n@Description Список вероятных национальностей (определяется автоматически)",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "patronymic": {
                    "description": "Patronymic отчество человека\n@Description Отчество человека (опционально)",
                    "type": "string"
                },
                "surname": {
                    "description": "Surname фамилия человека\n@Description Фамилия человека\n@Required",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt время последнего обновления\n@Description Время последнего обновления записи",
                    "type": "string"
                }
            }
        },
        "updatePerson.Request": {
            "description": "Структура запроса для обновления информации о человеке",
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "age": {
                    "description": "Age новый возраст человека\n@Description Новый возраст человека (опционально)",
                    "type": "integer"
                },
                "gender": {
                    "description": "Gender новый пол человека\n@Description Новый пол человека (опционально)",
                    "type": "string"
                },
                "id": {
                    "description": "ID идентификатор человека\n@Description ID человека для обновления\n@Required",
                    "type": "integer"
                },
                "name": {
                    "description": "Name новое имя человека\n@Description Новое имя человека (опционально)",
                    "type": "string"
                },
                "national": {
                    "description": "National новые национальности\n@Description Новый список национальностей (опционально)",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "patronymic": {
                    "description": "Patronymic новое отчество человека\n@Description Новое отчество человека (опционально)",
                    "type": "string"
                },
                "surname": {
                    "description": "Surname новая фамилия человека\n@Description Новая фамилия человека (опционально)",
                    "type": "string"
                }
            }
        },
        "updatePerson.Response": {
            "description": "Структура ответа при обновлении информации о человеке",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Сообщение об ошибке (если есть)",
                    "type": "string"
                },
                "person": {
                    "description": "Person обновленные данные человека\n@Description Обновленные данные человека",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Person"
                        }
                    ]
                },
                "status": {
                    "description": "Status статус операции\n@Description HTTP статус операции",
                    "type": "string"
                }
            }
        },
        "utils.ErrorResponse": {
            "description": "Стандартная структура ответа при возникновении ошибки",
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error сообщение об ошибке\n@Description Описание ошибки",
                    "type": "string"
                },
                "status": {
                    "description": "Status HTTP статус\n@Description HTTP статус ошибки",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Person Enrichment API",
	Description:      "API сервис для обогащения данных о людях дополнительной информацией",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
