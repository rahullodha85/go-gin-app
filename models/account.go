package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	Email     string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	FirstName string    `gorm:"not null" json:"first-name"`
	LastName  string    `gorm:"not null" json:"last-name"`
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

type AccountEp struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first-name"`
	LastName  string    `json:"last-name"`
	CreatedAt time.Time	`json:"created-at"`
	UpdatedAt time.Time	`json:"updated-at"`
}
