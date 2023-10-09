package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Title    string
	Director string
}
