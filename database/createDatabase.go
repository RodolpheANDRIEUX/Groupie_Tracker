package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DatabaseInstence struct {
	Db *sql.DB
}

var DataBase DatabaseInstence

func Database() {

	//Database IP :172.20.0.2
	DB, err := sql.Open("mysql", "username:password@tcp(mysql_server:3306)/gp-db")
	if err != nil {
		fmt.Println("Can't connect to the DataBase: gp-db")
	}
	fmt.Println("Connexion to gp-db : Success!")

	//CREATION DES TABLES
	//CREATION OF GROUP TABLE

	queryGroupTable := `CREATE TABLE IF NOT EXISTS Band(
        GroupID int  PRIMARY KEY,
        GroupName VARCHAR(255),
    	Image VARCHAR(255),
    	FirstAlbum VARCHAR(255),
    	CreationDate int,
        Relation VARCHAR(255)
    );`
	// @TODO : essayer de lier les tables avec vue ou voir comment pour les groupe

	// exécution de la requête SQL
	_, err = DB.Query(queryGroupTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table GROUP created (if not exists)")

	//CREATION OF ARTIST TABLE

	// requête SQL pour créer la table si elle n'existe pas
	queryArtisteTable := `CREATE TABLE IF NOT EXISTS Artist (
		ArtistID int  PRIMARY KEY, 
		GroupeTableID int,
        ArtisteName VARCHAR(255),
		FOREIGN KEY (GroupeTableID) REFERENCES GroupID(GroupID)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryArtisteTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table ARTIST created (if not exists)")

	//CREATION OF DATES TABLE

	// requête SQL pour créer la table si elle n'existe pas
	queryDatesTable := `CREATE TABLE IF NOT EXISTS Dates (
        DatesID int  PRIMARY KEY,
        GroupeTableID int,
    	Location VARCHAR(255),
	    ConcertDate VARCHAR(255),
		FOREIGN KEY (GroupeTableID) REFERENCES GroupID(GroupID)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryDatesTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table ARTIST created (if not exists)")

	//ADD DATA FROM API
	InsertDataFromApi(DB)

	//GOOD PRACTICE
	DataBase = DatabaseInstence{Db: DB}
	DB.SetConnMaxLifetime(time.Minute * 100)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
