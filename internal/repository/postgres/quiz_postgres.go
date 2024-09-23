package postgres

import "github.com/jmoiron/sqlx"

type QuizPostgres struct {
	db *sqlx.DB
}

func NewQuizPostgres(db *sqlx.DB) *QuizPostgres {
	return &QuizPostgres{db: db}
}
