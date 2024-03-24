package main

import (
	"database/sql"
	"fmt"
)

func createTables() error {
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	_, err := createUserTable(db)
	if err != nil {
		return err
	}
	_, err = createCategoryTable(db)
	if err != nil {
		return err
	}
	_, err = createAccountTable(db)
	if err != nil {
		return err
	}
	_, err = createTransactionTable(db)
	if err != nil {
		return err
	}
	_, err = createAssignmentsTable(db)
	if err != nil {
		return err
	}
	_, err = createGoalsTable(db)
	if err != nil {
		return err
	}
	return nil
}

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

func createTransactionTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE transactions (
			id int AUTO_INCREMENT PRIMARY KEY,
			account_id int,
			date datetime,
			payee VARCHAR(100),
			amount int,
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

func createAssignmentsTable(db *sql.DB) (sql.Result, error) {

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

func createGoalsTable(db *sql.DB) (sql.Result, error) {

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
