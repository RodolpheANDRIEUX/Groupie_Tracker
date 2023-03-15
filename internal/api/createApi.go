package api

import (
	"Groupie-tracker/internal/database"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func CreateAPI(w http.ResponseWriter, r *http.Request) {

	//CREATE ARTIST API

	// Get data from database
	rows, err := database.Database.Query("SELECT ArtistName, Image, FirstAlbum, SpotifyFollowers, CreationDate FROM Artist")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var artists []database.Artist
	for rows.Next() {
		var artist database.Artist
		if err := rows.Scan(&artist.ArtistName, &artist.Image, &artist.FirstAlbum, &artist.SpotifyFollowers.Total, &artist.CreationDate); err != nil {
			log.Fatal(err)
		}
		artists = append(artists, artist)
	}

	//

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
