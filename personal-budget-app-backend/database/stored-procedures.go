package database

import (
	"log"
	"os"
)

func CreateSprocs() error {
	db := InitializeDB()
	defer db.Close()
	content, err := os.ReadFile("../SQL scripts/create-sprocs.sql")
    if err != nil {
        log.Fatal(err)
    }
    sqlCommands := string(content)

    // Execute SQL commands
    _, err = db.Exec(sqlCommands)
    if err != nil {
        log.Fatal(err)
    }
	return nil
}
