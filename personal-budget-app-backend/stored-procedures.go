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
			IN account_id INT
		)
		BEGIN
			DECLARE new_balance INT DEFAULT 0;
			DECLARE starting_balance INT;
		
			-- Retrieve the total amount from transactions for the given account_id
			SELECT SUM(amount) INTO new_balance
			FROM transactions
			WHERE account_id = account_id;
		
			-- Retrieve the starting balance for the account
			SELECT balance INTO starting_balance
			FROM accounts
			WHERE id = account_id;
		
			-- Calculate the new balance by adding the sum of transactions to the starting balance
			SET new_balance = new_balance + starting_balance;
		
			-- Update the accounts table with the new balance
			UPDATE accounts
			SET balance = new_balance
			WHERE id = account_id;
		END;
		`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
