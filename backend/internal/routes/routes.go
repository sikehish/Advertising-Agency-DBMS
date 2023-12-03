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

	app.Listen(":3000") // You can change the port as needed
}
