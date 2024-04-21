package database

import (
	"database/sql"
	"fmt"
	"log"
	"personal-budget-app-backend/secrets"

	"github.com/go-sql-driver/mysql"
)

func InitializeDB() *sql.DB {
	var db *sql.DB
	cfg := mysql.Config{
		User:                 secrets.DBUSER, // env.go, not in repo
		Passwd:               secrets.DBPASS, // env.go, not in repo
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "personal_budget",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("opened connection: ")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
