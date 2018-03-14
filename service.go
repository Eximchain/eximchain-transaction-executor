package main

import (
	"context"
	"errors"
	"log"

	vault "github.com/hashicorp/vault/api"
)

// Manages vault keys and executes transactions against an eximchain node
type TransactionExecutorService interface {
	GetKey(context.Context) (string, error)
}

// concrete implementation of TransactionExecutorService
type transactionExecutorService struct {
	vaultClient *vault.Client
}

func (svc transactionExecutorService) GetKey(_ context.Context) (string, error) {
	pathArg := "keys/singleton"
	vault := svc.vaultClient.Logical()
	secret, err := vault.Read(pathArg)
	if err != nil {
		log.Fatal(err)
		return "", ErrVault
	}
	return secret.Data["key"].(string), nil
}

// ErrVault is returned when there is an error accessing vault.
var ErrVault = errors.New("error accessing vault")
