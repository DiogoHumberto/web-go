package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"study.com/golang-web/config"
	"study.com/golang-web/db"
	"study.com/golang-web/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		authenticated, err := authenticateUser(email, senha)
		if err != nil || !authenticated {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := config.Store.Get(r, "session-name")
		session.Values["user"] = email
		session.Save(r, w)
		http.Redirect(w, r, "/new", http.StatusSeeOther)
		return
	}

	temp.ExecuteTemplate(w, "Login", nil)

}

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		err := createUser(username, email, senha)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	temp.ExecuteTemplate(w, "Create", nil)

}

func createUser(username, email, senha string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return db.DB.Create(&models.User{Username: username, Email: email, Password: string(hashedPassword)}).Error
}

// Verifica as credenciais do usuário
func authenticateUser(email, password string) (bool, error) {
	var user models.User

	result := db.DB.Where(&models.User{Email: email}).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil // Senha incorreta
	}

	return true, nil // Autenticação bem-sucedida
}
