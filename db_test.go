package main

import (
	"testing"
)

func NewTestDB() *BoltDB {
	db := &BoltDB{}
	err := db.open("eximchain_test.db")

	if err != nil {
		panic("cannot open db")
	}

	return db
}

func TestCreateToken(t *testing.T) {
	token, err := createToken()

	if err != nil {
		t.Fatalf("cannot create token: %s", err)
	}

	if len(token) == 0 {
		t.Fatalf("cannot create token: %s", token)
	}
}

func TestUser(t *testing.T) {
	db := NewTestDB()
	defer db.close()

	token, err := db.createUser("test@example.com")
	if err != nil {
		t.Fatalf("cannot create user %s", err)
	}

	if len(token) == 0 {
		t.Fatalf("cannot create user %s", token)
	}

	email, err := db.getUser(token)
	if err != nil {
		t.Fatalf("cannot get user %s", err)
	}

	if email != "test@example.com" {
		t.Fatalf("cannot get user %s", email)
	}

	token1, err := db.getTokenByEmail(email)
	if err != nil {
		t.Fatalf("cannot get token by email %s", err)
	}

	if token1 != token {
		t.Fatalf("cannot get token by email %s %s", token1, token)
	}

	db.deleteUserByToken(token1)

	email1, err := db.getUser(token1)
	if err != nil {
		t.Fatalf("cannot get user %s", err)
	}

	if len(email1) > 0 {
		t.Fatalf("cannot delete user %s", email1)
	}
}
