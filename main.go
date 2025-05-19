package main

import (
	"backend-konstruksi/config"
	"backend-konstruksi/routes"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Update CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Add logging middleware
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Incoming %s request to %s", c.Method(), c.Path())
		return c.Next()
	})

	config.ConnectDB()
	routes.SetupRoutes(app)

	log.Printf("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
