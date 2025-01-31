package controllers

import (
	"leap_store/app/configs"
	"leap_store/app/models"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (pc *ProductController) ProductHandler(w http.ResponseWriter, r *http.Request){

	r.ParseForm()

	name_product := r.FormValue("name_product")
	description := r.FormValue("description")
	image := r.FormValue("image")
	category := r.FormValue("category")
	price := r.FormValue("price")

	priceFloat64, err := strconv.ParseFloat(price, 64)

	if err != nil {
		panic("fail to convert price to float")
	}

	product := models.Product{
		Title: name_product,
		Price: priceFloat64,
		Category: category,
		Description: description,
		Image: image,

	}

	err = configs.DB.Create(&product).Error

	if err != nil {
		panic("Failed to create data product")
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)




}