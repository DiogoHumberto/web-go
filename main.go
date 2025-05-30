package main

import (
	"fmt"
	"net/http"

	"study.com/golang-web/config"
	"study.com/golang-web/db"
	"study.com/golang-web/routes"

	_ "github.com/lib/pq"
)

func main() {

	db.ConnectDataBase()

	config.InitializeStore()

	routes.UpAllRoutes()

	fmt.Println("Listening on :8080...")

	http.ListenAndServe(":8080", nil)
}
