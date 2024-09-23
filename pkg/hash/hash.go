package hash

import "golang.org/x/crypto/bcrypt"

type Hash interface {
	GenerateHashFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type PasswordHasher struct {
	cost int
}

func NewHash() *PasswordHasher {
	return &PasswordHasher{cost: bcrypt.DefaultCost}
}

func (h *PasswordHasher) GenerateHashFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *PasswordHasher) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
