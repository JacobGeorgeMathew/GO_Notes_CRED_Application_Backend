package repository

import (
	"database/sql"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User , error)
	GetUserByID(id int) (*models.User,error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo  {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password, created_at, updated_at) VALUES(?, ?, ?, NOW(), NOW())`

	result , err := r.db.Exec(query,user.Username,user.Email,user.Password)

	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	user.ID = int(id)
	return nil
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, created_at, updated_at 
			  FROM users WHERE email = ?`
	
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *userRepo) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, created_at, updated_at 
			  FROM users WHERE id = ?`
	
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
