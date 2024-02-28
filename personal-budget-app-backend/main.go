package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// var result sql.Result
	// log DBUSER and DBPASS
	dbUser := os.Getenv("DBUSER")
	fmt.Println("dbUser: " + dbUser)
	dbPass := os.Getenv("DBPASS")
	fmt.Println("dBPass: " + dbPass)
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	// for development only
	dropTables(db)
	fmt.Println("tables dropped")
	createUserTable(db)
	fmt.Println("user table created")
	createCategoryTable(db)
	fmt.Println("category table created")
	// API
	router := gin.Default()
	router.GET("/users/:id", getUserByID)
	// router.GET("/categories/:userID", getCategories2)
	router.POST("/signup", signup)
	router.POST("/categories", createCategory)

	router.Run("localhost:8080")
}

func createCategory(c *gin.Context) {
	var newCategory BudgetCategory
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := initializeDB()
	defer db.Close()
	category, err := addCategoryToDB(db, newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)

}

func signup(c *gin.Context) {
	fmt.Println(c.Request.Body)
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newUser)
	db := initializeDB()
	defer db.Close()
	user, err := createUser(db, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
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
	defer db.Close()
	fmt.Println("db initialized")
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// func getCategories2(c *gin.Context) {
// 	userIDString := c.Param("userID")
// 	userID, err := strconv.Atoi(userIDString)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "userID must be an integer"})
// 		return
// 	}
// 	var categories = new([]BudgetCategory)
// 	db := initializeDB()
// 	fmt.Println("db initialized")
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM budget_categories WHERE user_id = ?", userID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "categories not found"})
// 		return
// 	}
// 	for rows.Next() {
// 		var category BudgetCategory
// 		err := rows.Scan(&category.ID, &category

// 	for _, category := range budgetCategories {

// 		if category.UserID == userID {
// 			*categories = append(*categories, category)
// 		}
// 	}
// 	if len(*categories) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "categories not found"})
// 	} else if len(*categories) > 0 {
// 		c.JSON(http.StatusOK, categories)
// 	}
// }
