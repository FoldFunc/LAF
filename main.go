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

	app.Post("/addLostAndFound", handlers.AddLostAndFoundHandler)
	app.Post("/getLostAndFound", handlers.GetLostAndFoundHandler)
	app.Post("/claimLostAndFound", handlers.ClaimLostAndFoundHandler)
	app.Get("/showAllLostAndFound", handlers.ShowAllLostAndFoundHadler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error while listening to the 8080 port: ", err)
	}
}
