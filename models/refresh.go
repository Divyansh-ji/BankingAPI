package models

import "time"

type RefreshToken struct {
	ID        uint `gorm : "Primary key"`
	Token     string
	UserID    uint
	ExpiresAt time.Time
}
