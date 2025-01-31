package models

import "gorm.io/gorm"

type Kun struct{
	gorm.Model
	//ID int `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"` 
}

func (Kun) TableName() string {
	return "kun"
}