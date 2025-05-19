package routes

import (
	"backend-konstruksi/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Add logging middleware for API group
	api.Use(func(c *fiber.Ctx) error {
		log.Printf("API Request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	proyek := api.Group("/proyek")
	proyek.Get("/", controllers.GetAllProyek)
	proyek.Get("/:id", controllers.GetProyekByID)
	proyek.Post("/", controllers.CreateProyek)
	proyek.Put("/:id", controllers.UpdateProyek)
	proyek.Delete("/:id", controllers.DeleteProyek)

	kontraktor := api.Group("/kontraktor")
	kontraktor.Get("/", controllers.GetAllKontraktor)
	kontraktor.Get("/:id", controllers.GetKontraktorByID)
	kontraktor.Post("/", controllers.CreateKontraktor)
	kontraktor.Put("/:id", controllers.UpdateKontraktor)
	kontraktor.Delete("/:id", controllers.DeleteKontraktor)

	material := api.Group("/material")
	material.Get("/", controllers.GetAllMaterial)
	material.Get("/:id", controllers.GetMaterialByID) // Add this line
	material.Post("/", controllers.CreateMaterial)
	material.Put("/:id", controllers.UpdateMaterial)    // Implement this controller
	material.Delete("/:id", controllers.DeleteMaterial) // Implement this controller
	// Tambahkan endpoint lain nanti
}
