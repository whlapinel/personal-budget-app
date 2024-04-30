package util

import "os"

func GetHost() string {
	production := os.Getenv("PRODUCTION")
	var address string
	if production == "false" {
		address = "localhost"
	} else {
		address = "mariadb"
	}
	return address
}
