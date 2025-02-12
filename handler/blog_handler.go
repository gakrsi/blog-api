package handlers

import (
    "blog-api/config"
    "blog-api/models"
    "time"

    "github.com/gofiber/fiber/v2"
)

// CreateBlogPost godoc
// @Summary Create a new blog post
// @Description Create a new blog post with the given input data
// @Tags blog-posts
// @Accept json
// @Produce json
// @Param blogPost body models.BlogPost true "Blog post object"
// @Success 201 {object} models.BlogPost
// @Router /api/blog-post [post]
func CreateBlogPost(c *fiber.Ctx) error {
    blogPost := new(models.BlogPost)
    if err := c.BodyParser(blogPost); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    // blogPost.CreatedAt = time.Now()
    // blogPost.UpdatedAt = time.Now()

    if err := config.DB.Create(&blogPost).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not create blog post"})
    }

    return c.Status(201).JSON(blogPost)
}

// GetAllBlogPosts godoc
// @Summary Get all blog posts
// @Description Get a list of all blog posts
// @Tags blog-posts
// @Produce json
// @Success 200 {array} models.BlogPost
// @Router /api/blog-post [get]
func GetAllBlogPosts(c *fiber.Ctx) error {
    var blogPosts []models.BlogPost
    if err := config.DB.Find(&blogPosts).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch blog posts"})
    }
    return c.JSON(blogPosts)
}

// GetBlogPost godoc
// @Summary Get a blog post by ID
// @Description Get a single blog post by its ID
// @Tags blog-posts
// @Param id path int true "Blog post ID"
// @Produce json
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [get]
func GetBlogPost(c *fiber.Ctx) error {
    id := c.Params("id")
    var blogPost models.BlogPost

    if err := config.DB.First(&blogPost, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Blog post not found"})
    }

    return c.JSON(blogPost)
}

// UpdateBlogPost godoc
// @Summary Update a blog post
// @Description Update a blog post by its ID
// @Tags blog-posts
// @Accept json
// @Produce json
// @Param id path int true "Blog post ID"
// @Param blogPost body models.BlogPost true "Blog post object"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [patch]
func UpdateBlogPost(c *fiber.Ctx) error {
    id := c.Params("id")
    var blogPost models.BlogPost

    if err := config.DB.First(&blogPost, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Blog post not found"})
    }

    if err := c.BodyParser(&blogPost); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    blogPost.UpdatedAt = time.Now()
    config.DB.Save(&blogPost)

    return c.JSON(blogPost)
}

// DeleteBlogPost godoc
// @Summary Delete a blog post
// @Description Delete a blog post by its ID
// @Tags blog-posts
// @Param id path int true "Blog post ID"
// @Success 204 "No Content"
// @Router /api/blog-post/{id} [delete]
func DeleteBlogPost(c *fiber.Ctx) error {
    id := c.Params("id")
    var blogPost models.BlogPost

    if err := config.DB.First(&blogPost, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Blog post not found"})
    }

    config.DB.Delete(&blogPost)
    return c.SendStatus(204)
}