package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"leap_store/app/configs"
	"leap_store/app/models"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

type CartData struct {
	Carts []models.Cart
	Username string
	Total_Price_of_All_Product float64
}

type CheckOut_Controller struct{}

func (cc *CheckOut_Controller) SendEmail(w http.ResponseWriter, r *http.Request){
	//nqpp rphz nawi ujbl

	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	fmt.Println("ini adalah username dari co controller ===  ",username)

	var cart []models.Cart

	err := configs.DB.Where("username = ?", username).Find(&cart).Error

	if err != nil {
		http.Error(w, "Failed to search data on cart", http.StatusSeeOther)
		fmt.Println("Failed to search data on cart")
		return
	}

	var Total_Price_of_All_Product float64

	err = configs.DB.Model(&cart).Where("username = ?", username).Select("SUM(total_price)").Scan(&Total_Price_of_All_Product).Error

	if err != nil {
		http.Error(w, "failed to sum price all product", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("views/template/email.tmpl")

	if err != nil {
		log.Fatalf("failed to load html: %v", err)
		fmt.Println("Failed to load email html")
		return
	}

	data := CartData{
		Carts: cart,
		Username: username,
		Total_Price_of_All_Product: Total_Price_of_All_Product,
	}

	var body bytes.Buffer

	err = tmpl.Execute(&body, data)

	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		fmt.Println("Failed to execute template")
		return
	}

	m := gomail.NewMessage()

	m.SetHeader("From", "watchamachalit@gmail.com")

	m.SetHeader("To", "hafidzaulia29@gmail.com")

	m.SetHeader("Subject", "percobaan ke 1")

	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "watchamachalit@gmail.com", "nqpprphznawiujbl")

	if err := d.DialAndSend(m); err != nil{
		//http.Redirect(w , r, "/cart", http.StatusSeeOther)
		fmt.Println("Faield to send email")
		return
	}

	fmt.Println("Success to send email")
	//http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (cc *CheckOut_Controller) CheckOutHandler(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "session")
	username := session.Values["username"]

	

	result := configs.DB.Unscoped().Where("username = ?", username).Delete(&models.Cart{})

	if result != nil {
		http.Error(w, "Failed to delete data", http.StatusInternalServerError)
        fmt.Println("Failed to delete data:", result.Error)
		//http.Redirect(w, r, "/home", http.StatusSeeOther)
		//fmt.Println("Failed to delete data")
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)

}