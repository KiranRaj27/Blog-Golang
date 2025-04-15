// internal/repository/user/pg_user.go
package user

import (
	"github.com/kiranraj27/blog-golang/internal/config"
	domain "github.com/kiranraj27/blog-golang/internal/domain/user"
)

type pgRepo struct{}

func NewUserRepo() domain.Repository {
	return &pgRepo{}
}

func (r *pgRepo) Create(u *domain.User) error {
	return config.DB.Create(u).Error
}

func (r *pgRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
