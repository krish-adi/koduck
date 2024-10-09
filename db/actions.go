package db

import (
	"fmt"
	"log"
)

func List() {
	rows, err := db.Query(`SELECT database_name, schema_name, table_name FROM duckdb_tables();`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var localTables, remoteTables []string

	for rows.Next() {
		var databaseName, schemaName, tableName string
		err = rows.Scan(&databaseName, &schemaName, &tableName)
		if err != nil {
			log.Fatal(err)
		}

		if databaseName == "koduck" && schemaName == "main" {
			localTables = append(localTables, tableName)
		} else if databaseName == "unstructured" && schemaName == "knowledge_bases" {
			remoteTables = append(remoteTables, tableName)
		}
	}

	if len(localTables) > 0 {
		fmt.Println("Local Knowledge Bases:")
		for _, table := range localTables {
			fmt.Printf("• %s\n", table)
		}
		fmt.Println("")
	}

	fmt.Println("Remote Knowledge Bases:")
	for _, table := range remoteTables {
		fmt.Printf("• %s\n", table)
	}
}

func Pull(tableName string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s.%s AS SELECT * FROM %s.%s.%s;",
		localDatabase, localSchema, tableName, remoteDatabase, remoteSchema, tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pulled: ", tableName)
}

func Drop(tableName string) {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s.%s", localDatabase, localSchema, tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dropped: ", tableName)
}
