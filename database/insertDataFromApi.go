package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Group struct {
	GroupID      int      `json:"id"`
	GroupName    string   `json:"name"`
	Image        string   `json:"image"`   // URL
	Members      []string `json:"members"` //Pas dans db
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	// API rest link
}

type Artist struct {
	ArtistID     int
	ArtistName   string
	GroupTableID int //Je sais pas si c'est utile
}

type Dates struct {
	DatesID      int    `json:"id"`
	GroupTableID int    //Je sais pas si c'est utile
	Location     string `json:"location"`    // API rest link
	ConcertDate  string `json:"concertDate"` // API rest link
}

func InsertDataFromApi(db *sql.DB) {

	var artistsInfos []Group

	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/artists", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Printf("\nRequesting artists: %v\n", req) // debug

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	fmt.Printf("\nResponse: %v\n", resp) // debug

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&artistsInfos)
	if err != nil {
		fmt.Println("error decoding the file: ", err)
		return
	}

	for _, artist := range artistsInfos {
		fmt.Println(artist.GroupID, ".", artist.GroupName, "(", artist.CreationDate, ") - First album:", artist.FirstAlbum)
	}
	//

	//Idee de comment faire pour inserer les donnees dans la db en deux étapes, c'est une bonne solution

	//stmt, err := db.Prepare("INSERT INTO Artist (ID, ArtisteName, Image, Members, CreationDate, FirstAlbum, Location, ConcertDate, Relation) VALUES (?, ?, ?, ? ,? ,? ,? ,? ,?)")
	//if err != nil {
	//	print("Error while preparing the statement: ", err)
	//}
	//defer stmt.Close()
	//
	//for _, d := range artistsInfos {
	//
	//	}
	//
	//	_, err = stmt.Exec(d.ID, d.Name, d.CreationDate, d.Members, d.CreationDate, d.FirstAlbum, d.Location, d.ConcertDate, d.Relations)
	//	if err != nil {
	//		print("Error while executing the statement: ", err)
	//	}
	//}
}
