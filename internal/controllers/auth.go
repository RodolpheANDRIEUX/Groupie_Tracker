package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type User struct {
	Username string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckLoginStatus(w http.ResponseWriter, r *http.Request) {
	isLoggedIn := false // Remplacez par la logique pour vérifier l'état de connexion de l'utilisateur
	username := ""      // Remplacez par la logique pour obtenir le nom d'utilisateur de l'utilisateur connecté

	cookie, err := r.Cookie("username")
	if err == nil {
		isLoggedIn = true
		username = cookie.Value
	}

	w.Header().Set("X-Is-Logged-In", strconv.FormatBool(isLoggedIn))
	w.Header().Set("X-Username", username)
}
