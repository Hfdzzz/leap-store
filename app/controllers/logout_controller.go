package controllers

import "net/http"

func logoutHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}