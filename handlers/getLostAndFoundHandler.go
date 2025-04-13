package handlers

import (
	"log"

	"github.com/FoldFunc/LAF/database"
	"github.com/FoldFunc/LAF/structs"
	"github.com/gofiber/fiber/v2"
)

func GetLostAndFoundHandler(c *fiber.Ctx) error {
	log.Println("GetLostAndFoundHandler function called")
	var r structs.ReqestGetAndLostHandler
	if err := c.BodyParser(&r); err != nil {
		log.Println("Invalid JSON body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	if r.Name == "" || r.Notes == "" || r.Where == "" || r.WhoFound == "" {
		log.Println("The field is empty")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	if err := database.GetLostAndFoundDatabase(r); err != nil {
		log.Println("Error in GetLostAndFoundDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	log.Println("Reqest succesfull")
	return c.Status(fiber.StatusCreated).JSON(map[string]string{"message": "Lost and found added succesfully"})

}
