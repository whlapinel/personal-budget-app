package main

import (
	"fmt"
)

func createTables() error {
	db := initializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	_, err := createUserTable(db)
	if err != nil {
		return err
	}
	_, err = createCategoryTable(db)
	if err != nil {
		return err
	}
	_, err = createAccountTable(db)
	if err != nil {
		return err
	}
	_, err = createTransactionTable(db)
	if err != nil {
		return err
	}
	_, err = createAssignmentsTable(db)
	if err != nil {
		return err
	}
	_, err = createGoalsTable(db)
	if err != nil {
		return err
	}
	return nil
}






