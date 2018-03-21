package main

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/eximchain/eth-client/quorum"
	"github.com/eximchain/go-ethereum/accounts/keystore"
	"github.com/eximchain/go-ethereum/core/types"

	vault "github.com/hashicorp/vault/api"
)

// Manages vault keys and executes transactions against an eximchain node
type TransactionExecutorService interface {
	GetKey(context.Context) (string, error)
	GenerateKey(context.Context) (string, error)
}

// concrete implementation of TransactionExecutorService
type transactionExecutorService struct {
	vaultClient  *vault.Client
	quorumClient *quorum.Client
	keystore     *keystore.KeyStore
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

	address := "0x" + hex.EncodeToString(account.Address.Bytes())
	return address, nil
}

func (svc *transactionExecutorService) signTx(tx *types.Transaction) (*types.Transaction, error) {
	// TODO: Something less presumptuous to get the signing account
	wallets := svc.keystore.Wallets()
	if len(wallets) == 0 {
		log.Println("No wallets found")
		return nil, ErrKeystore
	}
	wallet := wallets[0]

	accounts := wallet.Accounts()
	if len(accounts) == 0 {
		log.Println("No accounts found")
		return nil, ErrKeystore
	}
	account := accounts[0]

	signedTx, err := svc.keystore.SignTx(account, tx, nil)
	if err != nil {
		log.Println(err)
		return nil, ErrKeystore
	}

	return signedTx, nil
}

// ErrVault is returned when there is an error accessing vault.
var ErrVault = errors.New("error accessing vault")

// ErrKeystore is returned when there is an error using the keystore
var ErrKeystore = errors.New("error using keystore")
