package main

import (
	"personal-budget-app-backend/middleware"
	"personal-budget-app-backend/routes"

	"personal-budget-app-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
	createDB := true // true if you want to delete database and start over
	if createDB {
		err := database.CreateDB()
		if err != nil {
			panic(err)
		}
	}
	createDummyData := true // true if you want to create dummy data
	if createDummyData {
		err := database.CreateDummyData()
		if err != nil {
			panic(err)
		}
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(middleware.AuthenticateBFF)
	routes.RegisterRoutes(router)
	router.Run("localhost:8080")
}
