package database

import (
	"fmt"
)

func CreateDB() {
	err := CreateTables()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("tables created")
	}
	err = CreateSprocs()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("stored procedures created")
	}
}
