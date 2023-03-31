// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "MKN Support",
            "email": "mkn-notifyer@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/favorites": {
            "get": {
                "description": "Gets favorite proj",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets favorite proj",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Project"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "Logout user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/owned_projects": {
            "get": {
                "description": "Gets all owned proj",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets all owned proj",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Project"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/project": {
            "post": {
                "description": "Creates project",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "summary": "Creates project",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "description": "Gets all proj",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets all proj",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Project"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Signup user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signup user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/undelivered_notifications": {
            "get": {
                "description": "Gets undelivired notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets undelivired notifications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Notification"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/upcoming": {
            "get": {
                "description": "Gets upcoming notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets upcoming notifications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Notification"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}": {
            "put": {
                "description": "Updates project",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "change"
                ],
                "summary": "Updates project",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes project",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Deletes project",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/collaborator": {
            "post": {
                "description": "Adds collaborators",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "summary": "Adds collaborators",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes collaborator",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Deletes collaborator",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/collaborators": {
            "get": {
                "description": "Gets collaborators",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets collaborators",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Collaborations"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/section": {
            "post": {
                "description": "Creates section",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "summary": "Creates section",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/sections": {
            "get": {
                "description": "Gets all sections",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets all sections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Section"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/{section}": {
            "put": {
                "description": "Updates section",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "change"
                ],
                "summary": "Updates section",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes section",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Deletes section",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/{section}/notification": {
            "post": {
                "description": "Creates notification",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "summary": "Creates notification",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/{section}/notifications": {
            "get": {
                "description": "Gets All Notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets All Notifications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Notification"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        },
        "/{project}/{section}/{notification}": {
            "get": {
                "description": "Gets Notification by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Gets Notification",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Notification"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "change"
                ],
                "summary": "Updates notifications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes notifications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Deletes notifications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.Collaborations": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "ds.Notification": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "error_status": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "section_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "ds.Project": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "ds.Section": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "internal_app.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MKN API",
	Description:      "Notification backend service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
