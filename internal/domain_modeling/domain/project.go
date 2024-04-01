package domain

import "time"

// Project represents a collaboration project within the system.
type Project struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
