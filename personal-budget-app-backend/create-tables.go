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

func createAccountTable(db *sql.DB) (sql.Result, error) {
	
	query :=
		`CREATE TABLE accounts (
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			type VARCHAR(100),
			bank_name VARCHAR(100),
			starting_balance FLOAT,
			FOREIGN KEY (email) REFERENCES users(email)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func createTransactionTable(db *sql.DB) (sql.Result, error) {
	
	query :=
		`CREATE TABLE transactions (
			id int AUTO_INCREMENT PRIMARY KEY,
			account_id int,
			date datetime,
			payee VARCHAR(100),
			amount FLOAT,
			memo VARCHAR(100),
			category_id int,
			FOREIGN KEY (account_id) REFERENCES accounts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
