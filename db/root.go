package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"

	"github.com/krish-adi/koduck/paths"
)

var db *sql.DB

func InitDB() {

	var err error

	db, err = sql.Open("duckdb", fmt.Sprintf("%s/koduck.db?access_mode=READ_WRITE", paths.KoduckDir))
	if err != nil {
		log.Fatal(err)
	}

}

func CloseDB() {
	db.Close()
}

func RunQuery(query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
