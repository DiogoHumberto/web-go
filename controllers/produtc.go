package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"study.com/golang-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProdutcs := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", allProdutcs)

}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.InsertProduct(name, description, priceConv, amountConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.UpdateProduct(id, name, description, priceConv, amountConv)
	}
	http.Redirect(w, r, "/", 301)
}
