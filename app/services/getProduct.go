package services

import (
	"encoding/json"
	"errors"
	
	"leap_store/app/models"
	"net/http"
	//"golang.org/x/tools/go/analysis/passes/defers"
)

func GetAllProducts() ([]models.Product, error){
	url := "https://fakestoreapi.com/products"

	resp, err := http.Get(url)

	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to fetch products: " + resp.Status)
	}

	var products []models.Product

	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil

}