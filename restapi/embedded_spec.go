package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.parking.v1+json"
  ],
  "produces": [
    "application/io.goswagger.parking.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The product of a parking.io Behnam Hajian",
    "title": "A Parking lot application",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/lot": {
      "get": {
        "tags": [
          "lot"
        ],
        "operationId": "findLots",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a list the lots",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/lot"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "lot"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/lot"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/lot"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/lot/status/{id}": {
      "get": {
        "tags": [
          "lot"
        ],
        "operationId": "getStstus",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/lot"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/lot/{id}": {
      "put": {
        "tags": [
          "lot"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/lot"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/lot"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "lot"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/lot/{lotId}/ticket/{ticketId}/carLeaves": {
      "post": {
        "tags": [
          "ticket"
        ],
        "operationId": "leaveParking",
        "responses": {
          "200": {
            "description": "returns an integer as price",
            "schema": {
              "$ref": "#/definitions/ticket"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "lotId",
          "in": "path",
          "required": true
        },
        {
          "type": "integer",
          "format": "int64",
          "name": "ticketId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/ticket/getTicket/lot/{lotId}/carSize/{carSize}": {
      "get": {
        "tags": [
          "ticket"
        ],
        "operationId": "getTicket",
        "responses": {
          "200": {
            "description": "returns a ticket",
            "schema": {
              "$ref": "#/definitions/ticket"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "lotId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "carSize",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "lot": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "address": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "lotType": {
          "type": "string",
          "minLength": 1
        },
        "mediumCarSize": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "remainderMediumCarSpaces": {
          "type": "integer",
          "format": "int64"
        },
        "remainderSmallCarSpaces": {
          "type": "integer",
          "format": "int64"
        },
        "smallCarSize": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "ticket": {
      "type": "object",
      "properties": {
        "amount": {
          "type": "integer",
          "format": "float64"
        },
        "carSize": {
          "type": "string",
          "minLength": 1
        },
        "entryTime": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "leaveTime": {
          "type": "string"
        },
        "licenseNumber": {
          "type": "string",
          "minLength": 1
        },
        "lotId": {
          "type": "integer",
          "format": "int64"
        },
        "spotNumber": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
}
