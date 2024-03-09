package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
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
	router.Use(authenticateBFF)
	router.GET("/hello", sayHello)
	router.GET("/users/:email", getUserByEmail)
	router.GET("/categories", getCategories)
	router.POST("/users", postUser)
	router.POST("/signin", authenticateUser)
	router.POST("/categories", postCategory)
	router.GET("/transactions", getTransactions)
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

func getTransactions(c *gin.Context) {

}

func authenticateUser(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("creds", creds)
	db := initializeDB()
	defer db.Close()
	rows := db.QueryRow("SELECT password FROM users WHERE email = ?", creds.Email)
	var password string
	err := rows.Scan(&password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	if password != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "password does not match"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func postCategory(c *gin.Context) {
	var newCategory Category
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := newCategory.create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCategory)
}

func postUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newUser)
	if err := newUser.create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func getUserByEmail(c *gin.Context) {
	email := c.Param("email");
	fmt.Println(email)
	db := initializeDB()
	defer db.Close()
	fmt.Println("db initialized")
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func getCategories(c *gin.Context) {
	// authenticate
	var token string
	if token = c.GetHeader("Authorization"); token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no token"})
		return
	}
	fmt.Println("token", token)
	// validate token
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	// get categories
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories WHERE email = ?", claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
		return
	}
	var categories []Category
	var category Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.UserID, &category.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
			return
		} else {
			categories = append(categories, category)
		}
	}
	c.JSON(http.StatusOK, categories)
}
