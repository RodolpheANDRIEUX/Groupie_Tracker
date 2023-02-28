package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=username password=password dbname=gp-db port=6666 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"time"
//
//	_ "github.com/lib/pq"
//)
//
//func connect() (*sql.DB, error) {
//	bin, err := ioutil.ReadFile("/run/secrets/db-password")
//	if err != nil {
//		return nil, err
//	}
//	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
//}
//
//func blogHandler(w http.ResponseWriter, r *http.Request) {
//	db, err := connect()
//	if err != nil {
//		w.WriteHeader(500)
//		return
//	}
//	defer db.Close()
//
//	rows, err := db.Query("SELECT title FROM blog")
//	if err != nil {
//		w.WriteHeader(500)
//		return
//	}
//	var titles []string
//	for rows.Next() {
//		var title string
//		err = rows.Scan(&title)
//		titles = append(titles, title)
//	}
//	json.NewEncoder(w).Encode(titles)
//}
//
//func prepare() error {
//	db, err := connect()
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	for i := 0; i < 60; i++ {
//		if err := db.Ping(); err == nil {
//			break
//		}
//		time.Sleep(time.Second)
//	}
//
//	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
//		return err
//	}
//
//	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id SERIAL, title VARCHAR)"); err != nil {
//		return err
//	}
//
//	for i := 0; i < 5; i++ {
//		if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", fmt.Sprintf("Blog post #%d", i)); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func test() {
//	log.Print("Prepare db...")
//	if err := prepare(); err != nil {
//		log.Fatal(err)
//	}
//
//	log.Print("Listening 8000")
//	r := mux.NewRouter()
//	r.HandleFunc("/", blogHandler)
//	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
//}
