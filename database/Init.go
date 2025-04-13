package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var once sync.Once

func Init() {
	log.Println("Init function called")
	once.Do(func() {
		var err error
		DB, err = sql.Open("sqlite3", "./lost&found.db")
		if err != nil {
			log.Fatal(err)
		}

		_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS laf (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			whofound TEXT NOT NULL,
			"where" TEXT NOT NULL,
			notes TEXT NOT NULL,
			date_added TEXT NOT NULL,
			date_claimed TEXT NOT NULL
		)`)
		if err != nil {
			log.Fatal(err)
		}
	})
}
