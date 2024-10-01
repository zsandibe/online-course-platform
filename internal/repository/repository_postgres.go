package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/online-course-platform/internal/domain"
	"github.com/zsandibe/online-course-platform/internal/entity"
	"github.com/zsandibe/online-course-platform/internal/repository/postgres"
)

type UserRepository interface {
	SignUp(user domain.SignUpRequest) (*entity.User, error)
	SignIn(credentials domain.SignInRequest) (*entity.User, error)
}

type CourseRepository interface{}

type QuizRepository interface{}

type LessonRepository interface{}

type PaymentRepository interface{}

type Repository struct {
	UserRepository    UserRepository
	CourseRepository  CourseRepository
	QuizRepository    QuizRepository
	LessonRepository  LessonRepository
	PaymentRepository PaymentRepository
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:    postgres.NewUserPostgres(db),
		CourseRepository:  postgres.NewCoursePostgres(db),
		QuizRepository:    postgres.NewQuizPostgres(db),
		LessonRepository:  postgres.NewLessonPostgres(db),
		PaymentRepository: postgres.NewPaymentPostgres(db),
	}
}
