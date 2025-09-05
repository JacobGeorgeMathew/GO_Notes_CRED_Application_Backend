package handlers

import (
	"strconv"

	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/services"
	"github.com/gofiber/fiber/v2"
)

type APIStruct struct {
	Data interface{} `json:"data,omitempty"`
}

type NoteHandler struct {
	noteService services.NoteService
}

func NewNoteHandler(noteService services.NoteService) *NoteHandler {
	return &NoteHandler{noteService: noteService}
}

// Get user ID from context (set by auth middleware)
func getUserID(c *fiber.Ctx) (int, error) {
	userID := c.Locals("userID")
	if userID == nil {
		return 0, fiber.NewError(401, "Unauthorized")
	}
	return userID.(int), nil
}

func (h *NoteHandler) GetAllNotes(c *fiber.Ctx) error {

	notes , err := h.noteService.GetAllNotes()

	if err != nil {
		return  c.Status(500).JSON(fiber.Map{"message": "Error Occured"})
	}

	return c.JSON(APIStruct{Data: notes})
}


func (h *NoteHandler) GetUserNotes(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	notes, err := h.noteService.GetNotesByUserID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error occurred"})
	}

	return c.JSON(APIStruct{Data: notes})
}

func (h *NoteHandler) CreateNote(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	//fmt.Printf("User ID: %v",userID)

	var note models.CreateNoteStruct 
	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//Set user ID from token
	note.UserID = userID

	err = h.noteService.CreateNote(&note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error creating note"})
	}

	return c.Status(201).JSON(APIStruct{Data: note})
}

func (h *NoteHandler) GetNote(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	note, err := h.noteService.GetNoteByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
	}

	// Check if note belongs to user
	if note.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied"})
	}

	return c.JSON(APIStruct{Data: note})
}

func (h *NoteHandler) UpdateNote(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	note.ID = id
	note.UserID = userID

	err = h.noteService.UpdateNote(&note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error updating note"})
	}

	return c.JSON(APIStruct{Data: note})
}

func (h *NoteHandler) DeleteNote(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	err = h.noteService.DeleteNote(id, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error deleting note"})
	}

	return c.Status(204).Send(nil)
}