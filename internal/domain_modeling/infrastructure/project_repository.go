package infrastructure

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"gorm.io/gorm"
)

// ProjectRepository interface for project data access abstraction.
type ProjectRepository interface {
	Create(project *domain.Project) error
	FindAll() ([]domain.Project, error)
	Delete(id uint) error
}

type gormProjectRepository struct {
	db *gorm.DB
}

// NewProjectRepository creates a new instance of a GORM-based project repository.
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &gormProjectRepository{db: db}
}

// Create adds a new project to the database.
func (r *gormProjectRepository) Create(project *domain.Project) error {
	return r.db.Create(project).Error
}

// FindAll retrieves all projects from the database.
func (r *gormProjectRepository) FindAll() ([]domain.Project, error) {
	var projects []domain.Project
	result := r.db.Find(&projects)
	return projects, result.Error
}

// Delete removes a project by its ID.
func (r *gormProjectRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Project{}, id).Error
}
