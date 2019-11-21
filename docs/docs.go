// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-21 16:47:20.795075 +0530 IST m=+0.096116327

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
        "termsOfService": "https://www.exnomy.com/terms-and-privacy",
        "contact": {
            "name": "Exnomy Support",
            "url": "https://www.exnomy.com/terms-and-privacy",
            "email": "info@dexhigh.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/blockbankcoltd/Exnomy/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/assets": {
            "get": {
                "description": "**Get All the assets available on the exchange**\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failure]\n\"data\": {\n\"tokens\": [\n{\n\"symbol\": \"string\"\t\t\t\t\t\t\t\t[Symbol of the asset, ex - BTC]\n\"name\": \"string\",\t\t\t\t\t\t\t\t[Name of the asset, ex - Bitcoin]\n\"Address\": \"string\", \t\t\t\t\t\t\t[Exnomy Smart contract specific address]\n\"Decimal\": int,\t\t\t\t\t\t\t\t\t[Decimal of the asset]\n\"updatedAt\": \"time\",\t\t\t\t\t\t\t\t[Time when asset updated in DB]\n\"createdAt\": \"time\"\t\t\t\t\t\t\t\t[Time when asset created in DB]\n}\n]\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show all assets present at exnomy.com",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Response"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/candles": {
            "get": {
                "description": "**Get high, low, open, close and volume data for an interval defined between 'from' to 'to' unix timestamps.**\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"status\": int,  \t\t   \t      [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", \t\t\t      [Description String for Success and Failure]\n\"data\": {\n\"candles\": [  \t\t\t      [Candle Information]\n{\n\"time\": int,          [Unix timestamp of Candle, ex - 1574232000]\n\"open\": \"string\",     [Open Price in this TimeStamp]\n\"close\": \"string\",    [Close Price in this Timestamp]\n\"low\": \"string\",      [Lowest Trade Price in this timestamp]\n\"high\": \"string\",     [Highest timetsamp in this timestamp]\n\"volume\": \"string\"    [Total traded volume in this timestamp]\n}\n]\n}\n}\n` + "`" + `` + "`" + `` + "`" + `",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Bar data from 'from' to 'to'.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Market ID, ex XRP-BTC",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Start unix timestamp of interval",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "End unix timestamp of interval",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "granularity, In secounds, ex - 5minute ==\u003e 5 * 60 ==\u003e 300",
                        "name": "granularity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Bar"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/orderbook": {
            "get": {
                "description": "**Find the latest snapshot version 2 corresponding to MarketID**\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"status\": int,    \t\t\t\t[Status Code, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", \t\t\t\t[Description String for Success and Failure]\n\"data\": {\n\"orderBook\": {\n\"sequence\": int, \t\t[Database specific Sequence Number]\n\"bids\": \t\t\t\t[\"Array of Objects\", [Array bids Object]\n[\n\"string\", \t\t[Bid Price, ex - \"0.00003108\"]\n\"string\", \t\t[Bif Amount, ex - \"330\"]\n]\n],\n\"asks\": \t\t\t\t[Array of Objects, [Array asks Object]\n[\n\"string\", \t\t[Ask price, ex - \"0.00003108\"]\n\"string\", \t\t[Ask Amount, ex - \"330\"]\n]\n]\n}\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get OrderBook for a particular market ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MarketID (ex: XRP-BTC)",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SnapshotV2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/trades": {
            "get": {
                "description": "**Find all the trades for a particular market**\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failure]\n\"data\": {\n\"count\": int,   [Size of the trades array object]\n\"trades\": [     [Array of Objects]\n{\n\"id\": int,   [DB Speacific Trade ID ]\n\"transactionID\": 36671, [DB Speacific Transaction ID]\n\"transactionHash\": [Hash of Trade Transaction on Blockchain ex - 0xfba0a551880a7a9458fdcd9e55078645da4a11c399be91efdd19a6a82014e731],\n\"status\": \"string\", [Status of the Trade transaction, ex - successful, pending, canceled]\n\"marketID\": \"string\", [Market ID, ex - XRP-BTC]\n\"maker\": \"string\", [Maker Account Address ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3]\n\"taker\": \"string\", [Taker Account Address ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3],\n\"takerSide\": \"string\", [Side of Taker order ex - but OR sell]\n\"makerOrderID\": \"string, [OrderID of Maker, ex - 0xbf9abcaa08bb820f791206564d261e3e36309ac16f5b091e507f0b14fb369656],\n\"takerOrderID\": \"string, [OrderID of Taker, ex - 0xbf9abcaa08bb820f791206564d261e3e36309ac16f5b091e507f0b14fb369656],\n\"sequence\": uint, [DB specific sequence number]\n\"amount\": \"uint\", [Amount/Quantity of trade, ex - 230]\n\"price\": \"string\", [Price of Trade, ex - 0.00003096]\n\"executedAt\": \"time\", [Time When Trade executed,  ex - 2019-11-20T08:36:26Z]\n\"createdAt\": \"time\",  [Time When Trade created,  ex - 2019-11-20T08:36:26Z]\n\"updatedAt\": \"time\",  [Time When Trade executed,  ex - 2019-11-20T08:36:26Z]\n}\n]\n}\n}\n` + "`" + `` + "`" + `` + "`" + `",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show all trades for a market",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MarketID (ex XRP-BTC)",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.QueryTradeResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ApiError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                }
            }
        },
        "api.Bar": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "number"
                },
                "high": {
                    "type": "number"
                },
                "low": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "time": {
                    "type": "integer"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "api.QueryTradeResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "trades": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Trade"
                    }
                }
            }
        },
        "api.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "desc": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "api.SnapshotV2": {
            "type": "object",
            "properties": {
                "asks": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=365) %!s(*ast.BasicLit=\u0026{366 5 2}) string}"
                    }
                },
                "bids": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=329) %!s(*ast.BasicLit=\u0026{330 5 2}) string}"
                    }
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "models.Trade": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "executedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "maker": {
                    "type": "string"
                },
                "makerOrderID": {
                    "type": "string"
                },
                "marketID": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "sequence": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "taker": {
                    "type": "string"
                },
                "takerOrderID": {
                    "type": "string"
                },
                "takerSide": {
                    "type": "string"
                },
                "transactionHash": {
                    "type": "string"
                },
                "transactionID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ExnomyAuthToken": {
            "type": "apiKey",
            "name": "Authorization (DEX-EXONOMY-AUTH should be like {address}#DEX-EXONOMY-AUTH@{time}#{signature})",
            "in": "header"
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
	Version:     "1.0",
	Host:        "api.exnomy.com",
	BasePath:    "",
	Schemes:     []string{"https"},
	Title:       "Exnomy API Swagger Documentation",
	Description: "This is Exnomy API Server.",
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
