// internal/usecase/user/service.go
package user

import (
	"errors"
	"strings"

	domain "github.com/kiranraj27/blog-golang/internal/domain/user"
	"github.com/kiranraj27/blog-golang/pkg/hash"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Register(u *domain.User) error {
	// basic validation
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	existing, _ := s.repo.FindByEmail(u.Email)
	if existing.ID != 0 {
		return errors.New("email already in use")
	}
	hashed, err := hash.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return s.repo.Create(u)
}

func (s *Service) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
