package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/users/:id", getUserByID)
	router.GET("/categories/:userID", getCategories)

	router.Run("localhost:8080")
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func getCategories(c *gin.Context) {
	userID := c.Param("userID")

	var categories = new([]BudgetCategory)
	for _, category := range budgetCategories {

		if category.UserID == userID {
			*categories = append(*categories, category)
		}
	}
	if len(*categories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "categories not found"})
	} else if len(*categories) > 0 {
		c.JSON(http.StatusOK, categories)
	}
}
