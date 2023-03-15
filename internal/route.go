package internal

import (
	"Groupie-tracker/internal/api"
	"Groupie-tracker/internal/controllers"
	"net/http"
)

func Init_routes(server *http.ServeMux) {

	server.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("page/Assets"))))

	server.HandleFunc("/register", controllers.Register)
	server.HandleFunc("/login", controllers.Login)
	server.HandleFunc("/api", api.CreateAPI)
	server.HandleFunc("/", controllers.Home)

}
