package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/online-course-platform/internal/domain"
	"github.com/zsandibe/online-course-platform/internal/entity"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (p *UserPostgres) SignUp(user domain.SignUpRequest) (*entity.User, error) {
	return &entity.User{}, nil
}

func (p *UserPostgres) SignIn(credentials domain.SignInRequest) (*entity.User, error) {
	return &entity.User{}, nil
}
