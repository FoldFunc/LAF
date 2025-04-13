package handlers

import (
	"log"

	"github.com/FoldFunc/LAF/database"
	"github.com/FoldFunc/LAF/structs"
	"github.com/gofiber/fiber/v2"
)

func ClaimLostAndFoundHandler(c *fiber.Ctx) error {
	log.Println("ClaimLostAndFound function called")

	var r structs.ReqestClaimLostAndFound
	if err := c.BodyParser(&r); err != nil {
		log.Println("Invalid JSON body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid json body"})
	}

	if err := database.ClaimLostAndFoundDatabase(r); err != nil {
		log.Println("Error in ClaimLostAndFoundDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	log.Println("Reqes succesfull")

	return c.JSON(fiber.Map{
		"message": "Lost and found claimed",
	})
}
