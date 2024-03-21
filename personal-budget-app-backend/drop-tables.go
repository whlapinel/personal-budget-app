package main

import (
	"database/sql"
)

func dropTables() (sql.Result, error) {
	db := initializeDB()
	defer db.Close()
	query :=
		`DROP TABLE if exists assignments, goals, transactions, categories, accounts, users`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, err
}
