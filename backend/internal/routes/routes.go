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
	app.Get("/api/v1/analytics/invoice-payments", handlers.GetInvoicePaymentStats)
	app.Get("/api/v1/analytics/invoice-quarterly", handlers.GetQuarterlyInvoiceOverview)
	app.Get("/api/v1/analytics/invoice-quarterly", handlers.GetQuarterlyInvoiceOverview)
	app.Get("/api/v1/analytics/ads-duration", handlers.GetAdvertisementDurationStats)
	app.Get("/api/v1/analytics/total-salary", handlers.GetTotalSalaryExpenseHandler)
	app.Get("/api/v1/analytics/num-employees", handlers.GetNumEmployeesHandler)
	app.Get("/api/v1/analytics/avg-salary", handlers.GetAvgSalaryHandler)

	//Contact us
	app.Post("/api/v1/add-contact", handlers.AddContactInfo)
	app.Get("/api/v1/get-contacts", handlers.GetAllContacts)

	app.Listen(":3000") // You can change the port as needed
}
