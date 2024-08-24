package entity

import "time"

type Payment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CourseID  int64     `json:"course_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
