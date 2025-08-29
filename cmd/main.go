package main

import (
	"fmt"
	"log"

	handlers "github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	api := app.Group("/api")

	notes := api.Group("/notes")

	notes.Get("/", handlers.GetAllNotes)

	log.Fatal(app.Listen("127.0.0.1:5000"))
}
