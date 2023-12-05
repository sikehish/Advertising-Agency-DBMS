package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/models"
	"gorm.io/gorm"
)

func GetAllClients(c *fiber.Ctx) error {
	var clients []models.Client
	configs.DB.Find(&clients)
	return c.Status(200).JSON(clients)
}

func GetClientByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client
	err := configs.DB.First(&client, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Client not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	//OR
	// client = config.Database.Find(&dog, id)

    // if client.RowsAffected == 0 {
    //     return c.SendStatus(404)
    // }

	return c.Status(200).JSON(client)
}


func AddClient(c *fiber.Ctx) error {
    client := new(models.Client)

    // Parse the request body into the client struct
    if err := c.BodyParser(client); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request payload",
            "details": fiber.Map{
                "message": err.Error(),
            },
        })
    }

    // Check if a record with the same ClientID already exists
    if err := configs.DB.First(client, client.ClientID).Error; err == nil {

        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Client creation failed",
            "details": fiber.Map{
                "message": "Client with the same ID already exists",
            },
        })
    } 
    // configs.DB.First(client, client.ClientID) queries the database to check if a record with the same ClientID already exists. If it finds a matching record, it loads the data into the same client struct. If no record is found, it returns an error, and the existing client struct remains unchanged.

    // Create a new client in the database
    if err := configs.DB.Create(client).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal server error",
            "details": fiber.Map{
                "message": err.Error(),
            },
        })
    }

    return c.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClient(c *fiber.Ctx) error {
    id := c.Params("id")

    client := new(models.Client)
    if err := c.BodyParser(client); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    existingClient := new(models.Client)
    if err := configs.DB.First(existingClient, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Client not found"})
    }


    if err := configs.DB.Model(existingClient).Updates(client).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update client"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"updatedClient": existingClient}})
}


func DeleteClient(c *fiber.Ctx) error {

	clientID := c.Params("id")

	var existingClient models.Client
	result := configs.DB.First(&existingClient, clientID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Client not found"})
	}

	// Delete the client from the database
	configs.DB.Delete(&existingClient)

	// Return a success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Client deleted successfully", "deletedClient": existingClient})
}
