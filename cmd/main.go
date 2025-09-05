package main

import (
	"fmt"
	"log"
	// "time"
	// "strconv"
	// "os"

	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/config"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/database"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/handlers"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/middleware"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/services"

	"github.com/gofiber/fiber/v2"
	 "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/helmet"
	// "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	noteRepo := repository.NewNoteRepo(db)
	userRepo := repository.NewUserRepo(db)

	// Initialize services
	noteService := services.NewNoteService(noteRepo)
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	noteHandler := handlers.NewNoteHandler(noteService)
	authHandler := handlers.NewAuthHandler(authService)
	
	app := fiber.New()

	// 🍪 CORS configuration for cookies
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5173", // Your frontend URLs
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true, // 🔥 IMPORTANT: Allow cookies to be sent
	}))

	// API routes
	api := app.Group("/api")

	// Auth routes with rate limiting
	auth := api.Group("/auth",) // 10 requests per minute
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)
	// auth.Post("/refresh", authHandler.RefreshToken)
	// auth.Post("/forgot-password", authHandler.ForgotPassword)
	// auth.Post("/reset-password", authHandler.ResetPassword)

	// Protected routes (authentication required)
	protected := api.Group("/", middleware.AuthMiddleware())
	protected.Get("/me", authHandler.GetMe)

	// User profile routes
	//profile := protected.Group("/profile")
	// profile.Get("/", authHandler.GetProfile)
	// profile.Put("/", authHandler.UpdateProfile)
	// profile.Post("/change-password", authHandler.ChangePassword)

	// Notes routes with rate limiting
	notes := protected.Group("/notes",)
	notes.Get("/", noteHandler.GetUserNotes)         // Get user's notes
	notes.Post("/", noteHandler.CreateNote)          // Create note
	notes.Get("/:id", noteHandler.GetNote)           // Get specific note
	notes.Put("/:id", noteHandler.UpdateNote)        // Update note
	notes.Delete("/:id", noteHandler.DeleteNote)     // Delete note (soft delete)
	//notes.Post("/:id/restore", noteHandler.RestoreNote) // Restore soft deleted note

	// Admin routes (admin role required)
	admin := protected.Group("/admin",)
	//admin.Get("/users", authHandler.GetAllUsers)
	admin.Get("/notes", noteHandler.GetAllNotes)
	//admin.Get("/stats", noteHandler.GetStats)
	//admin.Put("/users/:id/status", authHandler.UpdateUserStatus)

	fmt.Println("Server Started")
	log.Fatal(app.Listen("127.0.0.1:5000"))
}

// import ( 
// 	"fmt"
// 	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/utils"
// )

// func main()  {
// 	value , err := utils.GenerateJWT(123)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(value)
// 	fmt.Println("")
// 	fmt.Println(utils.ValidateJWT(value))
// }
