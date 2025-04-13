package handlers

import (
	"github.com/FoldFunc/LAF/database"
	"github.com/FoldFunc/LAF/structs"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetLostAndFoundHandler(c *fiber.Ctx) error {
	log.Println("GetLostAndFoundHandler function called")

	var r structs.ReqestGetAndLostHandler
	if err := c.BodyParser(&r); err != nil {
		log.Println("Invalid JSON body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	err, fetchedData := database.GetLostAndFoundDatabase(r)
	if err != nil {
		log.Println("Error in GetLostAndFoundDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	log.Println("Request successful")
	log.Println("Fetched data count:", len(fetchedData))

	// Optional: Deduplicate the data if that's what you're trying to do
	uniqueMap := make(map[string]bool)
	var uniqueData []structs.LostAndFound

	for _, item := range fetchedData {
		key := item.Name + item.WhoFound + item.Where + item.Notes + item.DateAdded
		if !uniqueMap[key] {
			uniqueMap[key] = true
			uniqueData = append(uniqueData, item)
		}
	}

	return c.JSON(fiber.Map{
		"message": "Lost and found fetched",
		"body":    uniqueData,
	})
}

