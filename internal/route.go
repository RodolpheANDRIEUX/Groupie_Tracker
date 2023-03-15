package internal

import (
	"Groupie-tracker/internal/api"
	"Groupie-tracker/internal/controllers"
	"net/http"
)

func Init_routes(server *http.ServeMux) {
	server.HandleFunc("/register", controllers.Register)
	server.HandleFunc("/login", controllers.Login)

	server.HandleFunc("/api", api.CreateAPI)

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../assets"))))

}
