package api

import (
	"Groupie-tracker/internal/database"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
)

func CreateAPI(w http.ResponseWriter, r *http.Request) {

	query := `
	SELECT a.ArtistID, a.ArtistName, a.Image, a.FirstAlbum, a.CreationDate,
	       GROUP_CONCAT(DISTINCT m.MemberName) as MemberNames,
	       GROUP_CONCAT(DISTINCT d.ConcertDate) as ConcertDates,
	       GROUP_CONCAT( d.ConcertLocation) as ConcertLocations
	FROM Artist a
	JOIN Members m ON a.ArtistID = m.ArtistID
	JOIN Dates d ON a.ArtistID = d.ArtistID
	GROUP BY a.ArtistID
`

	rows, err := database.Database.Query(query)
	if err != nil {
		print("error")
		log.Fatal(err)
	}
	defer rows.Close()

	artistData := make([]database.Artist, 0)

	for rows.Next() {
		var artistInfo database.Artist
		var memberNames string
		var concertDates string
		var concertLocations string

		err = rows.Scan(&artistInfo.YtrackID, &artistInfo.ArtistName, &artistInfo.Image, &artistInfo.FirstAlbum, &artistInfo.CreationDate,
			&memberNames,
			&concertDates,
			&concertLocations)
		if err != nil {
			log.Fatal(err)
		}
		artistInfo.Members = strings.Split(memberNames, ",")
		artistInfo.ConcertDates = strings.Split(concertDates, ",")

		// Ajouter le code pour créer la structure DatesLocations
		concertLocationsSlice := strings.Split(concertLocations, ",")
		datesInfo := make(map[string][]string)
		for i, concertDate := range artistInfo.ConcertDates {
			if i < len(concertLocationsSlice) {
				location := concertLocationsSlice[i]
				datesInfo[location] = append(datesInfo[location], concertDate)
			} else {
				log.Printf("Erreur : L'index %d est en dehors des limites de concertLocationsSlice (longueur %d)\n", i, len(concertLocationsSlice))
			}
		}
		artistInfo.Dates = datesInfo

		artistData = append(artistData, artistInfo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artistData)

}
