package internal

import (
	"Groupie-tracker/internal/api"
	"Groupie-tracker/internal/controllers"
	"net/http"
)

func Init_routes(server *http.ServeMux) {

	server.HandleFunc("/register", controllers.Register)
	server.HandleFunc("/login", controllers.Login)

<<<<<<<<< Temporary merge branch 1
	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../page"))))
=========
	server.HandleFunc("/api", api.CreateAPI)

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../assets"))))

>>>>>>>>> Temporary merge branch 2
}
