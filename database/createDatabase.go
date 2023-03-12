package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func CreateDatabase() *sql.DB {

	//Database IP :172.20.0.2
	DB, err := sql.Open("mysql", "username:password@tcp(mysql_server:3306)/gp-db")
	if err != nil {
		fmt.Println("Can't connect to the DataBase: gp-db")
	}
	fmt.Println("Connexion to gp-db : Success!")

	//CREATION DES TABLES

	//CREATION OF ARTIST TABLE
	queryArtistTable := `CREATE TABLE IF NOT EXISTS Artist(
        ArtistID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        ArtistName VARCHAR(255),
    	Image VARCHAR(255),
    	FirstAlbum VARCHAR(255),
    	SpotifyFollowers int,
    	CreationDate DATE
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryArtistTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table ARTIST created (if not exists)")

	//CREATION OF MEMBERS TABLE
	queryMembersTable := `CREATE TABLE IF NOT EXISTS Members (
		MemberID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		MemberName VARCHAR(255),
        ArtistTableID int,
		FOREIGN KEY (ArtistTableID) REFERENCES Artist(ArtistID)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryMembersTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table MEMBERS created (if not exists)")

	//CREATION OF DATES TABLE
	queryDatesTable := `CREATE TABLE IF NOT EXISTS Dates (
        DatesID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        ArtistTableID int,
    	ConcertLocation VARCHAR(255),
	    ConcertDate VARCHAR(255),
		FOREIGN KEY (ArtistTableID) REFERENCES Artist(ArtistID)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryDatesTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table DATES created (if not exists)")

	//CREATION OF ALBUM TABLE
	queryAlbumTable := `CREATE TABLE IF NOT EXISTS Album (
        AlbumID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        ArtistTableID int,
    	AlbumName VARCHAR(255),
	    AlbumImage VARCHAR(255),
    	AlbumReleaseDate VARCHAR(255),
		FOREIGN KEY (ArtistTableID) REFERENCES Artist(ArtistID)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryAlbumTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table ALBUM created (if not exists)")

	//CREATION OF USERS TABLE
	queryUsersTable := `CREATE TABLE IF NOT EXISTS Users (
    UserID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    UserName varchar(255) NOT NULL UNIQUE KEY,
    Password varchar(255) NOT NULL
);`

	// exécution de la requête SQL
	_, err = DB.Query(queryUsersTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table USERS created (if not exists)")

	//GOOD PRACTICE
	DB.SetConnMaxLifetime(time.Minute * 100)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	return DB
}
