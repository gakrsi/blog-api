package tests

import (
    "blog-api/config"
    "blog-api/models"
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func TestCreateBlogPost(t *testing.T) {
    app := fiber.New()
    config.ConnectDB()

    blogPost := models.BlogPost{
        Title:       "Test Title",
        Description: "Test Description",
        Body:        "Test Body",
    }

    jsonBytes, _ := json.Marshal(blogPost)

    req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewBuffer(jsonBytes))
    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, 201, resp.StatusCode)
}