package controllers

import (
	"Groupie-tracker/internal/database"
	"database/sql"
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
	fileNameRegister := "register.html"
	t := template.Must(template.ParseFiles("page/" + fileNameRegister))

	if r.Method == http.MethodPost {
		username := r.FormValue("usernameInput")
		password, _ := HashPassword(r.FormValue("passwordInput"))

		var user User
		err := database.Database.Db.QueryRow("SELECT UserName FROM Users WHERE UserName=?", username).Scan(&user.Username)

		switch {
		case err == sql.ErrNoRows:
			_, err = database.Database.Db.Exec("INSERT INTO Users(UserName, Password) VALUES(?, ?)", username, password)
			if err != nil {
				println("Error 2")
			}
		case err != nil:
			panic(err)
		default:
			// Si l'utilisateur existe déjà, recharge la page avec un message d'erreur
			t.Execute(w, "L'utilisateur existe déjà. Veuillez choisir un autre nom d'utilisateur.")
			return
		}
	}

	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fileNameLogin := "login.html"
	t := template.Must(template.ParseFiles("page/" + fileNameLogin))

	if r.Method == http.MethodPost {
		username := r.FormValue("usernameInput")
		passwordInput := r.FormValue("passwordInput")

		var passwordDB string
		err := database.Database.Db.QueryRow("SELECT Password FROM Users WHERE UserName = ?", username).Scan(&passwordDB)
		if err != nil {
			panic(err.Error())
		}

		if CheckPasswordHash(passwordInput, passwordDB) {
			println("Password correct")
		} else {
			println("Password incorrect")
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
