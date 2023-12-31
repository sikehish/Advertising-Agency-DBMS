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

// ---------------------------------------------------------------

func GetQuarterlyInvoiceOverview(c *fiber.Ctx) error {
	var quarterlyOverviews []struct {
		Quarter       int     `json:"quarter"`
		NumInvoices   int     `json:"numInvoices"`
		TotalAmount   float64 `json:"totalAmount"`
		AverageAmount float64 `json:"averageAmount"`
	}
	
	// Execute the SQL query using GORM
	err := configs.DB.Table("invoices").
	Select("EXTRACT(QUARTER FROM invoice_date) AS quarter, COUNT(*) AS num_invoices, SUM(total_amount) AS total_amount, AVG(total_amount) AS average_amount").
	Group("quarter").
	Order("quarter").
	Scan(&quarterlyOverviews).Error
	
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	
	return c.JSON(quarterlyOverviews)
}


// ---------------------------------------------------------------


func GetAdvertisementDurationStats(c *fiber.Ctx) error {
	var durationStats []struct {
		DurationCategory string `json:"durationCategory"`
		NumAds           int    `json:"numAds"`
		TotalCost        float64 `json:"totalCost"`
	}

	// Execute the SQL query using GORM
	err := configs.DB.Raw(`
        SELECT
            CASE
                WHEN duration_days <= 20 THEN 'Less than 20 days'
                WHEN duration_days > 20 AND duration_days <= 40 THEN 'Between 20 and 40 days'
                WHEN duration_days > 40 THEN 'More than 40 days'
                ELSE 'Uncategorized'
            END AS duration_category,
            COUNT(*) AS num_ads,
            SUM(cost) AS total_cost
        FROM
            advertisement_placements
        GROUP BY
            duration_category
    `).Scan(&durationStats).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(durationStats)
}

// ---------------------------------------------------------------

// TotalSalaryExpenseHandler retrieves the total salary expense for each department
func GetTotalSalaryExpenseHandler(c *fiber.Ctx) error {
	var totalSalaryExpenses []struct {
		Department          string  `json:"department"`
		TotalSalaryExpense  float64 `json:"totalSalaryExpense"`
	}

	err := configs.DB.Table("employees").
		Select("department, SUM(salary) AS total_salary_expense").
		Group("department").
		Scan(&totalSalaryExpenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(totalSalaryExpenses)
}

// ---------------------------------------------------------------

// NumEmployeesHandler retrieves the total number of employees in each department
func GetNumEmployeesHandler(c *fiber.Ctx) error {
	var numEmployees []struct {
		Department string `json:"department"`
		NumEmployees int `json:"numEmployees"`
	}

	err := configs.DB.Table("employees").
		Select("department, COUNT(*) AS num_employees").
		Group("department").
		Scan(&numEmployees).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(numEmployees)
}

// ---------------------------------------------------------------

func GetAvgSalaryHandler(c *fiber.Ctx) error {
	var avgSalaries []struct {
		Department string  `json:"department"`
		AvgSalary  float64 `json:"avgSalary"`
	}

	err := configs.DB.Table("employees").
		Select("department, AVG(salary) AS avg_salary").
		Group("department").
		Scan(&avgSalaries).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(avgSalaries)
}