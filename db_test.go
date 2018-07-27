package main

import (
	"testing"
)

func NewTestDB() *BoltDB {
	db := &BoltDB{}
	err := db.Open("eximchain_test.db")

	if err != nil {
		panic("cannot open db")
	}

	return db
}

func TestUser(t *testing.T) {
	db := NewTestDB()
	defer db.Close()

	token, _ := db.CreateUser("test@example.com")

	if len(token) == 0 {
		t.Error("cannot create user")
	}

	email, _ := db.GetUser(token)

	if email != "test@example.com" {
		t.Error("cannot get user")
	}

	token1, _ := db.GetTokenByEmail(email)

	if token1 != token {
		t.Error("cannot get token by email")
	}

	db.DeleteUserByToken(token1)

	email1, _ := db.GetUser(token1)

	if len(email1) > 0 {
		t.Error("cannot delete user")
	}
}
