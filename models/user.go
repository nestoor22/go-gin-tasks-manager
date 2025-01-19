package models

import (
	"time"
)
import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time // Auto-managed by GORM
	UpdatedAt time.Time // Auto-managed by GORM
}
