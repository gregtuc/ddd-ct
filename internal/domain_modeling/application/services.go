package application

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/infrastructure"
)

// DomainModelService handles operations on domain models.
type DomainModelService struct {
	repo *infrastructure.DomainModelRepository
}

// NewDomainModelService creates a new service with the given repository.
func NewDomainModelService(repo *infrastructure.DomainModelRepository) *DomainModelService {
	return &DomainModelService{repo: repo}
}

// CreateDomainModel creates and persists a new domain model.
func (s *DomainModelService) CreateDomainModel(name, description string) (*domain.DomainModel, error) {
	model := &domain.DomainModel{Name: name, Description: description}
	err := s.repo.Create(model)
	return model, err
}
