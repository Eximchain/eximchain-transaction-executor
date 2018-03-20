package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGetKeyEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//		req := request.(getKeyRequest)
		v, err := svc.GetKey(ctx)
		if err != nil {
			return getKeyResponse{v, err.Error()}, nil
		}
		return getKeyResponse{v, ""}, nil
	}
}

func makeGenerateKeyEndpoint(svc TransactionExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//		req := request.(getKeyRequest)
		v, err := svc.GenerateKey(ctx)
		if err != nil {
			return generateKeyResponse{v, err.Error()}, nil
		}
		return generateKeyResponse{v, ""}, nil
	}
}

func decodeGetKeyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getKeyRequest
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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getKeyRequest struct{}

type generateKeyRequest struct{}

type getKeyResponse struct {
	Key string `json:"key"`
	Err string `json:"err,omitempty"`
}

type generateKeyResponse struct {
	Address string `json:"address"`
	Err     string `json:"err,omitempty"`
}
