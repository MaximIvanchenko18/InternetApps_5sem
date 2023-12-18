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
        "/api/cargo": {
            "get": {
                "description": "Возвращает все доступные грузы с опциональной фильтрацией по названию и диапазону цены",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Грузы"
                ],
                "summary": "Получить все грузы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название груза для фильтрации",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Нижний порог цены для фильтрации",
                        "name": "low_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Верхний порог цены для фильтрации",
                        "name": "high_price",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.GetAllCargosResponse"
                        }
                    }
                }
            }
        },
        "/api/cargo/": {
            "post": {
                "description": "Добавить новый груз",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Грузы"
                ],
                "summary": "Добавить груз",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Изображение груза",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Название",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Английское название",
                        "name": "en_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Категория",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Цена",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Вес",
                        "name": "weight",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Объем",
                        "name": "capacity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/cargo/{cargo_id}": {
            "get": {
                "description": "Возвращает подробную информацию об одном грузе",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Грузы"
                ],
                "summary": "Получить один груз",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Cargo"
                        }
                    }
                }
            },
            "put": {
                "description": "Изменить данные о грузе",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Грузы"
                ],
                "summary": "Изменить груз",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Изображение груза",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Название",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Английское название",
                        "name": "en_name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Категория",
                        "name": "category",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Цена",
                        "name": "price",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Вес",
                        "name": "weight",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Объем",
                        "name": "capacity",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Описание",
                        "name": "description",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Удаляет груз по id",
                "tags": [
                    "Грузы"
                ],
                "summary": "Удалить груз",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/cargo/{cargo_id}/add_to_flight": {
            "post": {
                "description": "Добавить выбранный груз в черновик полета",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Грузы"
                ],
                "summary": "Добавить груз в полет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AddToFlightResp"
                        }
                    }
                }
            }
        },
        "/api/flights": {
            "get": {
                "description": "Возвращает все полеты с фильтрацией по статусу и дате формирования",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Полеты"
                ],
                "summary": "Получить все полеты",
                "parameters": [
                    {
                        "type": "string",
                        "description": "статус полета",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "начальная дата формирования",
                        "name": "form_date_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "конечная дата формирвания",
                        "name": "form_date_end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllFlightsResponse"
                        }
                    }
                }
            }
        },
        "/api/flights/user_confirm": {
            "put": {
                "description": "Сформировать полет пользователем",
                "tags": [
                    "Полеты"
                ],
                "summary": "Сформировать полет",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.FlightOutput"
                        }
                    }
                }
            }
        },
        "/api/flights/{flight_id}": {
            "get": {
                "description": "Возвращает подробную информацию о полете",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Полеты"
                ],
                "summary": "Получить один полет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.FlightResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Позволяет изменить тип ракеты",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Полеты"
                ],
                "summary": "Указать тип ракеты",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Тип ракеты",
                        "name": "rocket_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.SwaggerUpdateFlightRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.UpdateFlightResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет полет по id",
                "tags": [
                    "Полеты"
                ],
                "summary": "Удалить полет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/flights/{flight_id}/change_cargo/{cargo_id}": {
            "put": {
                "description": "Изменить количество груза в полете",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Полеты"
                ],
                "summary": "Изменить количество груза в полете",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "количество груза",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/flights/{flight_id}/delete_cargo/{cargo_id}": {
            "delete": {
                "description": "Удалить груз из полета",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Полеты"
                ],
                "summary": "Удалить груз из полета",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id груза",
                        "name": "cargo_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllCargosResponse"
                        }
                    }
                }
            }
        },
        "/api/flights/{flight_id}/moderator_confirm": {
            "put": {
                "description": "Завершить или отклонить полет модератором",
                "tags": [
                    "Полеты"
                ],
                "summary": "Завершить полет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id полета",
                        "name": "flight_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "подтвердить",
                        "name": "confirm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "Авторизует пользователя по логину, паролю и отдаёт jwt токен для дальнейших запросов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.SwaggerLoginResp"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "post": {
                "description": "Выход из аккаунта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Выйти из аккаунта",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/sign_up": {
            "post": {
                "description": "Регистрация нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.RegisterResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.SwaggerUpdateFlightRequest": {
            "type": "object",
            "properties": {
                "rocket_type": {
                    "type": "string"
                }
            }
        },
        "ds.Cargo": {
            "type": "object",
            "required": [
                "capacity",
                "category",
                "description",
                "en_name",
                "name",
                "price",
                "weight"
            ],
            "properties": {
                "capacity": {
                    "description": "m^3",
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "en_name": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "price": {
                    "description": "Rubles",
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                },
                "weight": {
                    "description": "kg",
                    "type": "number"
                }
            }
        },
        "schemes.AddToFlightResp": {
            "type": "object",
            "properties": {
                "cargo_count": {
                    "type": "integer"
                }
            }
        },
        "schemes.AllCargosResponse": {
            "type": "object",
            "properties": {
                "cargos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Cargo"
                    }
                }
            }
        },
        "schemes.AllFlightsResponse": {
            "type": "object",
            "properties": {
                "flights": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemes.FlightOutput"
                    }
                }
            }
        },
        "schemes.FlightOutput": {
            "type": "object",
            "properties": {
                "completion_date": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "customer": {
                    "type": "string"
                },
                "formation_date": {
                    "type": "string"
                },
                "moderator": {
                    "type": "string"
                },
                "rocket_type": {
                    "type": "string"
                },
                "shipment_status": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schemes.FlightResponse": {
            "type": "object",
            "properties": {
                "cargos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Cargo"
                    }
                },
                "flight": {
                    "$ref": "#/definitions/schemes.FlightOutput"
                }
            }
        },
        "schemes.FlightShort": {
            "type": "object",
            "properties": {
                "cargo_count": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schemes.GetAllCargosResponse": {
            "type": "object",
            "properties": {
                "cargos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Cargo"
                    }
                },
                "draft_flight": {
                    "$ref": "#/definitions/schemes.FlightShort"
                }
            }
        },
        "schemes.LoginReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30
                },
                "password": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "schemes.RegisterReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30
                },
                "password": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "schemes.RegisterResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "schemes.SwaggerLoginResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "schemes.UpdateFlightResponse": {
            "type": "object",
            "properties": {
                "flight": {
                    "$ref": "#/definitions/schemes.FlightOutput"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Cargo transfer to ISS",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
