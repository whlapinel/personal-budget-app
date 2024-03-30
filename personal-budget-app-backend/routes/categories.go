package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"

	"github.com/gin-gonic/gin"
)

func RegisterCategoriesRoutes(router *gin.Engine) error {
	router.GET("/categories/:email", GetCategories)
	router.POST("/categories", PostCategory)
	return nil
}

func GetCategories(c *gin.Context) {
	var category models.Category
	fmt.Println("running getCategoriesByEmail")
	// get categories
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
		return
	}
	var categories []models.Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Email, &category.Name)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
			return
		} else {
			categories = append(categories, category)
		}
	}
	c.JSON(http.StatusOK, categories)
}

func PostCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.BindJSON(&newCategory); err != nil {
		fmt.Println("error in c.BindJSON(&newCategory): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newCategory: ", newCategory)
	if err := newCategory.Save(); err != nil {
		fmt.Println("error in newCategory.Save(): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCategory)
}
