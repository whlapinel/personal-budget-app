package database

import (
	"fmt"
	"log"
	"os"
)

func CreateTables() error {
	db := InitializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	content, err := os.ReadFile("../SQL scripts/create-tables.sql")
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