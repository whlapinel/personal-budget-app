package viewmodels

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestNetWorthReport(t *testing.T) {
	err := godotenv.Load("../../.env.backend")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	err = godotenv.Load("../../.env.backend.testing")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// test GetNetWorthReport
	email := "test@test.com"
	data, err := GetNetWorthReport(email)
	log.Print()
	if err != nil {
		t.Errorf("GetNetWorthReport failed: %v", err)
	}
	if data == nil {
		t.Errorf("GetNetWorthReport failed: data is nil")
	}
}
