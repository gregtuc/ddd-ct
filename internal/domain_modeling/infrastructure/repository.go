package infrastructure

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DomainModelRepository implements the repository interface using GORM.
type DomainModelRepository struct {
	db *gorm.DB
}

// NewDomainModelRepository creates a new repository instance.
func NewDomainModelRepository(db *gorm.DB) *DomainModelRepository {
	return &DomainModelRepository{db: db}
}

// InitDB initializes the database and migrates the DomainModel.
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("domainmodels.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.DomainModel{})
	return db, nil
}

// Create inserts a new DomainModel into the database.
func (r *DomainModelRepository) Create(model *domain.DomainModel) error {
	return r.db.Create(model).Error
}
