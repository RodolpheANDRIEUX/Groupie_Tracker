package main

import (
	"Groupie-tracker/database"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {

	const port = ":3000" //le port exposé n'était pas le bon

	// @todo : utiliser os.env pour open la db

	server := http.NewServeMux()

	// display the OS to see if it's the docker container who's running the app
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		os := runtime.GOOS
		_, err := w.Write([]byte("Hello from " + os + "!"))
		if err != nil {
			log.Fatal(err)
		}
	})

	//CREATE DATABASE
	DB := database.CreateDatabase()

	//ADD DATA IN IT FROM APIs
	database.PopulateDatabase(DB)
	//api.GetToken()

	fmt.Print("(http://localhost:3000) Server started on port", port)
	log.Fatal(http.ListenAndServe(port, server))

}
