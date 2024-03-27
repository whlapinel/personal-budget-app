package main

import (
	"database/sql"
)

func dropSprocs() error {
	db := initializeDB()
	defer db.Close()
	_, err := dropSprocUpdateAccountBalance(db)
	if err != nil {
		return err
	}
	return nil
}

func createSprocs() error {
	db := initializeDB()
	defer db.Close()
	_, err := createSprocUpdateAccountBalance(db)
	if err != nil {
		return err
	}
	_, err = createSprocUpdateCategoryBalance(db)
	if err != nil {
		return err
	}
	return nil
}

func dropSprocUpdateAccountBalance(db *sql.DB) (sql.Result, error) {
	query :=
		`
		DROP PROCEDURE IF EXISTS update_account_balance;
		`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func createSprocUpdateAccountBalance(db *sql.DB) (sql.Result, error) {
	// when new transaction is added, update account balance accordingly
	query :=
		`
		CREATE PROCEDURE update_account_balance(
			IN account_id INT,
			IN amount INT
		)
		BEGIN
			DECLARE new_balance INT DEFAULT 0;
			DECLARE old_balance INT DEFAULT 0;
			SELECT balance INTO old_balance FROM accounts WHERE id = account_id;
			SET new_balance = old_balance + amount;
			UPDATE accounts SET balance = new_balance WHERE id = account_id;
		END;
		`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func createSprocUpdateCategoryBalance(db *sql.DB) (sql.Result, error) {
	// when 1. assignments for a category are updated  or new transactions 
	// for a category are added/removed/changed, this sproc should be called to 
	// update category balance accordingly

	query :=
		`
		CREATE PROCEDURE update_category_balance(
			IN category_id INT,
			IN amount INT
		)
		BEGIN
			DECLARE new_balance INT DEFAULT 0;
			DECLARE old_balance INT DEFAULT 0;
			SELECT balance INTO old_balance FROM categories WHERE id = category_id;
			SET new_balance = old_balance + amount;
			UPDATE categories SET balance = new_balance WHERE id = category_id;
		END;
		`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
