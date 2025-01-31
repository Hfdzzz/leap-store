package configs

import (
	"fmt"
	//"leap_store/app/configs"
	"leap_store/app/models"
	"log"

	//"time"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func generateCartData() models.Cart {
	price := rand.Float64()*9000 + 1000  // Harga antara 1000 - 10000
	amount := float64(rand.Intn(10) + 1) // Jumlah antara 1 - 10
	return models.Cart{
		Username:      faker.Username(),
		Name_Product:   faker.Word(),
		Image_Product:  faker.URL(),
		Price_Product:  price,
		Total_Price:    price * amount,
		Amount_Product: amount,
		
	}
}

// Insert Data Dummy ke PostgreSQL dalam batch
func InsertDummyData(db *gorm.DB, total int) {
	batchSize := 5000 // Batch insert 10 ribu data sekali jalan
	totalInserted := 0

	for totalInserted < total {
		var carts []models.Cart
		for i := 0; i < batchSize; i++ {
			carts = append(carts, generateCartData())
		}

		if err := DB.Create(&carts).Error; err != nil {
			log.Fatal("Gagal insert data:", err)
		}

		totalInserted += batchSize
		fmt.Printf("Inserted %d/%d rows...\n", totalInserted, total)
	}

	fmt.Println("✅ Selesai! 20 juta data telah dimasukkan.")
}

// func main() {
	
	
// 	insertDummyData(DB, 20000000) // Insert 20 juta data
// }