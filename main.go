// filepath: backend-konstruksi/main.go
package main

import (
	"backend-konstruksi/config"
	"backend-konstruksi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Enable CORS
	app.Use(cors.New())

	// Connect ke MongoDB
	config.ConnectDB()

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server
	app.Listen(":3000")
}
