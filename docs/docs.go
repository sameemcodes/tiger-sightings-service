// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mohamed Sameem",
            "email": "mmmohamedsameem@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health-Controller"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/tiger/v1/checkIfTigerExists/{tigerId}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Check if a tiger exists by tigerId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tigerId",
                        "name": "tigerId",
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
        "/tiger/v1/create_new": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Create a new tiger",
                "parameters": [
                    {
                        "description": "Tiger body with timestamp format yyyy-mm-dd HH:ii:ss",
                        "name": "tiger",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tiger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/tiger/v1/deletebyTigerId/{tigerId}": {
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Delete a tiger by tigerId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tigerId",
                        "name": "tigerId",
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
        "/tiger/v1/fetch_all": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Get all tigers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tiger/v1/tiger_id/{tigerId}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Get a tiger by tigerId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tigerId",
                        "name": "tigerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tiger/v1/update": {
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tiger-Controller"
                ],
                "summary": "Update a tiger",
                "parameters": [
                    {
                        "description": "Tiger",
                        "name": "tiger",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tiger"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/tigerSighting/v1/create_new": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Create a new tiger sighting",
                "parameters": [
                    {
                        "description": "TigerSightingData",
                        "name": "sightingData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TigerSightingData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/tigerSighting/v1/deletebySightingId/{sightingId}": {
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Delete a tiger sighting by sightingId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sightingId",
                        "name": "sightingId",
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
        "/tigerSighting/v1/fetch_all": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Get all tiger sightings",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tigerSighting/v1/sighting_id/{sightingId}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Get a tiger sighting by sightingId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sightingId",
                        "name": "sightingId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tigerSighting/v1/tiger_id/{tigerId}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Get all tiger sightings for a tiger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tigerId",
                        "name": "tigerId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tigerSighting/v1/update": {
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TigerSighting-Controller"
                ],
                "summary": "Update a tiger sighting",
                "parameters": [
                    {
                        "description": "TigerSightingData",
                        "name": "sightingData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TigerSightingData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/v1/create_new": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Controller"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/v1/deletebyUserId/{userId}": {
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Controller"
                ],
                "summary": "Delete a user by userId",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/user/v1/fetch_all": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Controller"
                ],
                "summary": "Get all users",
                "responses": {}
            }
        },
        "/user/v1/update": {
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Controller"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/v1/user_id/{userId}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Controller"
                ],
                "summary": "Get a user by userId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Tiger": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "last_seen_coordinates_lat": {
                    "type": "number"
                },
                "last_seen_coordinates_lon": {
                    "type": "number"
                },
                "last_seen_timestamp": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tiger_id": {
                    "type": "string"
                }
            }
        },
        "models.TigerSightingData": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "sighting_id": {
                    "type": "string"
                },
                "sighting_image": {
                    "type": "string"
                },
                "tiger_id": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "tiggerhall-kittens",
	Description:      "Tiggerhall-Kittens",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
