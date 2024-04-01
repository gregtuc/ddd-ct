package application

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/infrastructure"
)

// ProjectService provides operations for managing projects.
type ProjectService struct {
	repo infrastructure.ProjectRepository
}

// NewProjectService creates a new project service with the given repository.
func NewProjectService(repo infrastructure.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// CreateProject handles the creation of a new project.
func (s *ProjectService) CreateProject(name, description string) (*domain.Project, error) {
	project := &domain.Project{Name: name, Description: description}
	err := s.repo.Create(project)
	return project, err
}

// GetAllProjects retrieves all projects.
func (s *ProjectService) GetAllProjects() ([]domain.Project, error) {
	return s.repo.FindAll()
}

// DeleteProject handles the deletion of a project.
func (s *ProjectService) DeleteProject(id uint) error {
	return s.repo.Delete(id)
}
