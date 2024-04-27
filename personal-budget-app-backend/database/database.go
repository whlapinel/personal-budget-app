package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB() *sql.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("MARIADB_USER"),
		Passwd:               os.Getenv("MARIADB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "mariadb:3306", // Change this to the service name and port in your docker-compose.yml
		DBName:               "personal_budget",
		AllowNativePasswords: true,
	}

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		fmt.Println("Connecting to the database...")
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("Failed to open database: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			db.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("Connected to database.")
		return db
	}

	log.Fatal("Failed to connect to the database after 5 attempts. Exiting...")
	return nil // This line will never be reached, but it's good practice to return a value on all code paths.
}
