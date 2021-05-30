// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/adverts/": {
            "get": {
                "description": "get all adverts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advert"
                ],
                "summary": "get adverts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/advert.AdvertInfo"
                            }
                        }
                    }
                }
            }
        },
        "/v1/categories/": {
            "get": {
                "description": "get all categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/category.CategoryInfo"
                            }
                        }
                    }
                }
            }
        },
        "/v1/rediness/": {
            "get": {
                "description": "check health",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "check health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.health"
                        }
                    }
                }
            }
        },
        "/v1/users/": {
            "get": {
                "description": "get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.UserInfo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "description": "create a user",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.NewUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.NewUser"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advert.AdvertInfo": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/category.CategoryInfo"
                },
                "category_uuid": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.UserInfo"
                },
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "category.CategoryInfo": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "handlers.health": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "user.NewUser": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "phone",
                "roles"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "user.UserInfo": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "localhost:3000",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Swagger SALES-API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
