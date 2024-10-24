package routes

import (
	"fmt"
	"trocup-message/handlers"
	"trocup-message/middleware"

	"github.com/gofiber/fiber/v2"
)

func MessageRoutes(app *fiber.App) {
	// PUBLIC
	app.Get("/health", handlers.HealthCheck)

	// PRIVATE : Routes protégées par le middleware ClerkAuthMiddleware
	api := app.Group("/api", middleware.ClerkAuthMiddleware)

	api.Get("/messages", handlers.GetMessages)
	// app.Get("/messages/:id", handlers.GetMessageByID)
	// api.Post("/messages", handlers.CreateMessage)
	// api.Put("/messages/:id", handlers.UpdateMessage)
	// api.Delete("/messages/:id", handlers.DeleteMessage)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
