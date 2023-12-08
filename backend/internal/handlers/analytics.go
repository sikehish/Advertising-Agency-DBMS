package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"gorm.io/gorm"
)

func GetInvoicePaymentStats(c *fiber.Ctx) error {
	var invoiceStats [] struct{
			Amount         float64 `json:"amount"`
			InvoicesCount  int     `json:"invoicesCount"`
			PaymentStatus  string  `json:"paymentStatus"`
	}

	

	err := configs.DB.Table("invoices").
		Select("SUM(total_amount) AS amount, COUNT(payment_status) AS invoices_count, payment_status").
		Group("payment_status").
		Scan(&invoiceStats).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

    return c.JSON(invoiceStats)
}

// QuarterlyInvoiceOverview represents the structure for the query result
type QuarterlyInvoiceOverview struct {
	Quarter       int     `json:"quarter"`
	NumInvoices   int     `json:"numInvoices"`
	TotalAmount   float64 `json:"totalAmount"`
	AverageAmount float64 `json:"averageAmount"`
}

// GetQuarterlyInvoiceOverview handles the request for quarterly invoice overview
func GetQuarterlyInvoiceOverview(c *fiber.Ctx) error {
	var quarterlyOverviews []QuarterlyInvoiceOverview

	// Execute the SQL query using GORM
	err := configs.DB.Table("invoices").
		Select("EXTRACT(QUARTER FROM invoice_date) AS quarter, COUNT(*) AS num_invoices, SUM(total_amount) AS total_amount, AVG(total_amount) AS average_amount").
		Group("quarter").
		Order("quarter").
		Scan(&quarterlyOverviews).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the data in JSON format
	return c.JSON(quarterlyOverviews)
}
