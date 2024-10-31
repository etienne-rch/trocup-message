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
	app.Get("/ws", handlers.HandleConnections)

	// PRIVATE : Routes protégées par le middleware ClerkAuthMiddleware
	api := app.Group("/api", middleware.ClerkAuthMiddleware)

	api.Get("/messages", handlers.GetMessages)
	api.Get("/messages/:id", handlers.GetMessageByID)
	api.Get("/messages/rooms/:id", handlers.GetMessagesByRoomID)
	api.Post("/messages", handlers.CreateMessage)
	api.Delete("/messages/:id", handlers.DeleteMessage)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
