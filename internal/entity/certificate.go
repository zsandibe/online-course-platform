package entity

import "time"

type Certificate struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	CourseID int64     `json:"course_id"`
	IssuedAt time.Time `json:"issued_at"`
}
