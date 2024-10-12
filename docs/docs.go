// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a server for fullstack-examination-2024.",
    "title": "fullstack-examination-2024 API",
    "contact": {},
    "license": {
      "name": "Apache 2.0"
    },
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "string"
                    }
                  }
                }
              ]
            }
          }
        }
      }
    },
    "/todos": {
      "get": {
        "tags": [
          "todos"
        ],
        "summary": "Find all todos",
        "parameters": [
          {
            "name": "key_word",
            "in": "query",
            "description": "Filter todos by task name, task Description, task priority, and task status",
            "required": false,
            "type": "string"
          },
          {
            "name": "page",
            "in": "query",
            "description": "Page number for pagination",
            "required": false,
            "type": "integer",
            "default": 1
          },
          {
            "name": "per_page",
            "in": "query",
            "description": "Number of todos per page",
            "required": false,
            "type": "integer",
            "default": 10
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "Data": {
                      "type": "array",
                      "items": {
                        "$ref": "#/definitions/model.Todo"
                      }
                    }
                  }
                }
              ]
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "todos"
        ],
        "summary": "Create a new todo",
        "parameters": [
          {
            "description": "json",
            "name": "request",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/handler.CreateRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseError"
                },
                {
                  "type": "object",
                  "properties": {
                    "data": {
                      "$ref": "#/definitions/model.Todo"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      }
    },
    "/todos/status": {
      "get": {
        "tags": [
          "todos"
        ],
        "summary": "Retrieve all todos with a specific status",
        "parameters": [
          {
            "name": "key_word",
            "in": "query",
            "description": "Filter todos by status",
            "required": false,
            "type": "string",
            "enum": [
              "to-do",
              "in-progress",
              "completed"
            ]
          },
          {
            "name": "page",
            "in": "query",
            "description": "Page number for pagination",
            "required": false,
            "type": "integer",
            "default": 1
          },
          {
            "name": "per_page",
            "in": "query",
            "description": "Number of todos per page",
            "required": false,
            "type": "integer",
            "default": 10
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "Data": {
                      "type": "array",
                      "items": {
                        "$ref": "#/definitions/model.Todo"
                      }
                    }
                  }
                }
              ]
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      }
    },
      "/todos/withoutPagination":{
"get": {
        "tags": [
          "todos"
        ],
        "summary": "Find all todos without pagination",
        "parameters": [],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "Data": {
                      "type": "array",
                      "items": {
                        "$ref": "#/definitions/model.Todo"
                      }
                    }
                  }
                }
              ]
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      }
      },
    "/todos/{id}": {
      "get": {
        "tags": [
          "todos"
        ],
        "summary": "Find a todo",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "Data": {
                      "$ref": "#/definitions/model.Todo"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "todos"
        ],
        "summary": "Update a todo",
        "parameters": [
          {
            "description": "body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/handler.UpdateRequestBody"
            }
          },
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/handler.ResponseData"
                },
                {
                  "type": "object",
                  "properties": {
                    "Data": {
                      "$ref": "#/definitions/model.Todo"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "summary": "Delete a todo",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/handler.ResponseError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "handler.CreateRequest": {
      "type": "object",
      "required": [
        "task"
      ],
      "properties": {
        "task": {
          "type": "string"
        }
      }
    },
    "handler.Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "handler.ResponseData": {
      "type": "object",
      "properties": {
        "data": {
          "description": "Data is the response data."
        }
      }
    },
    "handler.ResponseError": {
      "type": "object",
      "properties": {
        "errors": {
          "description": "Errors is the response errors.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/handler.Error"
          }
        }
      }
    },
    "handler.UpdateRequestBody": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/model.Status"
        },
        "task": {
          "type": "string"
        }
      }
    },
    "model.Status": {
      "type": "string",
      "enum": [
        "created",
        "processing",
        "done"
      ]
    },
    "model.Todo": {
      "type": "object",
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "status": {
          "$ref": "#/definitions/model.Status"
        },
        "task": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        }
      }
    }
  }
}`


// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "fullstack-examination-2024 API",
	Description:      "This is a server for fullstack-examination-2024.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
