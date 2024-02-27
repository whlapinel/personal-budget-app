package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	// var result sql.Result
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	// for development only
	dropTables(db)
	fmt.Println("tables dropped")
	createTables(db)
	fmt.Println("tables created")
	seedUserTestData(db)
	fmt.Println("user test data seeded")
	seedCategoryTestData(db)
	fmt.Println("category test data seeded")

	// API
	router := gin.Default()
	router.GET("/users/:id", getUserByID)
	router.GET("/categories/:userID", getCategories2)

	router.Run("localhost:8080")
}

func getUserByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be an integer"})
		return
	}
	fmt.Println(id)
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	var user User
	//FIXME this returns 500 when user is not found instead of 404
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func getCategories2(c *gin.Context) {
	userIDString := c.Param("userID")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userID must be an integer"})
		return
	}
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
