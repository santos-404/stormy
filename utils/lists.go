/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package utils

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func ListAllPasswords() {
	dbPath := GetDBPath()

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(service []byte, bucket *bolt.Bucket) error {
			fmt.Println("Service:", string(service))

			cursor := bucket.Cursor()
			for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
				fmt.Printf("\tUsername: %s\n", k)
			}
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to list all passwords: %v", err)
	}

}

func ListPasswordsByService(service string) {
	dbPath := GetDBPath()

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(service))
		if bucket == nil {
			return fmt.Errorf("service %s not found", service)
		}

		cursor := bucket.Cursor()

		fmt.Println("Service:", string(service))
		for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
			fmt.Printf("\tUsername: %s\n", k)
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to list passwords: %v", err)
	}
}
