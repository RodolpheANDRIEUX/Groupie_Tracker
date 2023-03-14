package controllers

import (
	"Groupie-tracker/internal/database"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var userValue User

type User struct {
	Username string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./internal/page/register.html"))

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
			t.Execute(w, "L'utilisateur existe déjà. Veuillez choisir un autre nom d'utilisateur.")
			return
		} else {
			_, err = database.Database.Exec("INSERT INTO Users(UserName, Password) VALUES(?, ?)", username, password)
			session, _ := CookieStorage.Get(r, username)
			session.Values["Username"] = userValue.Username

		}
	}

	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("page/login.html"))

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

	t.Execute(w, nil)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
