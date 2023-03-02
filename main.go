package main

import (
	"Groupie-tracker/database"
	"net/http"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	database.Database()

	http.ListenAndServe("localhost:8080", server)
}
