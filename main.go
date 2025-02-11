package main

import (
    "blog-api/config"
    "blog-api/routes"
    _ "blog-api/docs" 
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/swagger"
)

// @title Blog API
// @version 1.0
// @description Blog API with CRUD Operations
// @host blog-api-dag9.onrender.com
// @BasePath /
// @schemes https
func main() {
    // Initialize database
    config.ConnectDB()

    app := fiber.New(fiber.Config{
        EnableTrustedProxyCheck: true,
        TrustedProxies: []string{"0.0.0.0/0"},  // Be careful with this in production
    })

    // CORS
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Serve Swagger files
    app.Get("/swagger/*", swagger.New(swagger.Config{
        URL: "doc.json",  // The url pointing to API definition
        DeepLinking: false,
    }))

    // Setup routes
    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}