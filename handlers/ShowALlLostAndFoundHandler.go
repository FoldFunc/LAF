package handlers

import (
	"log"

	"github.com/FoldFunc/LAF/database"
	"github.com/gofiber/fiber/v2"
)

func ShowAllLostAndFoundHadler(c *fiber.Ctx) error {
	log.Println("ShowAllLostAndFoundHadler function called")

	fetchedData, err := database.ShowAllLostAndFoundDatabase()
	if err != nil {
		log.Println("Error in ShowAllLostAndFoundDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Fetching successful",
		"body":    fetchedData,
	})
}
