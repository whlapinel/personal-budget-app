package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Category struct {
	ID    int     `json:"id"`
	Email string  `json:"email"`
	Name  string  `json:"name"`
	Goals *[]Goal `json:"goals"` // not stored in DB, but should be retrieved along with category
}

func (bc *Category) create() error {
	fmt.Println("Creating category")
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO categories (email, name) VALUES (?, ?)", bc.Email, bc.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getCategoriesByID(c *gin.Context) {
	var category Category
	// get categories
	id := c.Param("id")
	fmt.Println("id: ", id)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
		return
	}
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Email, &category.Name)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
			return
		}
	}
	c.JSON(http.StatusOK, category)
}

func getCategoriesByEmail(c *gin.Context) {
	var category Category
	fmt.Println("running getCategoriesByEmail")
	// get categories
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
		return
	}
	var categories []Category
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
	// get goals for each category
	for i, category := range categories {
		rows, err := db.Query("SELECT * FROM goals WHERE category_id = ?", category.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
			return
		}
		var goals []Goal
		for rows.Next() {
			var goal Goal
			var tempDate []uint8
			err := rows.Scan(&goal.ID, &goal.Email, &goal.Name, &goal.Amount, &tempDate, &goal.CategoryID, &goal.Periodicity)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
				return
			} else {
				goal.TargetDate, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing goal target date"})
					return
				}
				goals = append(goals, goal)
			}
		}
		categories[i].Goals = &goals
	}

	c.JSON(http.StatusOK, categories)
}

func postCategory(c *gin.Context) {
	var newCategory Category
	if err := c.BindJSON(&newCategory); err != nil {
		fmt.Println("error in c.BindJSON(&newCategory): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newCategory: ", newCategory)
	if err := newCategory.create(); err != nil {
		fmt.Println("error in newCategory.create(): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCategory)
}
