package database

import (
	"database/sql"
)

func CreateAssignmentsTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE assignments (
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			category_id int,
			month VARCHAR(100),
			year int,
			amount int,
			FOREIGN KEY (email) REFERENCES users(email),
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func CreateUserTable(db *sql.DB) (sql.Result, error) {

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

func CreateTransactionTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE transactions (
			id int AUTO_INCREMENT PRIMARY KEY,
			account_id int,
			date datetime,
			payee VARCHAR(100),
			amount int,
			memo VARCHAR(100),
			category_id int,
			email VARCHAR(100),
			FOREIGN KEY (email) REFERENCES users(email),
			FOREIGN KEY (account_id) REFERENCES accounts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func CreateGoalsTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE goals (
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			amount int,
			target_date datetime,
			category_id int,
			periodicity VARCHAR(100),
			FOREIGN KEY (email) REFERENCES users(email),
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func CreateAccountTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE accounts (
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			type VARCHAR(100),
			bank_name VARCHAR(100),
			starting_balance int,
			balance int,
			FOREIGN KEY (email) REFERENCES users(email)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func CreateCategoryTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE categories (		
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			balance int,
			FOREIGN KEY (email) REFERENCES users(email)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
