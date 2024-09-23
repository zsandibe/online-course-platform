package postgres

import "github.com/jmoiron/sqlx"

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}
