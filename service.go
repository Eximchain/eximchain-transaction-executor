package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"time"

	"github.com/eximchain/eth-client/quorum"
	"github.com/eximchain/go-ethereum/accounts"
	"github.com/eximchain/go-ethereum/accounts/keystore"
	"github.com/eximchain/go-ethereum/core/types"
	"github.com/eximchain/go-ethereum/crypto"
	"github.com/go-kit/kit/transport/http/jsonrpc"
      ethRlp "github.com/eximchain/go-ethereum/rlp"
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

	Web3ClientVersion(context.Context, interface{}) (interface{}, error)
	Web3Sha3(context.Context, interface{}) (interface{}, error)
	NetVersion(context.Context, interface{}) (interface{}, error)
	NetPeerCount(context.Context, interface{}) (interface{}, error)
	NetListening(context.Context, interface{}) (interface{}, error)
	EthProtocolVersion(context.Context, interface{}) (interface{}, error)
	EthSyncing(context.Context, interface{}) (interface{}, error)
	EthCoinbase(context.Context, interface{}) (interface{}, error)
	EthMining(context.Context, interface{}) (interface{}, error)
	EthHashrate(context.Context, interface{}) (interface{}, error)
	EthGasPrice(context.Context, interface{}) (interface{}, error)
	EthAccounts(context.Context, interface{}) (interface{}, error)
	EthBlockNumber(context.Context, interface{}) (interface{}, error)
	EthGetBalance(context.Context, interface{}) (interface{}, error)
	EthGetStorageAt(context.Context, interface{}) (interface{}, error)
	EthGetTransactionCount(context.Context, interface{}) (interface{}, error)
	EthGetBlockTransactionCountByHash(context.Context, interface{}) (interface{}, error)
	EthGetBlockTransactionCountByNumber(context.Context, interface{}) (interface{}, error)
	EthGetUncleCountByBlockHash(context.Context, interface{}) (interface{}, error)
	EthGetUncleCountByBlockNumber(context.Context, interface{}) (interface{}, error)
	EthGetCode(context.Context, interface{}) (interface{}, error)
	EthSign(context.Context, string, string) (interface{}, error)
	EthSignTransaction(context.Context, string, string, int64, uint64, int64, string)(interface{}, error)
	EthSendRawTransaction(context.Context, interface{}) (interface{}, error)
	EthCall(context.Context, interface{}) (interface{}, error)
	EthEstimateGas(context.Context, interface{}) (interface{}, error)
	EthGetBlockByHash(context.Context, interface{}) (interface{}, error)
	EthGetBlockByNumber(context.Context, interface{}) (interface{}, error)
	EthGetTransactionByHash(context.Context, interface{}) (interface{}, error)
	EthGetTransactionByBlockHashAndIndex(context.Context, interface{}) (interface{}, error)
	EthGetTransactionByBlockNumberAndIndex(context.Context, interface{}) (interface{}, error)
	EthGetTransactionReceipt(context.Context, interface{}) (interface{}, error)
	EthGetUncleByBlockHashAndIndex(context.Context, interface{}) (interface{}, error)
	EthGetUncleByBlockNumberAndIndex(context.Context, interface{}) (interface{}, error)
	EthNewFilter(context.Context, interface{}) (interface{}, error)
	EthNewBlockFilter(context.Context, interface{}) (interface{}, error)
	EthNewPendingTransactionFilter(context.Context, interface{}) (interface{}, error)
	EthUninstallFilter(context.Context, interface{}) (interface{}, error)
	EthGetFilterChanges(context.Context, interface{}) (interface{}, error)
	EthGetFilterLogs(context.Context, interface{}) (interface{}, error)
	EthGetLogs(context.Context, interface{}) (interface{}, error)
	EthGetWork(context.Context, interface{}) (interface{}, error)
	EthSubmitWork(context.Context, interface{}) (interface{}, error)
	EthSubmitHashrate(context.Context, interface{}) (interface{}, error)
}

// concrete implementation of TransactionExecutorService
type transactionExecutorService struct {
	vaultClient   *vault.Client
	quorumClient  quorum.Client
	quorumAddress string
	keystore      *keystore.KeyStore
	accountCache  map[string]accounts.Account
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
	// account, present := svc.accountCache[from]
	// if !present {
	//   return "", ErrAccountMissing
	// }

	accs := svc.keystore.Accounts()
	var account accounts.Account

	for _, a := range accs {
		if ethCommon.HexToAddress(from) == a.Address {
			account = a
			break
		}
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

func (svc transactionExecutorService) Web3ClientVersion(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "web3_clientVersion")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) Web3Sha3(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "web3_sha3")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) NetVersion(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "net_version")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) NetPeerCount(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "net_peerCount")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) NetListening(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "net_listening")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthProtocolVersion(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_protocolVersion")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthSyncing(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_syncing")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthCoinbase(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_coinbase")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthMining(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_mining")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthHashrate(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_hashrate")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGasPrice(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_gasPrice")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthAccounts(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_accounts")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthBlockNumber(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_blockNumber")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetBalance(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getBalance")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetStorageAt(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getStorageAt")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetTransactionCount(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getTransactionCount")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetBlockTransactionCountByHash(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getBlockTransactionCountByHash")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetBlockTransactionCountByNumber(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getBlockTransactionCountByNumber")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetUncleCountByBlockHash(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getUncleCountByBlockHash")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetUncleCountByBlockNumber(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getUncleCountByBlockNumber")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetCode(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getCode")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func (svc transactionExecutorService) EthSign(ctx context.Context, address, data string) (interface{}, error) {
	accs := svc.keystore.Accounts()
	var account accounts.Account

	for _, a := range accs {
		if ethCommon.HexToAddress(address) == a.Address {
			account = a
			break
		}
	}

	password := ""
	signature, err := svc.keystore.SignHashWithPassphrase(account, password, signHash(ethCommon.FromHex(data)))
	if err != nil {
		return nil, err
	}

	signature[64] += 27

	return ethCommon.ToHex(signature), nil
}

func (svc transactionExecutorService) EthSignTransaction(ctx context.Context, from string, to string, amount int64, gasLimit uint64, gasPrice int64, hexData string)(interface{}, error) {
	accs := svc.keystore.Accounts()
	var account accounts.Account

	for _, a := range accs {
		if ethCommon.HexToAddress(from) == a.Address {
			account = a
			break
		}
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

      rlpData, err := ethRlp.EncodeToBytes(tx)

	if err != nil {
		log.Println("Error: RLP encoding")
		log.Println(err)
		return "", err
	}

      str := ethCommon.ToHex(rlpData)

      return str, nil
}

func (svc transactionExecutorService) EthSendRawTransaction(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_sendRawTransaction")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthCall(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_call")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthEstimateGas(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_estimateGas")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetBlockByHash(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getBlockByHash")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetBlockByNumber(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getBlockByNumber")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetTransactionByHash(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getTransactionByHash")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetTransactionByBlockHashAndIndex(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getTransactionByBlockHashAndIndex")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetTransactionByBlockNumberAndIndex(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getTransactionByBlockNumberAndIndex")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetTransactionReceipt(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getTransactionReceipt")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetUncleByBlockHashAndIndex(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getUncleByBlockHashAndIndex")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetUncleByBlockNumberAndIndex(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getUncleByBlockNumberAndIndex")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthNewFilter(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_newFilter")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthNewBlockFilter(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_newBlockFilter")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthNewPendingTransactionFilter(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_newPendingTransactionFilter")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthUninstallFilter(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_uninstallFilter")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetFilterChanges(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getFilterChanges")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetFilterLogs(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getFilterLogs")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetLogs(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getLogs")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthGetWork(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_getWork")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthSubmitWork(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_submitWork")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc transactionExecutorService) EthSubmitHashrate(ctx context.Context, params interface{}) (interface{}, error) {
	u, _ := url.Parse(svc.quorumAddress)
	client := jsonrpc.NewClient(u, "eth_submitHashrate")
	res, err := client.Endpoint()(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
