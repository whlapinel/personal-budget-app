package main

import (
	"database/sql"
)

func dropTables(db *sql.DB) (sql.Result, error) {
	query :=
		`DROP TABLE if exists categories, users`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, err
}
