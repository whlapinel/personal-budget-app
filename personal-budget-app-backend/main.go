package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
	wipeData := false // true if you want to delete tables and start over
	if wipeData {
		result, err := dropTables()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
			fmt.Println("tables dropped")
		}
		err = dropSprocs()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("stored procedures dropped")
		}
		err = createTables()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("tables created")
		}
		err = createSprocs()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("stored procedures created")
		}
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(authenticateBFF)

	// user routes
	router.GET("/users/:email", getUserByEmail)
	router.POST("/users", postUser)

	//categories routes
	router.GET("/categories/:email", getCategoriesByEmail)
	router.POST("/categories", postCategory)

	//accounts routes
	router.GET("/accounts/:email", getAccountsByEmail)
	router.POST("/accounts", postAccount)

	//transactions routes
	router.GET("/transactions/:email", getTransactionsByEmail)
	router.POST("/transactions/", postTransaction)

	//assignments routes
	router.GET("/assignments/:categoryID", getAssignmentsByCategoryID)
	router.GET("/assignments/email/:email", getAssignmentsByEmail)
	router.POST("/assignments", postAssignment)

	//goals routes... not sure this is needed since it's being sent as a property of categories
	router.GET("/goals/categoryID/:categoryID", getGoalByCategoryID)
	router.POST("/goals", postGoal)
	router.Run("localhost:8080")
}
