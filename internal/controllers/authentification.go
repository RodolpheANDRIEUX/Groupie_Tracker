package controllers

import (
	"Groupie-tracker/internal/database"
	"database/sql"
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var userValue User

func Register(w http.ResponseWriter, r *http.Request) {

	//TODO : hash the passowrd

	filename := "login.html"

	t := template.Must(template.ParseFiles("page/" + filename))

	username := r.FormValue("usernameInput")
	password := r.FormValue("passwordInput")

	if r.Method == http.MethodPost {

		var user string

		err := database.Database.Db.QueryRow("SELECT UserName FROM Users WHERE UserName=?", username).Scan(&user)
		switch {
		case err == sql.ErrNoRows:
			if err == nil {
				println("Error 1")
			}
			_, err = database.Database.Db.Exec("INSERT INTO Users(UserName, Password) VALUES(?, ?)", username, password)
			if err != nil {
				println("Error 2")
			}
			//CookieSession(w, r, username)
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
		case err != nil:
			println("Error 3")
		default:
			//CookieSession(w, r, username)
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
		}
	}
	t.Execute(w, nil)
}
