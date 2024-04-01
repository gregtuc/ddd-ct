package infrastructure

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
