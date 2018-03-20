package main

import (
	"context"
	"errors"
	"log"

	"github.com/eximchain/go-ethereum/accounts/keystore"

	vault "github.com/hashicorp/vault/api"
)

// Manages vault keys and executes transactions against an eximchain node
type TransactionExecutorService interface {
	GetKey(context.Context) (string, error)
	GenerateKey(context.Context) (string, error)
}

// concrete implementation of TransactionExecutorService
type transactionExecutorService struct {
	vaultClient *vault.Client
	keystore    *keystore.KeyStore
}

func (svc transactionExecutorService) GetKey(_ context.Context) (string, error) {
	pathArg := "keys/singleton"
	vault := svc.vaultClient.Logical()
	secret, err := vault.Read(pathArg)
	if err != nil {
		log.Println(err)
		return "", ErrVault
	}
	key, present := secret.Data["key"]
	if !present {
		log.Fatal("Vault entry found but key not present")
	}
	return key.(string), nil
}

func (svc transactionExecutorService) GenerateKey(_ context.Context) (string, error) {
	// TODO: Use a real password
	password := ""
	account, err := svc.keystore.NewAccount(password)
	if err != nil {
		log.Println(err)
		return "", ErrKeystore
	}
	return account.Address.Str(), nil
}

// ErrVault is returned when there is an error accessing vault.
var ErrVault = errors.New("error accessing vault")

// ErrKeystore is returned when there is an error using the keystore
var ErrKeystore = errors.New("error using keystore")
