package repository

import (
	     "database/sql"
			 "github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
				)

type NoteRepository interface {
	GetAllNotes() ([]models.Note,error)
}

type noteStruct struct {
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) *noteStruct {
	return &noteStruct{
		db: db,
	}
}

func (r *noteStruct) GetAllNotes() ([]models.Note, error){
	query := "SELECT * FROM Notes"

	rows , err := r.db.Query(query)

	if err != nil {
		return  nil,err
	}

	var notes []models.Note
	for rows.Next() {
		var note models.Note

		err := rows.Scan(&note.ID,&note.Title,&note.Content)
		if err != nil {
		return  nil,err
	  }
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}