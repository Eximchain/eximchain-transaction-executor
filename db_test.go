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

func TestCreateToken(t *testing.T) {
	token, err := CreateToken()

	if err != nil {
		t.Fatalf("cannot create token: %s", err)
	}

	if len(token) == 0 {
		t.Fatalf("cannot create token: %s", token)
	}
}

func TestUser(t *testing.T) {
	db := NewTestDB()
	defer db.Close()

	token, err := db.CreateUser("test@example.com")
	if err != nil {
		t.Fatalf("cannot create user %s", err)
	}

	if len(token) == 0 {
		t.Fatalf("cannot create user %s", token)
	}

	email, err := db.GetUser(token)
	if err != nil {
		t.Fatalf("cannot get user %s", err)
	}

	if email != "test@example.com" {
		t.Fatalf("cannot get user %s", email)
	}

	token1, err := db.GetTokenByEmail(email)
	if err != nil {
		t.Fatalf("cannot get token by email %s", err)
	}

	if token1 != token {
		t.Fatalf("cannot get token by email %s %s", token1, token)
	}

	db.DeleteUserByToken(token1)

	email1, err := db.GetUser(token1)
	if err != nil {
		t.Fatalf("cannot get user %s", err)
	}

	if len(email1) > 0 {
		t.Fatalf("cannot delete user %s", email1)
	}
}
