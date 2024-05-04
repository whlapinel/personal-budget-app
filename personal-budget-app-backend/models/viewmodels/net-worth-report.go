package viewmodels

import (
	"fmt"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
	"time"
)

type NetWorthReportEntry struct {
	Email    string `json:"email"`
	Month    int    `json:"month"`
	Year     int    `json:"year"`
	NetWorth int    `json:"netWorth"`
}

func getAccountBalanceByDate(email string, account_id int, date time.Time) (*int, error) {
	// this will get the account balance for the last day of the month
	fmt.Println("getAccountBalanceByDate email: ", email)
	fmt.Println("getAccountBalanceByDate date: ", date)
	// convert date to last day of month
	lastDayOfMonth := date.AddDate(0, 0, -1)
	fmt.Println("getAccountBalanceByDate lastDayOfMonth: ", lastDayOfMonth)

	// get account balance for date

	var accountBalance int
	db := database.InitializeDB()
	defer db.Close()
	err := db.QueryRow("SELECT IFNULL(SUM(amount), 0) FROM transactions WHERE email = ? AND account_id = ? AND date <= ?", email, account_id, lastDayOfMonth).Scan(&accountBalance)
	if err != nil {
		return nil, err
	}

	// get account balances for last day of month
	return &accountBalance, nil
}

func getNetWorthByMonth(email string, month int, year int) (*int, error) {
	// this will get the net worth for the last day of the month
	fmt.Println("getNetWorthByMonth email: ", email)
	var netWorthForMonth int
	// loop through accounts
	accounts, err := models.GetAccounts(email)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		// get account balance for last day of month
		accountBalance, err := getAccountBalanceByDate(email, account.ID, time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local))
		if err != nil {
			return nil, err
		}
		// add account balance to net worth
		netWorthForMonth += *accountBalance
	}

	// return net worth
	fmt.Println("getNetWorthByMonth netWorthForMonth: ", netWorthForMonth)

	return &netWorthForMonth, nil
}

func GetNetWorthReport(email string) (*[]NetWorthReportEntry, error) {
	// get current month
	today := time.Now()
	thisMonth := int(today.Month())
	thisYear := today.Year()
	db := database.InitializeDB()
	defer db.Close()
	var netWorths []NetWorthReportEntry
	// get net worth for last 3 months
	for i := 2; i >= 0; i-- {
		currMonth := thisMonth - i
		currYear := thisYear
		if currMonth <= 0 {
			currMonth = 12 + currMonth
			currYear = thisYear - 1
		}
		// get net worth for month
		netWorth, err := getNetWorthByMonth(email, currMonth, currYear)
		if err != nil {
			return nil, err
		}
		// append entry to array
		netWorths = append(netWorths, NetWorthReportEntry{
			Email:    email,
			Month:    currMonth,
			Year:     currYear,
			NetWorth: *netWorth,
		})
	}
	return &netWorths, nil
}
