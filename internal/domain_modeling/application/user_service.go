package application

import (
	"errors"

	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/infrastructure"
)

type UserService struct {
	userRepo infrastructure.UserRepository
}

func NewUserService(userRepo infrastructure.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(name, email, password string) (*domain.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil && existingUser.ID != 0 {
		return nil, errors.New("email already in use")
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := s.userRepo.Create(user)
	return user, err
}
