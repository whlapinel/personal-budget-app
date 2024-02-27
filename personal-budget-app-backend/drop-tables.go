package main

import (
	"database/sql"
	"log"
)

func dropTables(db *sql.DB) (sql.Result, error) {
	query :=
		`DROP TABLE if exists users, categories`
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
