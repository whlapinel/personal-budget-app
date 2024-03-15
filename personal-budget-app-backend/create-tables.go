package main

import (
	"database/sql"
)

func createUserTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE users (
			email VARCHAR(100) PRIMARY KEY,
			first_name VARCHAR(100) NOT NULL, 
			last_name VARCHAR(100) NOT NULL,
			password VARCHAR(100) NOT NULL
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func createCategoryTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE categories (		
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			FOREIGN KEY (email) REFERENCES users(email)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
