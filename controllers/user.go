package controllers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Login", nil)

}
