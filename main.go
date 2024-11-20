package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"trocup-message/config"
	"trocup-message/handlers"
	"trocup-message/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Import Swagger handler for Fiber
	"github.com/joho/godotenv"

	_ "trocup-message/docs" // Import des docs Swagger générées
)

// @title TrocUp Message API
// @version 1.0
// @description API pour gérer les messages dans l'application TrocUp
// @termsOfService http://trocup-message.com/terms

// @contact.name Support TrocUp Message
// @contact.email support@trocup-message.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5004
// @BasePath /
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Initialize MongoDB
	config.InitMongo()

	// Initialize the message repository
	repository.InitMessageRepository()

	// Set up routes
	handlers.SetupRoutes(app)

	// Add Swagger documentation route
	app.Get("/swagger/*", swagger.HandlerDefault) // Route pour Swagger

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "5004" // Default port if not specified
	}

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		if err := config.Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
