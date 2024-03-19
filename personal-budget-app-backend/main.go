package main

// for an explanation of this code see the Go tutorial "Developing a RESTful API with Go and Gin"
// https://go.dev/doc/tutorial/web-service-gin

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	// for development only
	// result, err := dropTables(db)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Println("tables dropped")
	// }
	// result, err = createUserTable(db)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Println("user table created")
	// }
	// result, err = createCategoryTable(db)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Println("category table created")
	// }
	// result, err = createAccountTable(db)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Println("account table created")
	// }
	// result, err = createTransactionTable(db)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Println("transaction table created")
	// }
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

func getTransactionsByAccountID(c *gin.Context) {
	var transaction Transaction
	// get transactions
	accountID := c.Param("accountID")
	fmt.Println("accountID: ", accountID)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM transactions WHERE account_id = ?", accountID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting transactions"})
		return
	}
	var transactions []Transaction
	for rows.Next() {
		var tempDate []uint8
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &tempDate, &transaction.Payee, &transaction.Amount, &transaction.Memo, &transaction.CategoryID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting transactions"})
			return
			} else {
				transaction.Date, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing transaction date"})
					return
				}
				transactions = append(transactions, transaction)
			}
	}
	c.JSON(http.StatusOK, transactions)
}

func getAccountsByEmail(c *gin.Context) {
	var account Account
	// get accounts
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting accounts"})
		return
	}
	var accounts []Account
	for rows.Next() {
		err := rows.Scan(&account.ID, &account.Email, &account.Name, &account.Type, &account.BankName, &account.StartingBalance)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting accounts"})
			return
		} else {
			accounts = append(accounts, account)
		}
	}

	// get account balances by retrieving sum of transactions for each account,
	// add to each account struct instance

	for i, account := range accounts {
		rows, err := db.Query("SELECT SUM(amount) FROM transactions WHERE account_id = ?", account.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
			return
		}
		var balance sql.NullFloat64
		for rows.Next() {
			err := rows.Scan(&balance)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
				return
			}
			if balance.Valid {
				accounts[i].Balance = balance.Float64 + account.StartingBalance
			} else {
				accounts[i].Balance = account.StartingBalance
			}
		}
	}
	c.JSON(http.StatusOK, accounts)
}

func postAccount(c *gin.Context) {
	var newAccount Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newAccount)
	if err := newAccount.create(); err != nil {
		fmt.Println("error in newAccount.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAccount)
}

func postTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newTransaction.AccountID", newTransaction.AccountID)
	if err := newTransaction.create(); err != nil {
		fmt.Println("error in newTransaction.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTransaction)
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
	email := c.Param("email")
	fmt.Println(email)
	db := initializeDB()
	defer db.Close()
	fmt.Println("db initialized")
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
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
	c.JSON(http.StatusOK, categories)
}
