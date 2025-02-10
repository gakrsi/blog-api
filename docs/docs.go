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
        "/api/blog-post": {
            "get": {
                "description": "Get a list of all blog posts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog-posts"
                ],
                "summary": "Get all blog posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BlogPost"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new blog post with the given input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog-posts"
                ],
                "summary": "Create a new blog post",
                "parameters": [
                    {
                        "description": "Blog post object",
                        "name": "blogPost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            }
        },
        "/api/blog-post/{id}": {
            "get": {
                "description": "Get a single blog post by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog-posts"
                ],
                "summary": "Get a blog post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a blog post by its ID",
                "tags": [
                    "blog-posts"
                ],
                "summary": "Delete a blog post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update a blog post by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog-posts"
                ],
                "summary": "Update a blog post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Blog post object",
                        "name": "blogPost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BlogPost": {
            "type": "object",
            "required": [
                "body",
                "description",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Blog API",
	Description:      "A simple blog API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
