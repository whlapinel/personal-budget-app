package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// for development only
	wipeData := false // true if you want to delete tables and start over
	if (wipeData) {
		result, err := dropTables()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
			fmt.Println("tables dropped")
		}
		err = createTables()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("tables created")
		}
	}
	// end development only
	// API
	router := gin.Default()
	router.Use(authenticateBFF)
	router.GET("/hello", sayHello)
	router.GET("/users/:email", getUserByEmail)
	router.POST("/users", postUser)
	router.GET("/categories/:email", getCategoriesByEmail)
	router.POST("/categories", postCategory)
	router.GET("/accounts/:email", getAccountsByEmail)
	router.POST("/accounts", postAccount)
	router.GET("/transactions/:accountID", getTransactionsByAccountID)
	router.POST("/transactions/", postTransaction)
	router.GET("/assignments/:email", getAssignmentsByEmail)
	router.POST("/assignments", postAssignment)
	router.GET("/goals/:email", getGoalsByEmail)
	router.POST("/goals", postGoal)
	router.Run("localhost:8080")
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func authenticateBFF(c *gin.Context) {
	// authenticate
	var reqKey string
	if reqKey = c.GetHeader("API_KEY"); reqKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no key"})
		c.Abort()
		return
	}
	fmt.Println("API_KEY", reqKey)
	if reqKey != API_KEY {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid key"})
		c.Abort()
		return
	}
	c.Next()
}