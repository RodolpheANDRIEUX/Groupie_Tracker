package controllers

import (
	"Groupie-tracker/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		formType := r.FormValue("form_type")

		switch formType {
		case "login":
			username := r.FormValue("usernameLogin")
			password := r.FormValue("passwordLogin")

			var passwordDB string
			err := database.Database.QueryRow("SELECT Password FROM Users WHERE UserName = ?", username).Scan(&passwordDB)
			if err != nil {
				panic(err.Error())
			}

			if CheckPasswordHash(password, passwordDB) {
				println("Password correct")
				fmt.Println("Username: " + username)
				expiration := time.Now().Add(24 * time.Hour)
				cookie := http.Cookie{Name: "username", Value: username, Expires: expiration, Path: "/"}
				http.SetCookie(w, &cookie)
			} else {

				println("Password incorrect")
				//todo : Reload page with error message
				fmt.Printf("Login - Username: %s, Password: %s\n", username, password)
			}

		case "register":
			username := r.FormValue("usernameRegister")
			password, _ := HashPassword(r.FormValue("passwordRegister"))
			fmt.Printf("Register - Username: %s, Password: %s\n", username, password)

			var user int
			err := database.Database.QueryRow("SELECT COUNT(UserName) FROM Users WHERE UserName=?", username).Scan(&user)
			if err != nil {
				panic(err.Error())
			}

			if user > 0 {
				fmt.Println("User already exists")
			} else {
				_, err = database.Database.Exec("INSERT INTO Users(UserName, Password) VALUES(?, ?)", username, password)
				//session, _ := CookieStorage.Get(r, username)
				//session.Values["Username"] = userValue.Username
				fmt.Println("User created")
			}
		default:
			println("Error")
		}

	}

	htmlBytes, err := os.ReadFile("page/Home.html")
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier HTML" + err.Error())
		log.Printf("Erreur lors du chargement du fichier HTML\" %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	html := string(htmlBytes)

	if _, err := fmt.Fprint(w, html); err != nil {
		log.Printf("Erreur lors de l'écriture de la réponse HTTP: %v", err)
	}
}
