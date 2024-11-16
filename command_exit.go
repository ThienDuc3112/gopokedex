package main

import (
	"os"
)

func commandExit(*Config, []string) error {
	os.Exit(0)
	return nil
}
