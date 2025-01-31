package controllers

import (
	"html/template"
	"leap_store/app/configs"
	"leap_store/app/models"
	"leap_store/app/services"
	"net/http"
	//"github.com/gorilla/sessions"
)

type IndexPageData struct{
	Users []models.Users
	Product []models.Product
	//Cart []models.Cart
	Username string
	TotalAmount float64
}

type IndexController struct{}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	username, _ := session.Values["username"].(string)

	if username == ""{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var user []models.Users
	configs.DB.Find(&user)

	product, err := services.GetAllProducts()
	
	if err != nil{
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	var Cart []models.Cart

	var totalAmount float64

	configs.DB.Model(&Cart).Where("username = ?", username).Select("SUM(amount_product)").Scan(&totalAmount)

	data := IndexPageData{
		Users: user,
		Product: product,
		Username: username,
		TotalAmount: totalAmount,
	}

	tmpl,err := template.ParseFiles("views/index.tmpl")

	if err != nil {
		http.Error(w, "Failed to load index", http.StatusInternalServerError)
		//http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, data)
}