package main

import (
	"database/sql"
	"log"
)

func createTables(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE users (
			id int AUTO_INCREMENT PRIMARY KEY, 
			first_name VARCHAR(100), 
			last_name VARCHAR(100), 
			email VARCHAR(100)
			)`
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	
	query =
		`CREATE TABLE categories (		
			id int AUTO_INCREMENT PRIMARY KEY,
			user_id int,
			name VARCHAR(100),
			needed DECIMAL(10,2),
			assigned DECIMAL(10,2),
			spent DECIMAL(10,2),
			available DECIMAL(10,2)
			)`
	result, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
