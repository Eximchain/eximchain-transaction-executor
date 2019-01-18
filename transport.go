package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

func makePersonalNewAccountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "personal_newAccount"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GenerateKey(ctx)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeEthSendTransactionEndpoint(svc transactionExecutorService) endpoint.Endpoint {
	methodName := "eth_sendTransaction"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCTransactionParams)

		from := req[0].From
		to := req[0].To
		amount, _ := strconv.ParseInt(req[0].Value, 0, 64)
		gasLimit, _ := strconv.ParseUint(req[0].Gas, 0, 64)
		gasPrice, _ := strconv.ParseInt(req[0].GasPrice, 0, 64)
		data := req[0].Data

		txHash, err := svc.ExecuteTransaction(ctx, from, to, amount, gasLimit, gasPrice, data)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return txHash, nil
	}
}

func makeWeb3ClientVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "web3_clientVersion"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.Web3ClientVersion(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeWeb3Sha3Endpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "web3_sha3"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.Web3Sha3(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "net_version"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetVersion(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetPeerCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "net_peerCount"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetPeerCount(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeNetListeningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "net_listening"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.NetListening(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthProtocolVersionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_protocolVersion"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthProtocolVersion(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSyncingEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_syncing"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSyncing(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCoinbaseEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_coinbase"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthCoinbase(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthMiningEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_mining"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthMining(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_hashrate"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthHashrate(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGasPriceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_gasPrice"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGasPrice(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthAccountsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_accounts"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthAccounts(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_blockNumber"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthBlockNumber(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBalanceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getBalance"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBalance(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetStorageAtEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getStorageAt"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetStorageAt(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionCountEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getTransactionCount"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionCount(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getBlockTransactionCountByHash"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockTransactionCountByHash(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockTransactionCountByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getBlockTransactionCountByNumber"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockTransactionCountByNumber(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getUncleCountByBlockHash"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleCountByBlockHash(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleCountByBlockNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getUncleCountByBlockNumber"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleCountByBlockNumber(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetCodeEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getCode"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetCode(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSignEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_sign"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.([]interface{})
		address := req[0].(string)
		data := req[1].(string)

		res, err := svc.EthSign(ctx, address, data)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSignTransactionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_signTransaction"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RPCTransactionParams)

		from := req[0].From
		to := req[0].To
		amount, _ := strconv.ParseInt(req[0].Value, 0, 64)
		gasLimit, _ := strconv.ParseUint(req[0].Gas, 0, 64)
		gasPrice, _ := strconv.ParseInt(req[0].GasPrice, 0, 64)
		data := req[0].Data

		txHash, err := svc.EthSignTransaction(ctx, from, to, amount, gasLimit, gasPrice, data)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		return txHash, err
	}
}

func makeEthSendRawTransactionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_sendRawTransaction"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSendRawTransaction(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthCallEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_call"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthCall(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthEstimateGasEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_estimateGas"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthEstimateGas(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getBlockByHash"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockByHash(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetBlockByNumberEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getBlockByNumber"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetBlockByNumber(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByHashEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getTransactionByHash"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByHash(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getTransactionByBlockHashAndIndex"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByBlockHashAndIndex(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getTransactionByBlockNumberAndIndex"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionByBlockNumberAndIndex(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetTransactionReceiptEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getTransactionReceipt"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetTransactionReceipt(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockHashAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getUncleByBlockHashAndIndex"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleByBlockHashAndIndex(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetUncleByBlockNumberAndIndexEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getUncleByBlockNumberAndIndex"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetUncleByBlockNumberAndIndex(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_newFilter"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewFilter(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewBlockFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_newBlockFilter"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewBlockFilter(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthNewPendingTransactionFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_newPendingTransactionFilter"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthNewPendingTransactionFilter(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthUninstallFilterEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_uninstallFilter"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthUninstallFilter(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterChangesEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getFilterChanges"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetFilterChanges(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetFilterLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getFilterLogs"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetFilterLogs(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetLogsEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getLogs"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetLogs(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthGetWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_getWork"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthGetWork(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitWorkEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_submitWork"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSubmitWork(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func makeEthSubmitHashrateEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	methodName := "eth_submitHashrate"
	logger := log.WithField("method", methodName)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.EthSubmitHashrate(ctx, request)

		success := false
		if err == nil {
			success = true
		}
		logger = logger.WithFields(log.Fields{"success": success, "err": err})
		logger.Info("RPC call served")

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
