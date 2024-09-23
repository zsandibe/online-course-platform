package postgres

import "github.com/jmoiron/sqlx"

type LessonPostgres struct {
	db *sqlx.DB
}

func NewLessonPostgres(db *sqlx.DB) *LessonPostgres {
	return &LessonPostgres{db: db}
}
