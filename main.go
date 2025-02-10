package main

import (
    "blog-api/config"
    "blog-api/routes"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/swagger"
)

// @title Blog API
// @version 1.0
// @description A simple blog API
// @host localhost:3000
// @BasePath /api
func main() {
    config.ConnectDB()

    app := fiber.New()
    app.Use(cors.New())
    routes.SetupRoutes(app)
    app.Get("/swagger/*", swagger.HandlerDefault)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}




