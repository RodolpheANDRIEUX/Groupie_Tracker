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

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		os := runtime.GOOS
		_, err := w.Write([]byte("Hello from " + os + "!"))
		if err != nil {
			log.Fatal(err)
		}
	})

	database.Database()

	fmt.Print("(http://localhost:3000) Server started on port", port)
	http.ListenAndServe(port, server)

}
