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
	"trocup-message/routes"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// CORS activation for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                // Enable access from all domains
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS", // Allowed HTTP methods
	}))

	// Initialize MongoDB
	config.InitMongo()

	// Initialize Clerk
	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

	// Initialize the message repository
	// repository.InitMessageRepository()

	// Set up routes
	// handlers.SetupRoutes(app)
	routes.MessageRoutes(app)

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "5004" // Default port if not specified
	}

	// Start the WebSocket message handling in a goroutine
	go handlers.HandleMessages()

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
