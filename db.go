package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	bolt "github.com/coreos/bbolt"
)

type BoltDB struct {
	*bolt.DB
	userBucket []byte
}

func (db *BoltDB) Open(path string) error {
	var err error

	db.DB, err = bolt.Open(path, 0600, nil)

	if err != nil {
		return err
	}

	db.userBucket = []byte("users")

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(db.userBucket)

		if err != nil {
			return errors.New("create user bucket error")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (db *BoltDB) Close() error {
	err := db.DB.Close()
	return err
}

func (db *BoltDB) CreateUser(email string) (string, error) {
	if len(email) == 0 {
		return "", errors.New("user email is empty")
	}

	var token string

	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.userBucket)
		t, err := CreateToken()
		if err != nil {
			return err
		}

		token = t
		return b.Put([]byte(token), []byte(email))
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func (db *BoltDB) GetUser(token string) (string, error) {
	email := ""
	k := []byte(token)

	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.userBucket)
		v := b.Get(k)
		email = string(v)
		return nil
	})

	if err != nil {
		return "", err
	}

	return email, nil
}

func (db *BoltDB) GetTokenByEmail(email string) (string, error) {
	token := ""

	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.userBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) == email {
				token = string(k)
				return nil
			}
		}

		return nil
	})

	return token, nil
}

func (db *BoltDB) ListUsers() {
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.userBucket)
		c := b.Cursor()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Fprintf(w, "%s\t%s\n", v, k)
		}

		w.Flush()

		return nil
	})
}

func (db *BoltDB) DeleteUserByToken(token string) error {
	k := []byte(token)
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.userBucket)
		return b.Delete(k)
	})

	if err != nil {
		log.Println("DeleteUserByToken error", err)
	}

	return err
}
