package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/models"
	"gorm.io/gorm"
)

func GetAllClients(c *fiber.Ctx) error {
	var clients []models.Client
	config.DB.Find(&clients)
	return c.JSON(clients)
}

func GetClientByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var client models.Client
	if err := configs.DB.First(&client, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Client not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	return c.JSON(client)
}

// Implement other CRUD handlers for clients, campaigns, advertisements, etc.
