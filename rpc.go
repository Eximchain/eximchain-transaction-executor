package main

import (
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func MakeRPCHandler(svc transactionExecutorService) *jsonrpc.Server {
	m := make(jsonrpc.EndpointCodecMap)

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

	m["web3_clientVersion"] = jsonrpc.EndpointCodec{
		Endpoint: makeWeb3ClientVersionEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["web3_sha3"] = jsonrpc.EndpointCodec{
		Endpoint: makeWeb3Sha3Endpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["net_version"] = jsonrpc.EndpointCodec{
		Endpoint: makeNetVersionEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["net_peerCount"] = jsonrpc.EndpointCodec{
		Endpoint: makeNetPeerCountEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["net_listening"] = jsonrpc.EndpointCodec{
		Endpoint: makeNetListeningEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_protocolVersion"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthProtocolVersionEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_syncing"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSyncingEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_coinbase"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthCoinbaseEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_mining"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthMiningEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_hashrate"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthHashrateEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_gasPrice"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGasPriceEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_blockNumber"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthBlockNumberEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getBalance"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBalanceEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getStorageAt"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetStorageAtEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getTransactionCount"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetTransactionCountEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getBlockTransactionCountByHash"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBlockTransactionCountByHashEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getBlockTransactionCountByNumber"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBlockTransactionCountByNumberEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getUncleCountByBlockHash"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetUncleCountByBlockHashEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getUncleCountByBlockNumber"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetUncleCountByBlockNumberEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getCode"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetCodeEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_sign"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSignEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_sendRawTransaction"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSendRawTransactionEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_call"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthCallEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_estimateGas"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthEstimateGasEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getBlockByHash"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBlockByHashEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getBlockByNumber"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetBlockByNumberEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getTransactionByHash"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetTransactionByHashEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getTransactionByBlockHashAndIndex"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetTransactionByBlockHashAndIndexEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getTransactionByBlockNumberAndIndex"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetTransactionByBlockNumberAndIndexEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getTransactionReceipt"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetTransactionReceiptEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getUncleByBlockHashAndIndex"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetUncleByBlockHashAndIndexEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getUncleByBlockNumberAndIndex"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetUncleByBlockNumberAndIndexEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_newFilter"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthNewFilterEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_newBlockFilter"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthNewBlockFilterEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_newPendingTransactionFilter"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthNewPendingTransactionFilterEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_uninstallFilter"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthUninstallFilterEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getFilterChanges"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetFilterChangesEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getFilterLogs"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetFilterLogsEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getLogs"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetLogsEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_getWork"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthGetWorkEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_submitWork"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSubmitWorkEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	m["eth_submitHashrate"] = jsonrpc.EndpointCodec{
		Endpoint: makeEthSubmitHashrateEndpoint(svc),
		Decode:   decodeRPCRequest,
		Encode:   encodeRPCResponse,
	}

	handler := jsonrpc.NewServer(m)

	return handler
}
