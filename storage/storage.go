/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
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

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(service))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}

		// TODO: Encrypt the password before storing it
		return bucket.Put([]byte(username), []byte(password))
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
