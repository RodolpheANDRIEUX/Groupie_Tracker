package main

import (
	"log"
	"net/http"
	"runtime"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		os := runtime.GOOS
		_, err := w.Write([]byte("Hello from " + os + "!"))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", server))
}
