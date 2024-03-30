package database

import ()

func CreateTables() error {
    _, err := ExecuteScript("create_tables.sql", ";\n", "exec")
	return err
}
