package main

import (
	"github.com/FoldFunc/LAF/database"
	"github.com/FoldFunc/LAF/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.Init()
	app := fiber.New()

	app.Post("/addLostAndFound", handlers.GetLostAndFoundHandler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error while listening to the 8080 port: ", err)
	}
}
