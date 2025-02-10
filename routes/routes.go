package routes

import (
    "blog-api/handler"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    // Blog routes
    api.Post("/blog-post", handlers.CreateBlogPost)
    api.Get("/blog-post", handlers.GetAllBlogPosts)
    api.Get("/blog-post/:id", handlers.GetBlogPost)
    api.Patch("/blog-post/:id", handlers.UpdateBlogPost)
    api.Delete("/blog-post/:id", handlers.DeleteBlogPost)
}