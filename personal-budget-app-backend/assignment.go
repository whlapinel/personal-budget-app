package main

type Assignment struct {
	ID         int     `json:"id"`
	Email      string  `json:"email"`
	CategoryID int     `json:"category_id"`
	Month      string  `json:"month"`
	Year       string  `json:"year"`
	Amount     float64 `json:"amount"`
}

func (a *Assignment) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO assignments (email, category_id, month, year, amount) VALUES (?, ?, ?, ?, ?)", a.Email, a.CategoryID, a.Month, a.Year, a.Amount)
	if err != nil {
		return err
	}
	return nil
}
