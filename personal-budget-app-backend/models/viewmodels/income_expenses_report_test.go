package viewmodels

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetIncomeAndExpenses(t *testing.T) {
	err := godotenv.Load("../../.env.backend")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	err = godotenv.Load("../../.env.backend.testing")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// test GetIncomeAndExpenses
	email := "test@test.com"
	data, err := GetIncomeAndExpenses(email)
	if err != nil {
		t.Errorf("GetIncomeAndExpenses failed: %v", err)
	}
	if data == nil {
		t.Errorf("GetIncomeAndExpenses failed: data is nil")
	} else if len(*data) == 0 {
		t.Errorf("GetIncomeAndExpenses failed: data is empty")
	}
}
