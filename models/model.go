package models

type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Owner   string  `gorm:"size:100;not null"`
	Balance float64 `gorm:"not null;default:0"`
}
