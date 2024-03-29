package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"

	"github.com/gin-gonic/gin"
)
func RegisterUserRoutes(router *gin.Engine) error {
	router.POST("/users", PostUser)
	router.GET("/users/:email", GetUserByEmail)
	return nil
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	fmt.Println(email)
	db := database.InitializeDB()
	defer db.Close()
	fmt.Println("db initialized")
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

