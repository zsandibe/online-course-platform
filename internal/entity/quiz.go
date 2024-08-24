package entity

import "time"

type Quiz struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CourseID  int64     `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Question struct {
	ID        int64     `json:"id"`
	QuizID    int64     `json:"quiz_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	ID         int64     `json:"id"`
	QuestionID int64     `json:"question_id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
