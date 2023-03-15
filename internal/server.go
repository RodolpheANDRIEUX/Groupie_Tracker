package internal

import (
	database2 "Groupie-tracker/internal/database"
	"fmt"
	"log"
	"net/http"
)

func Init_server() {
	//TODO

	const port = ":3000" //le port exposé n'était pas le bon

	server := http.NewServeMux()

	Init_routes(server)

	//CREATE DATABASE
	database2.CreateDatabase()

	//ADD DATA IN IT FROM APIs
	database2.PopulateDatabase()

	fmt.Print("(http://localhost:3000) Server started on port", port)
	log.Fatal(http.ListenAndServe(port, server))

}

func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/data")

}
