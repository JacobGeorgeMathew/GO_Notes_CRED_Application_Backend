package repository

import (
	     "database/sql"
			 "github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
				)

type NoteRepository interface {
	GetAllNotes() ([]models.Note, error)
	GetNotesByUserID(userID int) ([]models.Note, error)
	CreateNote(note *models.CreateNoteStruct) error
	GetNoteByID(id int) (*models.Note, error)
	UpdateNote(note *models.Note) error
	DeleteNote(id, userID int) error
}

type noteStruct struct {
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) NoteRepository {
	return &noteStruct{db: db}
}

func (r *noteStruct) GetAllNotes() ([]models.Note, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM notes"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	
	return notes, rows.Err()
}

func (r *noteStruct) GetNotesByUserID(userID int) ([]models.Note, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM notes WHERE user_id = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	
	return notes, rows.Err()
}

func (r *noteStruct) CreateNote(note *models.CreateNoteStruct) error {
	query := `INSERT INTO notes (title, content, user_id, created_at, updated_at) 
			  VALUES (?, ?, ?, NOW(), NOW())`
	
	result, err := r.db.Exec(query, note.Title, note.Content, note.UserID)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	note.ID = int(id)
	return nil
}

func (r *noteStruct) GetNoteByID(id int) (*models.Note, error) {
	note := &models.Note{}
	query := `SELECT id, title, content, user_id, created_at, updated_at 
			  FROM notes WHERE id = ?`
	
	err := r.db.QueryRow(query, id).Scan(
		&note.ID, &note.Title, &note.Content, &note.UserID,
		&note.CreatedAt, &note.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return note, nil
}

func (r *noteStruct) UpdateNote(note *models.Note) error {
	query := `UPDATE notes SET title = ?, content = ?, updated_at = NOW() 
			  WHERE id = ? AND user_id = ?`
	
	_, err := r.db.Exec(query, note.Title, note.Content, note.ID, note.UserID)
	return err
}

func (r *noteStruct) DeleteNote(id, userID int) error {
	query := `DELETE FROM notes WHERE id = ? AND user_id = ?`
	_, err := r.db.Exec(query, id, userID)
	return err
}