package controllers

import (
	//"database/sql"
	"html/template"
	"leap_store/app/configs"
	"leap_store/app/models"
	"net/http"
	"strings"
	//"github.com/gorilla/sessions"
)

//var db *sql.DB

type RegisterController struct{}

func (rc *RegisterController) RegisterPage(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("views/register.tmpl"))
	tmpl.Execute(w, nil)
}

func (rc *RegisterController) RegisterHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	username := r.FormValue("username")

	if strings.Contains(username, " "){
		
	}

	if len(username) >= 7 {
		
	}

	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if password != confirmPassword {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user := models.Users{
		Name: username,
		Password: password,
	}

	

	 if err := configs.DB.Create(&user).Error; err != nil{
		http.Error(w, "failed to register, username already exist.", http.StatusConflict)
		return
	 }
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
	
}