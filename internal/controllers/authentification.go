package controllers

import (
	"Groupie-tracker/internal/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

var userValue User

type User struct {
	Username string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {

	htmlBytes, err := ioutil.ReadFile("page/register.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du fichier HTML", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	html := string(htmlBytes)

	// Écrire la réponse HTTP
	if _, err := fmt.Fprint(w, html); err != nil {
		log.Printf("Erreur lors de l'écriture de la réponse HTTP: %v", err)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("usernameInput")
		password, _ := HashPassword(r.FormValue("passwordInput"))

		var user int
		err := database.Database.QueryRow("SELECT COUNT(UserName) FROM Users WHERE UserName=?", username).Scan(&user)
		if err != nil {
			panic(err.Error())
		}

		if user > 0 {
			// Si l'utilisateur existe déjà, recharge la page avec un message d'erreur
			return
		} else {
			_, err = database.Database.Exec("INSERT INTO Users(UserName, Password) VALUES(?, ?)", username, password)
			session, _ := CookieStorage.Get(r, username)
			session.Values["Username"] = userValue.Username

		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	htmlBytes, err := ioutil.ReadFile("page/login.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du fichier HTML", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	html := string(htmlBytes)

	// Écrire la réponse HTTP
	if _, err := fmt.Fprint(w, html); err != nil {
		log.Printf("Erreur lors de l'écriture de la réponse HTTP: %v", err)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("usernameInput")
		passwordInput := r.FormValue("passwordInput")

		var passwordDB string
		err := database.Database.QueryRow("SELECT Password FROM Users WHERE UserName = ?", username).Scan(&passwordDB)
		if err != nil {
			panic(err.Error())
		}

		if CheckPasswordHash(passwordInput, passwordDB) {
			println("Password correct")
			session, _ := CookieStorage.Get(r, username)
			session.Values["Username"] = userValue.Username
			//todo : Redirect to next page
		} else {
			println("Password incorrect")
			//todo : Reload page with error message

		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
