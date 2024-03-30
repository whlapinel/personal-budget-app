package database

func CreateDummyData() error {
	_, err := ExecuteScript("add_dummy_data.sql", ";\n", "exec")
	return err
}
