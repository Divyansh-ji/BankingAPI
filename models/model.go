package models

type Account struct {
	UserID  uint    `gorm:"primary key"`
	Owner   string  `gorm:"size:100;not null"`
	Balance float64 `gorm:"not null;default:0"`
}
