package internal

import (
	"Groupie-tracker/internal/api"
	"Groupie-tracker/internal/controllers"
	"net/http"
)

func Init_routes(server *http.ServeMux) {

	server.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("page/Assets"))))

	http.HandleFunc("/", controllers.Home)

	server.HandleFunc("/auth", controllers.Authentification)
	server.HandleFunc("/register", controllers.Authentification)
	server.HandleFunc("/login", controllers.Authentification)

	server.HandleFunc("/api", api.CreateAPI)

}
