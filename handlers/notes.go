package handlers

import (
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/services"


	"github.com/gofiber/fiber/v2"
)

type APIStruct struct {
	Data interface{} `json:"data,omitempty"`
}

type NoteHandler struct {
	noteService services.NoteService
}

func NewNoteHandler(noteService services.NoteService) *NoteHandler  {
	return &NoteHandler{
		noteService: noteService,
	}
}

func (h *NoteHandler) GetAllNotes(c *fiber.Ctx) error {

	notes , err := h.noteService.GetAllNotes()

	if err != nil {
		return  c.Status(500).JSON(fiber.Map{"message": "Error Occured"})
	}

	return c.JSON(APIStruct{Data: notes})
}
