package infrastructure

import (
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/domain"
	"gorm.io/driver/mysql"
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

func InitDB() (*gorm.DB, error) {
	// Update "username", "password", "dbname", and "parameters" accordingly.
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
