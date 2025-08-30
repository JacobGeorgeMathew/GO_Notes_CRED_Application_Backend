package services

import (
	//"errors"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
)

type NoteService interface {
	GetAllNotes() ([]models.Note,error)
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