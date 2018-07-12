package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	ethCommon "github.com/eximchain/go-ethereum/common"
	"github.com/eximchain/go-ethereum/core/types"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

var RPCMethods = []string{
	"web3_clientVersion",
	"web3_sha3",
	"net_version",
	"net_peerCount",
	"net_listening",
	"eth_protocolVersion",
	"eth_syncing",
	"eth_coinbase",
	"eth_mining",
	"eth_hashrate",
	"eth_gasPrice",
	// "eth_accounts",
	"eth_blockNumber",
	// "eth_getBalance",
	"eth_getStorageAt",
	"eth_getTransactionCount",
	"eth_getBlockTransactionCountByHash",
	"eth_getBlockTransactionCountByNumber",
	"eth_getUncleCountByBlockHash",
	"eth_getUncleCountByBlockNumber",
	"eth_getCode",
	"eth_sign",
	// "eth_sendTransaction",
	"eth_sendRawTransaction",
	"eth_call",
	"eth_estimateGas",
	"eth_getBlockByHash",
	"eth_getBlockByNumber",
	"eth_getTransactionByHash",
	"eth_getTransactionByBlockHashAndIndex",
	"eth_getTransactionByBlockNumberAndIndex",
	"eth_getTransactionReceipt",
	"eth_getUncleByBlockHashAndIndex",
	"eth_getUncleByBlockNumberAndIndex",
	"eth_getCompilers",
	"eth_compileLLL",
	"eth_compileSolidity",
	"eth_compileSerpent",
	"eth_newFilter",
	"eth_newBlockFilter",
	"eth_newPendingTransactionFilter",
	"eth_uninstallFilter",
	"eth_getFilterChanges",
	"eth_getFilterLogs",
	"eth_getLogs",
	"eth_getWork",
	"eth_submitWork",
	"eth_submitHashrate",
	"db_putString",
	"db_getString",
	"db_putHex",
	"db_getHex",
	"shh_post",
	"shh_version",
	"shh_newIdentity",
	"shh_hasIdentity",
	"shh_newGroup",
	"shh_addToGroup",
	"shh_newFilter",
	"shh_uninstallFilter",
	"shh_getFilterChanges",
	"shh_getMessages",
}

func MakeRPCHandler(svc transactionExecutorService, quorumAddress string) *jsonrpc.Server {
	u, _ := url.Parse(quorumAddress)

	m := make(jsonrpc.EndpointCodecMap)

	for _, method := range RPCMethods {
		m[method] = jsonrpc.EndpointCodec{
			Endpoint: makeProxyEndpoint(u, method),
			Decode:   decodeRPCRequest,
			Encode:   encodeRPCResponse,
		}
	}

	m["eth_getBalance"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBalanceEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_sendTransaction"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSendTransactionEndpoint(svc),
		Decode:   decodeRPCTransactionRequest,
		Encode:   encodeRPCResponse,
	}

	m["personal_newAccount"] = jsonrpc.EndpointCodec{
		Endpoint: makePersonalNewAccountEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_accounts"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthAccountsEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	handler := jsonrpc.NewServer(m)

	return handler
}

func makeProxyEndpoint(u *url.URL, method string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCParams)
		log.Println("proxy", method, req)

		client := jsonrpc.NewClient(u, method)
		res, err := client.Endpoint()(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBalanceEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCParams)
		address := req[0]
		balance, err := svc.GetBalance(ctx, address)

		if err != nil {
			return nil, err
		}

		return balance, nil
	}
}

func makeEthSendTransactionEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCTransactionParams)
		account, present := svc.accountCache[req[0].From]
		if !present {
			return "", ErrAccountMissing
		}

		password := ""
		nonce, err := svc.quorumClient.PendingNonceAt(ctx, account.Address)
		if err != nil {
			return nil, err
		}

		to := ethCommon.HexToAddress(req[0].To)
		amount, _ := HexToBigInt(req[0].Value)
		gasLimit, _ := strconv.ParseUint(req[0].Gas, 0, 64)
		gasPrice, _ := HexToBigInt(req[0].GasPrice)
		data := ethCommon.FromHex(req[0].Data)

		tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
		tx, err = svc.keystore.SignTxWithPassphrase(account, password, tx, nil)
		if err != nil {
			return nil, err
		}

		err = svc.quorumClient.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}

		txHash := tx.Hash().String()
		return txHash, nil
	}
}

func makePersonalNewAccountEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		password := ""
		account, err := svc.keystore.NewAccount(password)
		if err != nil {
			return nil, err
		}

		address := "0x" + hex.EncodeToString(account.Address.Bytes())
		svc.accountCache[address] = account
		return address, nil
	}
}

func makeEthAccountsEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetVaultKey(ctx)
		if err != nil {
			return nil, err
		}

		return [1]string{v}, nil
	}
}

func decodeRPCRequest(ctx context.Context, msg json.RawMessage) (interface{}, error) {
	var req RPCParams
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeRPCResponse(ctx context.Context, result interface{}) (json.RawMessage, error) {
	b, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func decodeRPCTransactionRequest(ctx context.Context, msg json.RawMessage) (interface{}, error) {
	var req RPCTransactionParams
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type RPCParams = []string

type RPCTransactionParams = []RPCTransaction

type RPCTransaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	Nonce    string `json:"nonce"`
}
