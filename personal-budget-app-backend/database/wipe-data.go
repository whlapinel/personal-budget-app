package database

import (
	"fmt"
)

func WipeData() {
	result, err := DropTables()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		fmt.Println("tables dropped")
	}
	err = DropSprocs()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("stored procedures dropped")
	}
	err = CreateTables()
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
