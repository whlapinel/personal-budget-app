package database

import (
	"database/sql"
)

func DropTables()(sql.Result, error) {
	db := InitializeDB()
	defer db.Close()
	query :=
		`DROP TABLE if exists assignments, goals, transactions, categories, accounts, users`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, err
}
