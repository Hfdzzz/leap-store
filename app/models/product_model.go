package models

type Product struct{
	ID 			int `json:"id"`
	Title 		string `json:"title"`
	Price 		float64 `json:"price"`
	Category	string `json:"category"`
	Description string `json:"description"`
	Image 		string `json:"image"`
}

func (Product) TableName() string{
	return "product"
}