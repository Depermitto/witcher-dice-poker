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
            "name": "Piotr (Depermitto) Jabłoński",
            "email": "penciller@disroot.org"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hands": {
            "get": {
                "description": "Generate random dice poker hand",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hands"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Hand"
                        }
                    }
                }
            }
        },
        "/hands/eval": {
            "post": {
                "description": "Evaluate dice",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hands"
                ],
                "parameters": [
                    {
                        "description": "Raw dice to evaluate. Value range (1-6), array length (5)",
                        "name": "evalRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.evalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hand created from dice",
                        "schema": {
                            "$ref": "#/definitions/model.Hand"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/hands/switch": {
            "post": {
                "description": "Update dice poker hand",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hands"
                ],
                "parameters": [
                    {
                        "description": "Hand to modify along with list of dice indexes. Die at index will be switched with a new, randomly generated value. Dice indexes (1-5), array length (1-5)",
                        "name": "updateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.updateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Hand"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.evalRequest": {
            "type": "object",
            "properties": {
                "dice": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "handler.updateRequest": {
            "type": "object",
            "properties": {
                "hand": {
                    "$ref": "#/definitions/model.Hand"
                },
                "switches": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.Hand": {
            "type": "object",
            "properties": {
                "dice": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "leadval": {
                    "type": "integer"
                },
                "rank": {
                    "$ref": "#/definitions/model.HandRank"
                },
                "supval": {
                    "type": "integer"
                }
            }
        },
        "model.HandRank": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                6,
                7,
                8
            ],
            "x-enum-varnames": [
                "Nothing",
                "Pair",
                "TwoPairs",
                "ThreeOfAKind",
                "FiveHighStraight",
                "SixHighStraight",
                "FullHouse",
                "FourOfAKind",
                "FiveOfAKind"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Witcher Dice Poker API",
	Description:      "Webserver serving a complete implementation of Witcher 1 (2007) dice poker mini-game.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
