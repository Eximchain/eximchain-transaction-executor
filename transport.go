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
		from, to := req.From, req.To
		err := svc.ExecuteTransaction(ctx, from, to)
		if err != nil {
			return executeTransactionResponse{err.Error()}, nil
		}
		return executeTransactionResponse{""}, nil
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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getVaultKeyRequest struct{}

type generateKeyRequest struct{}

type executeTransactionRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
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
