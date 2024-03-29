package main

import (
	"personal-budget-app-backend/middleware"
	"personal-budget-app-backend/routes"

	"personal-budget-app-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
	createDB := false // true if you want to delete database and start over
	if createDB {
		database.CreateDB()
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(middleware.AuthenticateBFF)
	routes.RegisterRoutes(router)
	router.Run("localhost:8080")
}
