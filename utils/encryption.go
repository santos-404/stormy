/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package utils

import (
	"crypto/aes"
	"crypto/cipher"
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
		saltBytes = []byte(generateRandomStrings())
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

		bucket.Put([]byte("Salt"), saltBytes)
		return bucket.Put([]byte("MasterPassword"), hashedPassword)
	})
	if err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}

	fmt.Println("Master password added successfully!")
}

func encryptPassword(password string, db *bolt.DB) ([]byte, error) {
	masterPassword := getMasterPassword(db)
	block, err := aes.NewCipher(masterPassword)
	if err != nil {
		return nil, err
	}

	// Use AES-GCM for authenticated encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(password), nil)
	return ciphertext, nil
}

func getMasterPassword(db *bolt.DB) []byte {
	var masterPassword []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MasterPassword"))
		if bucket == nil {
			return fmt.Errorf("master password not set. \nTry: stormy set-master-password --help")
		}

		masterPassword = bucket.Get([]byte("MasterPassword"))

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to get master password: %v", err)
	}

	return masterPassword
}

func decryptPassword(encryptedPassword, masterPassword []byte, db *bolt.DB) (string, error) {

	var saltBytes []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MasterPassword"))
		if bucket == nil {
			return fmt.Errorf("master password not set")
		}
		saltBytes = bucket.Get([]byte("Salt"))

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to get salt: %v", err)
	}

	hashedPassword := pbkdf2.Key(masterPassword, saltBytes, 10000, 32, sha256.New)

	block, err := aes.NewCipher(hashedPassword)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedPassword) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := encryptedPassword[:nonceSize], encryptedPassword[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
