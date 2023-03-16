package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Artist struct {
	YtrackID         int    `json:"id"`
	ArtistName       string `json:"name"`
	Image            string `json:"image"` // URL
	SpotifyFollowers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	DatesAPI     string              `json:"relations"` // URL API
	Dates        map[string][]string `json:"datesLocations"`
	ConcertDates []string            `json:"concert_dates"`
}

type SearchResult struct {
	Artists struct {
		Items []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"artists"`
}

func GetSpotifyToken(artistName string, BearerToken string) string {

	// create a new HTTP client
	client := &http.Client{}
	// create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// set the query parameters
	q := url.Values{}
	q.Add("q", artistName)
	q.Add("type", "artist")
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()
	// add the authorization header
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	// make the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// decode the response JSON
	var result SearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// print the Artist ID
	if len(result.Artists.Items) > 0 {
		return result.Artists.Items[0].ID
	}
	return ""
}

func getSpotifyArtist(token string, BearerToken string) []byte {

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/artists/"+token, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", BearerToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error ReadAll the file: ", err)
		return nil
	}
	return body
}

func PopulateDatabase() {
	DB := Database
	var artists []Artist

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err = json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Println("Error:", err)
		return
	}

	BearerToken := "BQAyidSvyPBGNAoT64P7nlq_Q9OpRPFmKntsOriE0ssVv60Cs9sHu-Qc94Ldr5gaOcHY1vXwxEtoKG5o6zjW-Lw_VWabRPBQEhMFFlw0j3SCsyUQxLRuxHq_vgHcBQPcCxkyFZCVEJpKszXdHkao92Uer0hoeeEo8rnNo2kj-fzmdRYwItKjk87Vm9vBqACmW0WI9oPi"

	for _, artist := range artists {

		dateResp, err := http.Get(artist.DatesAPI)
		if err != nil {
			panic(err)
		}
		defer dateResp.Body.Close()

		err = json.NewDecoder(dateResp.Body).Decode(&artist)
		if err != nil {
			panic(err)
		}

		token := GetSpotifyToken(artist.ArtistName, BearerToken)
		spotifyBody := getSpotifyArtist(token, BearerToken)
		err = json.Unmarshal(spotifyBody, &artist)
		if err != nil {
			fmt.Println("error decoding the file: ", err)
			return
		}

		//fmt.Println(artist.ArtistName, "(", artist.CreationDate, ") - Members:", len(artist.Members), " - Followers:", artist.SpotifyFollowers.Total)
		SaveArtist(artist, DB)
	}
}

func SaveArtist(artist Artist, db *sql.DB) {

	qArtist, err := db.Prepare("INSERT IGNORE INTO Artist (ArtistName, Image, FirstAlbum, SpotifyFollowers, CreationDate) VALUES (?, ?, ?, ?, ? )")
	if err != nil {
		print("Error while preparing the statement1: ", err)
	}
	result, err := qArtist.Exec(artist.ArtistName, artist.Image, artist.FirstAlbum, artist.SpotifyFollowers.Total, artist.CreationDate)
	if err != nil {
		print("Error while executing the statement1: ", err)
	}

	artistID, err := result.LastInsertId()
	if err != nil {
		print("Error while getting last primary key: ", err)
	}

	for _, member := range artist.Members {
		qMembers, err := db.Prepare("INSERT IGNORE INTO Members (MemberName, ArtistID) VALUES (?, ?)")
		if err != nil {
			print("Error while preparing the statement2: ", err)
		}
		_, err = qMembers.Exec(member, artistID)
		if err != nil {
			print("Error while executing the statement2: ", err)
		}
	}

	for location, dates := range artist.Dates {
		for _, date := range dates {
			qDates, err := db.Prepare("INSERT IGNORE INTO Dates (ConcertLocation, ConcertDate, ArtistID) VALUES (?, ?, ?)")
			if err != nil {
				print("Error while preparing the statement3: ", err)
			}
			fmt.Println("attends encore")
			_, err = qDates.Exec(location, date, artistID)
			if err != nil {
				print("Error while executing the statement3: ", err)
			}
		}
	}
}
