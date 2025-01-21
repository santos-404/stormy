/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package utils

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	bolt "go.etcd.io/bbolt"
)

func ListAllPasswords() {
	dbPath := getDBPath()

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {

		color.Blue("------------------------------------------")
		color.Blue("|          Available Passwords:          |")
		color.Blue("------------------------------------------")
		color.Blue("|      Service      |      Username      |")
		color.Blue("------------------------------------------")

		empty := true

		tx.ForEach(func(service []byte, bucket *bolt.Bucket) error {
			if string(service) != "MasterPassword" {
				color.Green(string(service))
				cursor := bucket.Cursor()
				for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
					color.Yellow("\t\t\t%s", string(k))
					empty = false
				}
				color.Blue("------------------------------------------")
			}
			return nil
		})

		if empty {
			color.Red("There are no passwords stored yet.")
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to list all passwords: %v", err)
	}

}

func ListPasswordsByService(service string) {
	dbPath := getDBPath()

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

		color.Blue("------------------------------------------")
		color.Blue("|          Available Passwords:          |")
		color.Blue("------------------------------------------")
		color.Blue("|      Service      |      Username      |")
		color.Blue("------------------------------------------")

		cursor := bucket.Cursor()

		color.Green(string(service))
		for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
			color.Yellow("\t\t\t%s", string(k))
		}
		color.Blue("------------------------------------------")

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to list passwords: %v", err)
	}
}

func ListAllServices() {
	dbPath := getDBPath()

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		empty := true
		color.Blue("-------------------------")
		color.Blue("|  Available Services:  |")
		color.Blue("-------------------------")
		index := 1

		tx.ForEach(func(service []byte, b *bolt.Bucket) error {
			if string(service) != "MasterPassword" {
				color.Green("%2d. %s\n", index, string(service))
				index++
				empty = false
			}
			return nil
		})

		if empty {
			color.Red("There are no services stored yet.")
		} else {
			color.Blue("-------------------------")
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to list services: %v", err)
	}
}
