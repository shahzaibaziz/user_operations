// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "User Operation System",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/user": {
      "post": {
        "description": "register new user to database",
        "tags": [
          "service"
        ],
        "operationId": "registerUser",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successfully save user object into database",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "$ref": "#/responses/401"
          },
          "409": {
            "$ref": "#/responses/409"
          },
          "default": {
            "$ref": "#/responses/500"
          }
        }
      }
    },
    "/user/{user_id}": {
      "get": {
        "description": "get user base on id",
        "tags": [
          "service"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfully save user object into database",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "$ref": "#/responses/401"
          },
          "404": {
            "$ref": "#/responses/404"
          },
          "default": {
            "$ref": "#/responses/500"
          }
        }
      },
      "delete": {
        "description": "delete user",
        "tags": [
          "service"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "successfully delete user"
          },
          "401": {
            "$ref": "#/responses/401"
          },
          "404": {
            "$ref": "#/responses/404"
          },
          "default": {
            "$ref": "#/responses/500"
          }
        }
      },
      "patch": {
        "description": "update user name or address",
        "tags": [
          "service"
        ],
        "operationId": "updateUser",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/update_user"
            }
          },
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfully update user",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "$ref": "#/responses/401"
          },
          "404": {
            "$ref": "#/responses/404"
          },
          "default": {
            "$ref": "#/responses/500"
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "fetch all users base on filter",
        "tags": [
          "service"
        ],
        "operationId": "getAllUsers",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "string",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successfully save user object into database",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user_response"
              }
            }
          },
          "204": {
            "description": "no content in database"
          },
          "401": {
            "$ref": "#/responses/401"
          },
          "404": {
            "$ref": "#/responses/404"
          },
          "default": {
            "$ref": "#/responses/500"
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string",
          "example": "error code"
        },
        "message": {
          "type": "string",
          "example": "error message"
        }
      }
    },
    "update_user": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "maxLength": 32,
          "minLength": 3
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "name",
        "address",
        "email"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string",
          "maxLength": 32,
          "minLength": 3
        }
      }
    },
    "user_response": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "created_at": {
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
        }
      }
    }
  },
  "responses": {
    "400": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "401": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "403": {
      "description": "Forbidden",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "404": {
      "description": "Not Found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "409": {
      "description": "Conflict",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "500": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "User Operation System",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/user": {
      "post": {
        "description": "register new user to database",
        "tags": [
          "service"
        ],
        "operationId": "registerUser",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successfully save user object into database",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Conflict",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{user_id}": {
      "get": {
        "description": "get user base on id",
        "tags": [
          "service"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfully save user object into database",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "delete user",
        "tags": [
          "service"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "successfully delete user"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "description": "update user name or address",
        "tags": [
          "service"
        ],
        "operationId": "updateUser",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/update_user"
            }
          },
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfully update user",
            "schema": {
              "$ref": "#/definitions/user_response"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "fetch all users base on filter",
        "tags": [
          "service"
        ],
        "operationId": "getAllUsers",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "string",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successfully save user object into database",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user_response"
              }
            }
          },
          "204": {
            "description": "no content in database"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string",
          "example": "error code"
        },
        "message": {
          "type": "string",
          "example": "error message"
        }
      }
    },
    "update_user": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "maxLength": 32,
          "minLength": 3
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "name",
        "address",
        "email"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string",
          "maxLength": 32,
          "minLength": 3
        }
      }
    },
    "user_response": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "created_at": {
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
        }
      }
    }
  },
  "responses": {
    "400": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "401": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "403": {
      "description": "Forbidden",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "404": {
      "description": "Not Found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "409": {
      "description": "Conflict",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "500": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
