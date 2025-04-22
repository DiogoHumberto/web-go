package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"study.com/golang-web/config"
	"study.com/golang-web/controllers"
)

func UpAllRoutes() {

	routes := mux.NewRouter()

	// Rotas públicas
	routes.HandleFunc("/", controllers.Index)
	routes.HandleFunc("/login", controllers.Login)
	routes.HandleFunc("/user/create", controllers.Create)

	// Rotas protegidas
	protected := routes.PathPrefix("/").Subrouter()
	protected.Use(config.AuthMiddleware)

	protected.HandleFunc("/content", controllers.IndexContent)
	//routes.HandleFunc("/content", controllers.IndexContent)
	protected.HandleFunc("/about", controllers.AboutHandler)

	//routes.HandleFunc("/login", controllers.Login)
	protected.HandleFunc("/new", controllers.New)
	protected.HandleFunc("/insert", controllers.Insert)
	protected.HandleFunc("/delete", controllers.Delete)
	protected.HandleFunc("/edit", controllers.Edit)
	protected.HandleFunc("/update", controllers.Update)

	http.Handle("/", routes)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
