package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct{
	gorm.Model
	Username				string `gorm:"type:varchar(255);not null"`
	Name_Product 			string `gorm:"type:varchar(255);not null"`
	Image_Product 			string `gorm:"type:varchar(255);not null"`
	Price_Product 			float64 `gorm:"type:DECIMAL(10,2);not null"`
	Total_Price 			float64 `gorm:"type:DECIMAL(10,2);not null"`
	Amount_Product 			float64 `gorm:"type:DECIMAL(10,2);not null"`
	created_at 				time.Time
	updated_at 				time.Time
	deleted_at 				time.Time
}

func (Cart) TableName() string {
	return "cart"
}