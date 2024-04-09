package database

import (
	"database/sql"
	"fmt"
	"os"
	"personal-budget-app-backend/util"
	"strings"
)

func ExecuteScript(fileName string, separator string, scriptType string) (*sql.Rows, error) {
	rootDir, err := util.FindRootPath()
	if err != nil {
		return nil, err
	}
	scriptDirectory := (rootDir + "/SQL_scripts/")
	db := InitializeDB()
	fmt.Println("db initialized")
	defer db.Close()
	content, err := os.ReadFile(scriptDirectory + fileName)
	if err != nil {
		return nil, err
	}
	// Split the SQL commands on semicolon followed by newline
	sqlCommands := strings.Split(string(content), separator)
	for _, command := range sqlCommands {
		if strings.TrimSpace(command) == "" {
			continue // skip empty commands resulting from the split
		}
		if scriptType == "exec" {
			fmt.Println("running SQL command: ", command)
			_, err = db.Exec(command)
			if err != nil {
				return nil, err
			}
		} else if scriptType == "query" {
			fmt.Println("running SQL query: ", command)
			rows, err := db.Query(command)
			if err != nil {
				return nil, err
			}
			return rows, nil
		}
	}
	return nil, nil
}
