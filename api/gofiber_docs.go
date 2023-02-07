// Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplategofiber = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "miniyus",
            "url": "https://miniyus.github.io",
            "email": "miniyu97@gmail.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get login user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "get my info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.User"
                        }
                    }
                }
            }
        },
        "/api/auth/password": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "reset login user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "reset password",
                "parameters": [
                    {
                        "description": "reset password body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_auth.ResetPasswordStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "sign up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "sign up body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_auth.SignUp"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api_auth.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/revoke": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "revoke current jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "revoke token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.StatusResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/token": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "login  body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_auth.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_auth.TokenInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/groups": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get all group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "get all groups",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.ListResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "update group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "group pk",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "creat group",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/groups.UpdateGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.ResponseGroup"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_error.ValidationErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "create group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "create group",
                "parameters": [
                    {
                        "description": "creat group",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/groups.CreateGroup"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/groups.ResponseGroup"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_error.ValidationErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/groups/name/{name}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get group by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "get group by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.ResponseGroup"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/groups/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get group by pk",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "get group by pk",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pk",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.ResponseGroup"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "delete group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "bool"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/worker/status": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "jobs status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "jobs status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jobs.GetStatus"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/worker/{worker}/jobs": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get jobs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "get jobs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "worker name",
                        "name": "worker",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jobs.GetJobs"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/worker/{worker}/jobs/{job}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "get job",
                "parameters": [
                    {
                        "type": "string",
                        "description": "worker name",
                        "name": "worker",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "job id",
                        "name": "job",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jobs.GetJob"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api_error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "description": "health check your server",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "health check your server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.StatusResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_auth.ResetPasswordStruct": {
            "type": "object",
            "required": [
                "password",
                "password_confirm"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                }
            }
        },
        "api_auth.SignIn": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_auth.SignUp": {
            "type": "object",
            "required": [
                "email",
                "password",
                "password_confirm",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_auth.SignUpResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "api_auth.TokenInfo": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api_error.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api_error.ValidationErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "failed_fields": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "auth.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "group_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "groups.CreateAction": {
            "type": "object",
            "required": [
                "method",
                "resource"
            ],
            "properties": {
                "method": {
                    "$ref": "#/definitions/permission.Method"
                },
                "resource": {
                    "type": "string"
                }
            }
        },
        "groups.CreateGroup": {
            "type": "object",
            "required": [
                "name",
                "permissions"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.CreatePermission"
                    }
                }
            }
        },
        "groups.CreatePermission": {
            "type": "object",
            "required": [
                "actions",
                "name"
            ],
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.CreateAction"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "groups.ListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.ResponseGroup"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "groups.ResponseAction": {
            "type": "object",
            "properties": {
                "method": {
                    "$ref": "#/definitions/permission.Method"
                },
                "resource": {
                    "type": "string"
                }
            }
        },
        "groups.ResponseGroup": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.ResponseAction"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "groups.UpdateGroup": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.CreatePermission"
                    }
                }
            }
        },
        "jobs.GetJob": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "job_id": {
                    "type": "string"
                },
                "meta": {
                    "type": "object",
                    "additionalProperties": true
                },
                "status": {
                    "$ref": "#/definitions/worker.JobStatus"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "worker_name": {
                    "type": "string"
                }
            }
        },
        "jobs.GetJobs": {
            "type": "object",
            "properties": {
                "jobs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/worker.Job"
                    }
                }
            }
        },
        "jobs.GetStatus": {
            "type": "object",
            "properties": {
                "worker_count": {
                    "type": "integer"
                },
                "workers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/worker.StatusWorkerInfo"
                    }
                }
            }
        },
        "permission.Method": {
            "type": "string",
            "enum": [
                "GET",
                "POST",
                "PUT",
                "PATCH",
                "DELETE"
            ],
            "x-enum-varnames": [
                "GET",
                "POST",
                "PUT",
                "PATCH",
                "DELETE"
            ]
        },
        "utils.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        },
        "worker.Job": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "job_id": {
                    "type": "string"
                },
                "meta": {
                    "type": "object",
                    "additionalProperties": true
                },
                "status": {
                    "$ref": "#/definitions/worker.JobStatus"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "worker_name": {
                    "type": "string"
                }
            }
        },
        "worker.JobStatus": {
            "type": "string",
            "enum": [
                "success",
                "fail",
                "wait",
                "progress"
            ],
            "x-enum-varnames": [
                "SUCCESS",
                "FAIL",
                "WAIT",
                "PROGRESS"
            ]
        },
        "worker.StatusWorkerInfo": {
            "type": "object",
            "properties": {
                "is_running": {
                    "type": "boolean"
                },
                "job_count": {
                    "type": "integer"
                },
                "max_job_count": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Bearer token type",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfogofiber holds exported Swagger Info so clients can modify it
var SwaggerInfogofiber = &swag.Spec{
	Version:          "1.1.6",
	Host:             "localhost:9090",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "gofiber Swagger API Documentation",
	Description:      "gofiber API",
	InfoInstanceName: "gofiber",
	SwaggerTemplate:  docTemplategofiber,
}

func init() {
	swag.Register(SwaggerInfogofiber.InstanceName(), SwaggerInfogofiber)
}
