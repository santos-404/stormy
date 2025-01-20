/*
Copyright © 2025 NAME HERE javier.jsm21@gmail.com
*/
package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	bolt "go.etcd.io/bbolt"
)

func AddPassword(service, username, password string) {
	dbPath := getDBPath()
	fmt.Println(dbPath)

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Passwords"))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}

		data := fmt.Sprintf("username: %s, password: %s", username, password)

		return bucket.Put([]byte(service), []byte(data))
	})
	if err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}

	fmt.Println("Password added successfully!")
}

func getDBPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}
	return filepath.Join(homeDir, ".passwordmanager.db")
}
