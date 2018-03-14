package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	vault "github.com/hashicorp/vault/api"
)

func main() {
	vaultCFG := vault.DefaultConfig()
	vaultCFG.Address = "http://127.0.0.1:8200"

	var err error
	vaultClient, err := vault.NewClient(vaultCFG)
	if err != nil {
		log.Fatal(err)
	}

	// Magic token for dev vault server
	// TODO: Replace with AWS authentication
	vaultClient.SetToken("63b9d575-d4d3-bb0c-8ac0-d2562d142f6c")

	svc := transactionExecutorService{vaultClient: vaultClient}

	getKeyHandler := httptransport.NewServer(
		makeGetKeyEndpoint(svc),
		decodeGetKeyRequest,
		encodeResponse,
	)

	http.Handle("/get-key", getKeyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
