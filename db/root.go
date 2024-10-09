package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krish-adi/koduck/paths"
	_ "github.com/marcboeker/go-duckdb"
)

var db *sql.DB

// Local knowledge base variables
var localDatabase = "koduck"
var localSchema = "main"

// Remote knowledge base variables
var remoteDatabase = "unstructured"
var remoteSchema = "knowledge_bases"

func InitDB() {
	var err error

	db, err = sql.Open("duckdb", fmt.Sprintf("%s/koduck.db", paths.KoduckDir))
	if err != nil {
		log.Fatal(err)
	}

	KNOWLEDGE_BASE_IN_USE = ""
}

func AttachMD() {
	_, err := db.Exec("ATTACH 'md:unstructured';")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db.Close()
}
