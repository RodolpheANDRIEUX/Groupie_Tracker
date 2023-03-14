package controllers

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var CookieStorage = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func CookieSession(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := CookieStorage.Get(r, "session-name")

	session.Values["Username"] = username // Set the different session values

	err := session.Save(r, w) // Save it before we write to the response/return from the handler.

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
