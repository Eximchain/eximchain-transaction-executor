package main

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/eximchain/eth-client/quorum"
	"github.com/eximchain/go-ethereum/accounts"
	"github.com/eximchain/go-ethereum/accounts/keystore"
	"github.com/eximchain/go-ethereum/core/types"

	ethCommon "github.com/eximchain/go-ethereum/common"
	vault "github.com/hashicorp/vault/api"
)

// Manages vault keys and executes transactions against an eximchain node
type TransactionExecutorService interface {
	ExecuteTransaction(context.Context, string, string, int64, uint64, int64, string) (string, error)
	GetVaultKey(context.Context) (string, error)
	GenerateKey(context.Context) (string, error)
	GetBalance(context.Context, string) (int64, error)
	RunWorkload(context.Context, string, string, int64, uint64, int64, string, int, int)
	NodeSyncProgress(context.Context) (bool, uint64, uint64, error)
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

func (svc transactionExecutorService) ExecuteTransaction(ctx context.Context, from string, to string, amount int64, gasLimit uint64, gasPrice int64, hexData string) (string, error) {
	// TODO: Replace with vault backend
	account, present := svc.accountCache[from]
	if !present {
		return "", ErrAccountMissing
	}
	password := ""

	nonce, err := svc.quorumClient.PendingNonceAt(ctx, account.Address)
	if err != nil {
		log.Println("Error: PendingNonceAt")
		log.Println(err)
		return "", ErrQuorum
	}

	data := ethCommon.FromHex(hexData)

	tx := types.NewTransaction(nonce, ethCommon.HexToAddress(to), big.NewInt(amount), gasLimit, big.NewInt(gasPrice), data)
	// Chain ID must be nil for quorum
	tx, err = svc.keystore.SignTxWithPassphrase(account, password, tx, nil)
	if err != nil {
		log.Println("Error: Signing")
		log.Println(err)
		return "", ErrSigning
	}
	err = svc.quorumClient.SendTransaction(ctx, tx)
	if err != nil {
		log.Println("Error: SendTransaction")
		log.Println(err)
		return "", ErrQuorum
	}
	txHash := tx.Hash().String()
	return txHash, nil
}

func (svc transactionExecutorService) GetBalance(ctx context.Context, address string) (int64, error) {
	account, present := svc.accountCache[address]
	if !present {
		return int64(0), ErrAccountMissing
	}
	var blockNumber *big.Int
	blockNumber = nil
	balance, err := svc.quorumClient.BalanceAt(ctx, account.Address, blockNumber)
	if err != nil {
		log.Println(err)
		return int64(0), ErrQuorum
	}
	return balance.Int64(), nil
}

func (svc transactionExecutorService) RunWorkload(_ context.Context, from string, to string, amount int64, gasLimit uint64, gasPrice int64, hexData string, sleepSeconds int, numTransactions int) {
	ctx := context.Background()
	go svc.workload(ctx, from, to, amount, gasLimit, gasPrice, hexData, sleepSeconds, numTransactions)
}

func (svc transactionExecutorService) NodeSyncProgress(ctx context.Context) (bool, uint64, uint64, error) {
	syncProgress, err := svc.quorumClient.SyncProgress(ctx)
	if err != nil {
		log.Println(err)
		return false, uint64(0), uint64(0), ErrQuorum
	}

	// Syncing is complete
	if syncProgress == nil {
		return false, uint64(0), uint64(0), nil
	}

	// Syncing still in progress
	return true, syncProgress.CurrentBlock, syncProgress.HighestBlock, nil
}

func (svc transactionExecutorService) workload(ctx context.Context, from string, to string, amount int64, gasLimit uint64, gasPrice int64, hexData string, sleepSeconds int, numTransactions int) {
	sleepDuration := time.Duration(sleepSeconds) * time.Second
	for i := 0; i < numTransactions; i++ {
		_, err := svc.ExecuteTransaction(ctx, from, to, amount, gasLimit, gasPrice, hexData)
		if err != nil {
			log.Printf("Workload Error: %v", err)
		}
		time.Sleep(sleepDuration)
	}
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
