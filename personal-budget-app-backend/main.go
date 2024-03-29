package main

import (
	"personal-budget-app-backend/middleware"
	"personal-budget-app-backend/routes"

	"personal-budget-app-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
	wipeData := false // true if you want to delete tables and start over
	if wipeData {
		database.WipeData()
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(middleware.AuthenticateBFF)
	routes.RegisterRoutes(router)
	router.Run("localhost:8080")
}
