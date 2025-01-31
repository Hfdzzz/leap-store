package routes

import (
	"fmt"
	"leap_store/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() {

	indexController := controllers.IndexController{}
	
	registerController := controllers.RegisterController{}
	
	loginController := controllers.LoginController{}

	cartController := controllers.CartController{}

	productController := controllers.ProductController{}

	checkOutController := controllers.CheckOut_Controller{}

	
	r := mux.NewRouter()
	
	r.PathPrefix("/public/").Handler(
		http.StripPrefix("/public", http.FileServer(http.Dir("C:/xampp/htdocs/Belajar_Golang/Leap_Store/public"))),
	)
	
	r.HandleFunc("/register", registerController.RegisterPage).Methods("GET")
	r.HandleFunc("/register_account", registerController.RegisterHandler).Methods("POST")
	
	r.HandleFunc("/", loginController.LoginPage).Methods("GET")
	r.HandleFunc("/login_account", loginController.LoginHandler).Methods("POST")

	r.HandleFunc("/add_cart", cartController.CartHomeHandler).Methods("POST")
	
	r.HandleFunc("/home", indexController.Index).Methods("GET")

	r.HandleFunc("/input-product", productController.ProductHandler).Methods("POST")

	r.HandleFunc("/cart", cartController.CartPageHandler).Methods("GET")

	r.HandleFunc("/checkout", func (w http.ResponseWriter, r *http.Request) {

		checkOutController.SendEmail(w, r)

		checkOutController.CheckOutHandler(w, r)

		w.Write([]byte("Checkout succesfully"))
		
	}).Methods("POST")
	
	err := http.ListenAndServe(":8000", r)

	if err != nil {
		fmt.Println("Gagal menjalankan server")
	}
	
}
