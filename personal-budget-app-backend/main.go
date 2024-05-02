package main

import (
	"fmt"
	"log"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/database/dummy_data"
	"personal-budget-app-backend/middleware"
	"personal-budget-app-backend/routes"
	"personal-budget-app-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// for development only
	// load .env file
	err := godotenv.Load(".env.backend")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	createDB := true // true if you want to delete database and start over
	if createDB {
		fmt.Println("creating database")
		err := database.CreateDB()
		if err != nil {
			log.Fatal(err)
		}
	}
	createDummyData := true // true if you want to create dummy data
	if createDummyData {
		err := dummy_data.CreateDummyData()
		if err != nil {
			log.Fatal(err)
		}
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(middleware.AuthenticateBFF)
	routes.RegisterRoutes(router)
	router.Run(util.GetHost() + ":8080")
}
