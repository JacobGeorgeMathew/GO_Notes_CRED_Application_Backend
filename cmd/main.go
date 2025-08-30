package main

import (
	"fmt"
	"log"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/config"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/database"
	handlers "github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/handlers"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LoadConfig()

	
	// Initialize database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	//Dependency Injection Approach

	// Initialize repository layer
	noteRepo := repository.NewNoteRepo(db)

	// Initialize service layer
	noteService := services.NewNoteService(noteRepo)

	// Initialize handler layer
	noteHandler := handlers.NewNoteHandler(noteService)
	
	app := fiber.New()

	api := app.Group("/api")

	notes := api.Group("/notes")

	notes.Get("/", noteHandler.GetAllNotes)

	fmt.Println("Server Started")
	log.Fatal(app.Listen("127.0.0.1:5000"))
}
