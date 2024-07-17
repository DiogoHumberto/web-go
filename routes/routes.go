package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"study.com/golang-web/controllers"
)

var (
	store *sessions.CookieStore
	//users = map[string]string{} // map to store usernames and hashed passwords
)

func UpAllRoutes() {

	store = sessions.NewCookieStore([]byte(readKey()))

	routes := mux.NewRouter()

	routes.HandleFunc("/content", controllers.IndexContent)
	routes.HandleFunc("/about", controllers.AboutHandler)

	routes.HandleFunc("/", controllers.Index)
	routes.HandleFunc("/login", controllers.Login)
	routes.HandleFunc("/new", controllers.New)
	routes.HandleFunc("/insert", controllers.Insert)
	routes.HandleFunc("/delete", controllers.Delete)
	routes.HandleFunc("/edit", controllers.Edit)
	routes.HandleFunc("/update", controllers.Update)

	http.Handle("/", routes)
}

func readKey() string {
	// Carregar variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("SESSION_KEY")
	if key == "" {
		log.Fatal("SESSION_KEY environment variable not set")
	}
	return key
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
