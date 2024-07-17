package controllers

import (
	"log"
	"net/http"
	"strconv"

	"study.com/golang-web/db"
	"study.com/golang-web/models"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "content/about.html", nil)
}

func IndexContent(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "content/index.html", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	var produtcs []models.Product
	db.DB.Find(&produtcs)

	temp.ExecuteTemplate(w, "Index", produtcs)

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

		db.DB.Create(&models.Product{Name: name, Description: description, Price: priceConv, Amount: amountConv})
	}
	http.Redirect(w, r, "/", codeRedirect)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.DB.Delete(&models.Product{}, id)
	http.Redirect(w, r, "/", codeRedirect)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var product models.Product

	db.DB.First(&product, id)

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

		db.DB.Model(&models.Product{}).Where("id = ?", id).Updates(models.Product{Name: name, Description: description, Price: priceConv, Amount: amountConv})

	}
	http.Redirect(w, r, "/", codeRedirect)
}
