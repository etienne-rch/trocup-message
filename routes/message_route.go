package routes

import (
	"fmt"
	"trocup-message/handlers"
	"trocup-message/middleware"

	"github.com/gofiber/fiber/v2"
)

func MessageRoutes(app *fiber.App) {
	// Routes publiques : accessibles sans authentification
	public := app.Group("/api")

	public.Get("/health", handlers.HealthCheck)
	public.Get("/ws", handlers.HandleConnections)

	// Routes protégées : accessibles uniquement avec authentification
	protected := app.Group("/api/protected", middleware.ClerkAuthMiddleware)

	protected.Get("/messages", handlers.GetMessages)
	protected.Get("/messages/:id", handlers.GetMessageByID)
	protected.Get("/messages/rooms/:id", handlers.GetMessagesByRoomID)
	protected.Post("/messages", handlers.CreateMessage)
	protected.Delete("/messages/:id", handlers.DeleteMessage)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
