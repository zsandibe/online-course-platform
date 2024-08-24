package entity

import "time"

type Course struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID int64     `json:"instructor_id"`
	CategoryID   int64     `json:"category_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
