package controllers

import (
	//"database/sql"
	"html/template"
	"leap_store/app/configs"
	"leap_store/app/models"
	"net/http"

	"github.com/gorilla/sessions"
)

type LoginController struct{}

var store = sessions.NewCookieStore([]byte("secret-key"))

func (lc *LoginController) LoginPage(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("views/login.tmpl"))
	tmpl.Execute(w, nil)
}

func (lc *LoginController) LoginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.FormValue("username")
	passwordInput := r.FormValue("password")

	var user models.Users

	err := configs.DB.Model(&models.Users{}).Where("name = ?", username).First(&user).Error

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Error(w, "Invalid username", http.StatusUnauthorized)
	}

	if  passwordInput != user.Password {
		http.Error(w, "Invalid  password", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["username"] = user.Name
	session.Save(r, w)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
	

