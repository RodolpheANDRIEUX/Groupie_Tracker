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
	DB, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:6666)/gp-db")
	if err != nil {
		fmt.Println("Can't connect to the DataBase: gp-db")
	}
	fmt.Println("Connexion to gp-db : Success!")

	//@TODO : add var for collums

	// requête SQL pour créer la table si elle n'existe pas
	queryArtisteTable := `CREATE TABLE IF NOT EXISTS Artist (
        ID int AUTO_INCREMENT PRIMARY KEY,
        ArtisteName VARCHAR(100),
        Image VARCHAR(100),
		GroupeID int,
        CreationDate int,
        FirstAlbum VARCHAR(100),
        Location VARCHAR(100),
        ConcertDate VARCHAR(100),
        Relation VARCHAR(100)
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryArtisteTable)
	fmt.Println("Table ARTIST created (if not exists)")
	if err != nil {
		panic(err.Error())
	}

	// @TODO : essayer de lier les tables avec vue ou voir comment pour les groupes
	queryGroupTable := `CREATE TABLE IF NOT EXISTS GroupID (
        GroupeID int AUTO_INCREMENT FOREIGN KEY,
        GroupeName VARCHAR(100),
    	Artistname1 int,
    	Artistname2 int,
		Artistname3 int,
		Artistname4 int,
		Artistname5 int,
    );`

	// exécution de la requête SQL
	_, err = DB.Query(queryGroupTable)
	fmt.Println("Table GROUP created (if not exists)")
	if err != nil {
		panic(err.Error())
	}

	DataBase = DatabaseInstence{Db: DB}
	DB.SetConnMaxLifetime(time.Minute * 100)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
