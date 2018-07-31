package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

type Keyfile struct {
	Address string `json:"address"`
}

func makeEthAccountsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetVaultKey(ctx)
		if err != nil {
			return nil, err
		}

		var data Keyfile
		json.Unmarshal([]byte(v), &data)
		return []string{"0x" + data.Address}, nil
	}
}

func makePersonalNewAccountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GenerateKey(ctx)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeEthSendTransactionEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCTransactionParams)

		from := req[0].From
		to := req[0].To
		amount, _ := strconv.ParseInt(req[0].Value, 0, 64)
		gasLimit, _ := strconv.ParseUint(req[0].Gas, 0, 64)
		gasPrice, _ := strconv.ParseInt(req[0].GasPrice, 0, 64)
		data := req[0].Data

		txHash, err := svc.ExecuteTransaction(ctx, from, to, amount, gasLimit, gasPrice, data)

		if err != nil {
			return nil, err
		}
		return txHash, nil
	}
}

func makeWeb3ClientVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.Web3ClientVersion(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeWeb3Sha3Endpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.Web3Sha3(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetVersion(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetPeerCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetPeerCount(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetListeningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetListening(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthProtocolVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthProtocolVersion(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSyncingEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSyncing(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCoinbaseEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthCoinbase(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthMiningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthMining(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthHashrate(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGasPriceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGasPrice(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthBlockNumber(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBalanceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBalance(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetStorageAtEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetStorageAt(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionCount(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockTransactionCountByHash(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockTransactionCountByNumber(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleCountByBlockHash(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleCountByBlockNumber(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetCodeEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetCode(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSignEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSign(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCallEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthCall(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthEstimateGasEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthEstimateGas(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockByHash(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockByNumber(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByHash(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByBlockHashAndIndex(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByBlockNumberAndIndex(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionReceiptEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionReceipt(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleByBlockHashAndIndex(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleByBlockNumberAndIndex(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewFilter(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewBlockFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewBlockFilter(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewPendingTransactionFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewPendingTransactionFilter(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthUninstallFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthUninstallFilter(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterChangesEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetFilterChanges(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetFilterLogs(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetLogs(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetWork(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSubmitWork(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSubmitHashrate(ctx, request)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func decodeRPCRequest(ctx context.Context, msg json.RawMessage) (interface{}, error) {
	var req interface{}
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
