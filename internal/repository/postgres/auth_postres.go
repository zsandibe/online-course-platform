package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/online-course-platform/internal/domain"
	"github.com/zsandibe/online-course-platform/internal/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) SignUp(user domain.SignUpRequest) (*entity.User, error) {
	return &entity.User{}, nil
}

func (p *AuthPostgres) SignIn(credentials domain.SignInRequest) (*entity.User, error) {
	return &entity.User{}, nil
}
