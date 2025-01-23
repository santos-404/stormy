/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func getDBPath() string {
	var err error

	_ = godotenv.Load()
	path := os.Getenv("DB_PATH")

	if path == "" {
		path, err = os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}
	}

	return filepath.Join(path, ".stormy.db")
}

func SetPath(force bool) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which path do you want to use for the database?(./):")
	path, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read the input: %v", err)
	}
	newEnv := "DB_PATH=" + strings.TrimSpace(path)

	if !force {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are you sure you want to change the database path? [y/N]: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read the input: %v", err)
		}

		input = strings.ToLower(strings.TrimSpace(input))
		if input == "y" || input == "yes" {
			force = true
		}
	}

	if force {
		d1 := []byte(newEnv)
		err = os.WriteFile(".env", d1, 0644)
		if err != nil {
			log.Fatalf("Failed to write to .env file: %v", err)
		}
	}
}
