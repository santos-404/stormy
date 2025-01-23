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
	path := strings.TrimSpace(os.Getenv("DB_PATH"))

	if path == "" {
		path, err = os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}
	}
	dbPath := filepath.Join(path, ".stormy.db")

	absPath, err := filepath.Abs(dbPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	return absPath
}

func SetDBPath(force bool) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which path do you want to use for the database?(/home/user):")

	newPath, err := reader.ReadString('\n')
	newPath = strings.TrimSpace(newPath)
	newPath = strings.TrimSuffix(newPath, "/")

	if err != nil {
		log.Fatalf("Failed to read the input: %v", err)
	}
	newEnv := "DB_PATH=" + newPath

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
		moveDB(newPath)
		d1 := []byte(newEnv)
		err = os.WriteFile(".env", d1, 0644)
		if err != nil {
			log.Fatalf("Failed to write to .env file: %v", err)
		}
	}
}

func moveDB(newPath string) {
	oldPath := getDBPath()

	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		fmt.Println("reach here")
		return
	}

	newFileName := filepath.Join(newPath, ".stormy.db")

	err := os.Rename(oldPath, newFileName)
	if err != nil {
		log.Fatalf("Failed to move the database: %v", err)
	}
}
