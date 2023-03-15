package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// launch page/home.html

	htmlBytes, err := os.ReadFile("page/Home.html")
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier HTML" + err.Error())
		log.Printf("Erreur lors du chargement du fichier HTML\" %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	html := string(htmlBytes)

	// Écrire la réponse HTTP
	if _, err := fmt.Fprint(w, html); err != nil {
		log.Printf("Erreur lors de l'écriture de la réponse HTTP: %v", err)
	}
}
