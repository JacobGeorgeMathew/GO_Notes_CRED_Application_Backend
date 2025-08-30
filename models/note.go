package models

// import "time"

// type Note struct {
// 	ID          int       `json:"id" db:"id"`
// 	Title       string    `json:"title" db:"title"`
// 	Content     string    `json:"content" db:"content"`
// 	UserID      int       `json:"user_id" db:"user_id"`
// 	CreatedAt   time.Time `json:"created_at" db:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
// }

type Note struct {
 	ID          int       `json:"id" db:"id"`
 	Title       string    `json:"title" db:"title"`
 	Content     string    `json:"content" db:"content"`
}