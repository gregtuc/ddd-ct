package domain

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Projects []Project `gorm:"many2many:user_projects;"`
}
