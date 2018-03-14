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

func decodeGetKeyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getKeyRequest struct{}

type getKeyResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
