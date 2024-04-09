package util

import (
	"os"
	"path/filepath"
)

func FindRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		// Check if the current directory contains the go.mod file
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, err
		}
		// Move to the parent directory
		dir = filepath.Dir(dir)
	}
}
