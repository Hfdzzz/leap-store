package main

import (
	//"fmt"
	//"leap_store/app/configs"
	//"leap_store/app/models"
	"leap_store/cmd/cli"
	//"leap_store/routes"
	//"log"
	//"net/http"
)

func main() {


	cli.Execute()

	// configs.ConnectDB()
	
	// log.Println("Berhasil melewati konek db")

	
	// configs.DB.AutoMigrate(&models.Users{},&models.Kun{})
	
	// log.Println("Berhasil melewati auto migrate")
	
	// if err := configs.DB.AutoMigrate(&models.Users{},&models.Kun{}); err != nil{
	// 	log.Fatalf("Gagal buat table: %v", err)
	// 	fmt.Println("gagal buat table")
	// 	} 
		
	// 	log.Println("Berhasil")
		
	// routes.RegisterRoutes()
	// 	log.Println("Berhasil melewati register route")
	// //http.ListenAndServe(":8000", nil)

	
}