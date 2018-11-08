package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

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
		req := request.(jsonRpcRequest)
		res, err := svc.Web3ClientVersion(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeWeb3Sha3Endpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.Web3Sha3(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.NetVersion(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetPeerCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.NetPeerCount(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetListeningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.NetListening(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthProtocolVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthProtocolVersion(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSyncingEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthSyncing(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCoinbaseEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthCoinbase(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthMiningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthMining(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthHashrate(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGasPriceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGasPrice(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthAccountsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthAccounts(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println(request)
		req := request.(jsonRpcRequest)
		res, err := svc.EthBlockNumber(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBalanceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetBalance(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetStorageAtEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetStorageAt(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetTransactionCount(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetBlockTransactionCountByHash(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetBlockTransactionCountByNumber(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetUncleCountByBlockHash(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetUncleCountByBlockNumber(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetCodeEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetCode(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSignEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.([]interface{})
		address := req[0].(string)
		data := req[1].(string)

		res, err := svc.EthSign(ctx, address, data)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSignTransactionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCTransactionParams)

		from := req[0].From
		to := req[0].To
		amount, _ := strconv.ParseInt(req[0].Value, 0, 64)
		gasLimit, _ := strconv.ParseUint(req[0].Gas, 0, 64)
		gasPrice, _ := strconv.ParseInt(req[0].GasPrice, 0, 64)
		data := req[0].Data

		txHash, err := svc.EthSignTransaction(ctx, from, to, amount, gasLimit, gasPrice, data)

		return txHash, err
	}
}

func makeEthSendRawTransactionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthSendRawTransaction(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCallEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthCall(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthEstimateGasEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthEstimateGas(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetBlockByHash(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetBlockByNumber(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetTransactionByHash(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetTransactionByBlockHashAndIndex(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetTransactionByBlockNumberAndIndex(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionReceiptEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetTransactionReceipt(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetUncleByBlockHashAndIndex(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetUncleByBlockNumberAndIndex(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthNewFilter(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewBlockFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthNewBlockFilter(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewPendingTransactionFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthNewPendingTransactionFilter(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthUninstallFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthUninstallFilter(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterChangesEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetFilterChanges(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetFilterLogs(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetLogs(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthGetWork(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthSubmitWork(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(jsonRpcRequest)
		res, err := svc.EthSubmitHashrate(ctx, req)

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
