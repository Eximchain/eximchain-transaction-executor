package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGetVaultKeyEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetVaultKey(ctx)
		if err != nil {
			return getVaultKeyResponse{v, err.Error()}, nil
		}
		return getVaultKeyResponse{v, ""}, nil
	}
}

func makeGenerateKeyEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GenerateKey(ctx)
		if err != nil {
			return generateKeyResponse{v, err.Error()}, nil
		}
		return generateKeyResponse{v, ""}, nil
	}
}

func makeExecuteTransactionEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(executeTransactionRequest)
		from, to, amount, gasLimit, gasPrice := req.From, req.To, req.Amount, req.GasLimit, req.GasPrice
		err := svc.ExecuteTransaction(ctx, from, to, amount, gasLimit, gasPrice)
		if err != nil {
			return executeTransactionResponse{err.Error()}, nil
		}
		return executeTransactionResponse{""}, nil
	}
}

func makeRunWorkloadEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(runWorkloadRequest)
		from, to, amount, gasLimit, gasPrice, sleep, num := req.From, req.To, req.Amount, req.GasLimit, req.GasPrice, req.Sleep, req.Num
		svc.RunWorkload(ctx, from, to, amount, gasLimit, gasPrice, sleep, num)
		return runWorkloadResponse{}, nil
	}
}

func makeNodeSyncProgressEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		syncing, currentBlock, highestBlock, err := svc.NodeSyncProgress(ctx)
		if err != nil {
			return nodeSyncProgressResponse{false, uint64(0), uint64(0), err.Error()}, nil
		}
		return nodeSyncProgressResponse{syncing, currentBlock, highestBlock, ""}, nil
	}
}

func makeGetBalanceEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBalanceRequest)
		address := req.Address
		balance, err := svc.GetBalance(ctx, address)
		if err != nil {
			return getBalanceResponse{int64(0), err.Error()}, nil
		}
		return getBalanceResponse{balance, ""}, nil
	}
}

func decodeGetVaultKeyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getVaultKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGenerateKeyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request generateKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeExecuteTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request executeTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeRunWorkloadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request runWorkloadRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeNodeSyncProgressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request nodeSyncProgressRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getVaultKeyRequest struct{}

type generateKeyRequest struct{}

type executeTransactionRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   int64  `json:"amount"`
	GasLimit uint64 `json:"gasLimit"`
	GasPrice int64  `json:"gasPrice"`
}

type runWorkloadRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   int64  `json:"amount"`
	GasLimit uint64 `json:"gasLimit"`
	GasPrice int64  `json:"gasPrice"`
	Sleep    int    `json:"sleep"`
	Num      int    `json:"num"`
}

type nodeSyncProgressRequest struct{}

type getBalanceRequest struct {
	Address string `json:"address"`
}

type getVaultKeyResponse struct {
	Key string `json:"key"`
	Err string `json:"err,omitempty"`
}

type generateKeyResponse struct {
	Address string `json:"address"`
	Err     string `json:"err,omitempty"`
}

type executeTransactionResponse struct {
	Err string `json:"err,omitempty"`
}

type runWorkloadResponse struct{}

type nodeSyncProgressResponse struct {
	Syncing      bool   `json:"syncing"`
	CurrentBlock uint64 `json:"currentBlock"`
	HighestBlock uint64 `json:"highestBlock"`
	Err          string `json:"err,omitempty"`
}

type getBalanceResponse struct {
	Balance int64  `json:"balance"`
	Err     string `json:"err,omitempty"`
}
