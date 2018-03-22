package main

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"math/big"

	"github.com/eximchain/eth-client/quorum"
	"github.com/eximchain/go-ethereum/accounts"
	"github.com/eximchain/go-ethereum/accounts/keystore"
	"github.com/eximchain/go-ethereum/core/types"

	ethCommon "github.com/eximchain/go-ethereum/common"
	vault "github.com/hashicorp/vault/api"
)

// Manages vault keys and executes transactions against an eximchain node
type TransactionExecutorService interface {
	ExecuteTransaction(context.Context, string, string) error
	GetVaultKey(context.Context) (string, error)
	GenerateKey(context.Context) (string, error)
}

// concrete implementation of TransactionExecutorService
type transactionExecutorService struct {
	vaultClient  *vault.Client
	quorumClient quorum.Client
	keystore     *keystore.KeyStore
	accountCache map[string]accounts.Account
}

// Currently proof of concept only
func (svc transactionExecutorService) GetVaultKey(_ context.Context) (string, error) {
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
	svc.accountCache[address] = account
	return address, nil
}

func (svc transactionExecutorService) ExecuteTransaction(ctx context.Context, from string, to string) error {
	// TODO: Replace with vault backend
	account, present := svc.accountCache[from]
	if !present {
		return ErrAccountMissing
	}
	password := ""

	nonce, err := svc.quorumClient.PendingNonceAt(ctx, account.Address)
	if err != nil {
		log.Println("Error: PendingNonceAt")
		log.Println(err)
		return ErrQuorum
	}
	// TODO: Do something with these parameters
	amount := big.NewInt(1000000000000000000)
	gasLimit := uint64(1000000000)
	gasPrice := big.NewInt(0)
	data := make([]byte, 0, 0)

	tx := types.NewTransaction(nonce, ethCommon.HexToAddress(to), amount, gasLimit, gasPrice, data)
	// Chain ID must be nil for quorum
	tx, err = svc.keystore.SignTxWithPassphrase(account, password, tx, nil)
	if err != nil {
		log.Println("Error: Signing")
		log.Println(err)
		return ErrSigning
	}
	err = svc.quorumClient.SendRawTransaction(ctx, tx)
	if err != nil {
		log.Println("Error: SendTransaction")
		log.Println(err)
		return ErrQuorum
	}
	return nil
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

// ErrQuorum is returned when there is an error using the quorum client
var ErrQuorum = errors.New("error using quorum client")

// ErrAccountMissing is returned when the requested account could not be found
var ErrAccountMissing = errors.New("account not found")

// ErrSigning is returned when there is an error signing the transaction
var ErrSigning = errors.New("error signing transaction")
