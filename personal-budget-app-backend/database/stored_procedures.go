package database

func CreateSprocs() error {
    _, err := ExecuteScript("create_sprocs.sql", ";;\n", "exec")
	return err
}
