package services

import (
	//"errors"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
)

type NoteService interface {
	GetAllNotes() ([]models.Note, error)
	GetNotesByUserID(userID int) ([]models.Note, error)
	CreateNote(note *models.CreateNoteStruct) error
	GetNoteByID(id int) (*models.Note, error)
	UpdateNote(note *models.Note) error
	DeleteNote(id, userID int) error
}

type noteservice struct {
	noterepo repository.NoteRepository
}

func NewNoteService(s repository.NoteRepository) *noteservice {
	return &noteservice{
		noterepo: s,
	}
}

func (s *noteservice) GetAllNotes() ([]models.Note,error) {
	notes , err := s.noterepo.GetAllNotes()

	if err != nil {
		return nil,err
	}

	return  notes, nil
}

func (s *noteservice) GetNotesByUserID(userID int) ([]models.Note, error) {
	return s.noterepo.GetNotesByUserID(userID)
}

func (s *noteservice) CreateNote(note *models.CreateNoteStruct) error {
	return s.noterepo.CreateNote(note)
}

func (s *noteservice) GetNoteByID(id int) (*models.Note, error) {
	return s.noterepo.GetNoteByID(id)
}

func (s *noteservice) UpdateNote(note *models.Note) error {
	return s.noterepo.UpdateNote(note)
}

func (s *noteservice) DeleteNote(id, userID int) error {
	return s.noterepo.DeleteNote(id, userID)
}