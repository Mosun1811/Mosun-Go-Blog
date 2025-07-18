package models

import (
	"time"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string
	Content   string `gorm:"type:text"`
	AuthorID  uuid.UUID
	Author    User
	CreatedAt time.Time
	UpdatedAt time.Time
}
