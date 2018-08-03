package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eximchain/go-ethereum/accounts"
	"github.com/eximchain/go-ethereum/accounts/keystore"
)

func main() {
	switch os.Args[1] {
	case "server":
		RunServerCommand(os.Args[2:])
	case "user":
		RunUserCommand(os.Args[2:])
	case "local":
		svc := transactionExecutorService{
			keystore:      keystore.NewKeyStore("./keystore-local", keystore.StandardScryptN, keystore.StandardScryptP),
			quorumAddress: "http://localhost:8545",
			accountCache:  make(map[string]accounts.Account),
		}

		handler := MakeRPCHandler(svc)

		http.Handle("/rpc", handler)

		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
