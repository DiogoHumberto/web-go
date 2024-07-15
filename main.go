package main

import (
	"net/http"

	"study.com/golang-web/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.UpAllRoutes()
	http.ListenAndServe(":8000", nil)
}
