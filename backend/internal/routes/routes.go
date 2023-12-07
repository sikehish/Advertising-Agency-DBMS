package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/handlers"
)

func SetupRoutes() {
	app := fiber.New()

	// Clients
	app.Get("/api/v1/clients", handlers.GetAllClients)
	app.Get("/api/v1/clients/:id", handlers.GetClientByID)
	app.Post("/api/v1/clients", handlers.AddClient)
	app.Patch("/api/v1/clients/:id", handlers.UpdateClient)
	app.Delete("/api/v1/clients/:id", handlers.DeleteClient)


	//Analytics
	app.Get("/api/v1/analytics/invoices", handlers.GetInvoicePaymentStats)

	app.Listen(":3000") // You can change the port as needed
}
