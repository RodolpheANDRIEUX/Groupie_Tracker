package internal

import (
	"Groupie-tracker/internal/controllers"
	"net/http"
)

func Init_routes(server *http.ServeMux) {
	server.HandleFunc("/register", controllers.Register)
	server.HandleFunc("/login", controllers.Login)
}
