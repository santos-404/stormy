/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package utils

import (
	"crypto/sha256"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/pbkdf2"
)

func SetMasterPasword(masterPassword, salt string) {
	var saltBytes []byte

	if salt != "" {
		saltBytes = []byte(salt)
	} else {
		saltBytes = []byte("defaultSalt")
	}

	hashedPassword := pbkdf2.Key([]byte(masterPassword), saltBytes, 10000, 32, sha256.New)

	dbPath := getDBPath()
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("MasterPassword"))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}

		existingPassword := bucket.Get([]byte("MasterPassword"))
		if existingPassword != nil {
			return fmt.Errorf("master password already set; cannot overwrite it")
		}

		return bucket.Put([]byte("MasterPassword"), hashedPassword)
	})
	if err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}

	fmt.Println("Master password added successfully!")
}
