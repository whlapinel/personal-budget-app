package database

import (
	"fmt"
)

func CreateDB() error {
	err := CreateTables()
	if err != nil {
		return err
	} else {
		fmt.Println("tables created")
	}
	err = CreateSprocs()
	if err != nil {
		return err
	} else {
		fmt.Println("stored procedures created")
	}
	err = CreateTriggers()
	if err != nil {
		return err
	} else {
		fmt.Println("triggers created")
	}
	return nil
}
