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