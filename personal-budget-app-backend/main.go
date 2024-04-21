package main

import (
	"fmt"
	"log"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/middleware"
	"personal-budget-app-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
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
		err := database.CreateDummyData()
		if err != nil {
			log.Fatal(err)
		}
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(middleware.AuthenticateBFF)
	routes.RegisterRoutes(router)
	router.Run("127.0.0.1:8080")
}
