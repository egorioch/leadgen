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
        "/api/buildings": {
            "get": {
                "description": "Возвращает список строений с возможностью фильтрации по городу, году и количеству этажей",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Получить список строений",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Город",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Год",
                        "name": "year_built",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество этажей",
                        "name": "floors",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Building"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request"
                    }
                }
            },
            "post": {
                "description": "Принимает данные о строении и сохраняет их в базе данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "создает строение",
                "parameters": [
                    {
                        "description": "Данные строения",
                        "name": "building",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Building"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Building"
                        }
                    },
                    "400": {
                        "description": "bad request"
                    },
                    "500": {
                        "description": "server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Building": {
            "type": "object",
            "required": [
                "city",
                "floors",
                "name",
                "year_built"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "floors": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year_built": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
