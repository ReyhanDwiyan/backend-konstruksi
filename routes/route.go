package routes

import (
	"backend-konstruksi/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	proyek := api.Group("/proyek")
	proyek.Get("/", controllers.GetAllProyek)
	proyek.Post("/", controllers.CreateProyek)
	proyek.Put("/:id", controllers.UpdateProyek)
	proyek.Delete("/:id", controllers.DeleteProyek)

	kontraktor := api.Group("/kontraktor")
	kontraktor.Get("/", controllers.GetAllKontraktor)
	kontraktor.Post("/", controllers.CreateKontraktor)

	material := api.Group("/material")
	material.Get("/", controllers.GetAllMaterial)
	material.Post("/", controllers.CreateMaterial)
	// Tambahkan endpoint lain nanti
}
