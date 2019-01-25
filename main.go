package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eximchain/eth-client/quorum"
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
		quorumAddress := "http://localhost:8545"
		quorumClient, err := quorum.Dial(quorumAddress)
		if err != nil {
			log.Fatal(err)
		}

		svc := transactionExecutorService{
			keystore:      keystore.NewKeyStore("./keystore-local", keystore.StandardScryptN, keystore.StandardScryptP),
			quorumAddress: quorumAddress,
			quorumClient:  quorumClient,
			accountCache:  make(map[string]accounts.Account),
		}

		handler := MakeRPCHandler(svc)

		http.Handle("/", accessControl(handler))

		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
