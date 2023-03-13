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
	ArtistName       string `json:"name"`
	Image            string `json:"image"` // URL
	SpotifyFollowers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	DatesAPI     string   `json:"dates"` // API rest link
	Dates        []Dates
}

type Dates struct {
	ConcertLocation string `json:"location"`
	ConcertDate     string `json:"concertDate"`
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
	} else {
		fmt.Println("No Artist found")
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

func PopulateDatabase(DB *sql.DB) {
	var artists []Artist

	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/artists", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error ReadAll the file: ", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println("error unmarshal the file: ", err)
		return
	}

	BearerToken := "BQAyidSvyPBGNAoT64P7nlq_Q9OpRPFmKntsOriE0ssVv60Cs9sHu-Qc94Ldr5gaOcHY1vXwxEtoKG5o6zjW-Lw_VWabRPBQEhMFFlw0j3SCsyUQxLRuxHq_vgHcBQPcCxkyFZCVEJpKszXdHkao92Uer0hoeeEo8rnNo2kj-fzmdRYwItKjk87Vm9vBqACmW0WI9oPi"

	for _, artist := range artists {
		token := GetSpotifyToken(artist.ArtistName, BearerToken)
		spotifyBody := getSpotifyArtist(token, BearerToken)
		err = json.Unmarshal(spotifyBody, &artist)
		if err != nil {
			fmt.Println("error decoding the file: ", err)
			return
		}
		SaveArtist(artist, DB)
		//fmt.Println(artist.ArtistName, "(", artist.CreationDate, ") - Members:", len(artist.Members), " - Followers:", artist.SpotifyFollowers.Total)
	}
}

func SaveArtist(artist Artist, db *sql.DB) {

	stmt, err := db.Prepare("INSERT INTO Artist (ArtistName, Image, CreationDate, FirstAlbum ) VALUES (?, ?, ?, ? )")
	if err != nil {
		print("Error while preparing the statement: ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(artist.ArtistName, artist.Image, artist.CreationDate, artist.FirstAlbum)
	if err != nil {
		print("Error while executing the statement: ", err)
	}
}
