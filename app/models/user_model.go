package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	ID int `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"` 
	created_at time.Time
	updated_at time.Time
	deleted_at time.Time
}

func (Users) TableName() string {
	return "users"
}