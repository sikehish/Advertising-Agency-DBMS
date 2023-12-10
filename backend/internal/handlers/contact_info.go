package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/models"
)

func AddContactInfo(c *fiber.Ctx) error {
	var contact models.ContactUs

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"error": "Invalid JSON format",
		})
	}

	// Validate name, email, and message
	if contact.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"error": "Please enter a valid name.",
		})
	}

	if contact.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"error": "Please enter a valid email.",
		})
	}

	if contact.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"error": "Please enter a message.",
		})
	}

	// Check if email already exists
	var existingContact models.ContactUs
	result := configs.DB.First(&existingContact, "email = ?", contact.Email)

	if result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":"fail",
			"message": "Your email is already registered. We will get back to you shortly.",
		})
	}

	configs.DB.Create(&contact)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Submission successful!",
	})
}

func GetAllContacts(c *fiber.Ctx) error {
	var contacts []models.ContactUs
	configs.DB.Find(&contacts)
	response := map[string]interface{}{
		"length":  len(contacts),
		"results": contacts,
	}
	return c.JSON(response)
}