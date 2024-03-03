package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
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
	router.Use(authenticate)
	router.GET("/hello", sayHello)
	router.GET("/users/:id", getUserByID)
	router.GET("/categories", getCategories)
	router.POST("/signup", handleSignUp)
	router.POST("/signin", signIn)
	router.POST("/categories", handlePostCategory)
	router.GET("/transactions", getTransactions)
	router.Run("localhost:8080")
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func authenticate(c *gin.Context) {
	// authenticate
	var reqKey string
	if reqKey = c.GetHeader("API_KEY"); reqKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no key"})
		c.Abort()
		return
	}
	fmt.Println("API_KEY", reqKey)
	if reqKey != os.Getenv("API_KEY") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid key"})
		c.Abort()
		return
	}
	c.Next()
}



func getTransactions(c *gin.Context) {


}

func signIn(c *gin.Context) {
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
	// create token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//FIXME should be read from env
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating token"})
		return
	}
	// return token
	c.SetCookie("token", tokenString, 3600, "", "", false, false)
	// read the cookie that was just set
	c.JSON(http.StatusOK, tokenString)
}

func handlePostCategory(c *gin.Context) {
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
	// create category
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

func handleSignUp(c *gin.Context) {
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
