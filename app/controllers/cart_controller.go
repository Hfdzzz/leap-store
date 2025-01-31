package controllers

import (
	//"fmt"

	"fmt"
	"html/template"
	"leap_store/app/configs"
	"leap_store/app/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"
	//"strconv"
	//"text/template/parse"
)

type CartPageData struct{
	Users []models.Users
	Carts []models.Cart
	Total_Price_All_Product float64

}

type CartController struct{
	DB *gorm.DB
}

var username string

func (cc *CartController) CartPage(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("views/cartPage.tmpl"))
	tmpl.Execute(w, nil)
}

func (cc *CartController) CartHomeHandler(w http.ResponseWriter,  r *http.Request){
	r.ParseForm()

	session, _ := store.Get(r, "session")
	username = session.Values["username"].(string)

	name_product := r.FormValue("name_product")
	image := r.FormValue("image_product")
	
	priceSTR := r.FormValue("price")
	quantitySTR := r.FormValue("quantity")


	quantity, _ := strconv.ParseFloat(quantitySTR, 64)

	price, _ := strconv.ParseFloat(priceSTR, 64)

	total_price := price * quantity

	var existingCart models.Cart

	err := configs.DB.Where("name_product = ? AND username = ?", name_product, username).First(&existingCart).Error

	if err == nil {

		fmt.Println("Masuk kondisi baru")

		newQuantity := existingCart.Amount_Product + quantity 

		err := configs.DB.Model(&existingCart).Update("amount_product", newQuantity).Error

		if err != nil{
			http.Error(w, "Failed to update cart", http.StatusSeeOther)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		err = configs.DB.Model(&existingCart).Update("total_price", total_price).Error

		if err != nil {
			http.Error(w, "Failed to update total price", http.StatusSeeOther)
			http.Redirect(w, r,  "/home", http.StatusSeeOther)
			return
		}


		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	cart := models.Cart{
		Username: username,
		Name_Product: name_product,
		Image_Product: image,
		Price_Product: price,
		Amount_Product: quantity,
		Total_Price: total_price,

	}

	if err := configs.DB.Create(&cart).Error; err != nil{
		http.Error(w, "Failed to add to cart", http.StatusConflict)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
	return
}

func (cc *CartController) CartPageHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	var cart []models.Cart

	configs.DB.Where("username = ?", username).Find(&cart)

	var total_price_all_product float64

	configs.DB.Model(&cart).Where("username = ?", username).Select("SUM(total_price)").Scan(&total_price_all_product)

	tmpl, err := template.ParseFiles("views/cartPage.tmpl")

	if err != nil {
		http.Error(w, "Fail to load cart page", http.StatusSeeOther)
		return
	}

	data := CartPageData{
		Carts: cart,
		Total_Price_All_Product: total_price_all_product,
	}

	tmpl.Execute(w, data)
}