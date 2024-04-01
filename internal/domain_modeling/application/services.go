package application

import (
	"github.com/gtuc/ddd-ct/internal/domainmodeling/domain"
)

// DomainModelService handles operations on domain models.
type DomainModelService struct {
	// In a real application, this would interact with a repository
}

// NewDomainModel creates a new domain model with the given details.
func (s *DomainModelService) NewDomainModel(id, name, description string) domain.DomainModel {
	return domain.DomainModel{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
