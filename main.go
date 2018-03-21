package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/eximchain/eth-client/quorum"
	"github.com/eximchain/go-ethereum/accounts/keystore"

	httptransport "github.com/go-kit/kit/transport/http"
	vault "github.com/hashicorp/vault/api"
	awsauth "github.com/hashicorp/vault/builtin/credential/aws"
)

// Expects to be running in EC2
func GetRole() (string, error) {
	svc := ec2metadata.New(session.New())
	iam, err := svc.IAMInfo()
	if err != nil {
		return "", err
	}
	// Our instance profile conveniently has the same name as the role
	profile := iam.InstanceProfileArn
	splitArn := strings.Split(profile, "/")
	if len(splitArn) < 2 {
		return "", fmt.Errorf("no / character found in instance profile ARN")
	}
	role := splitArn[1]
	return role, nil
}

func LoginAws(v *vault.Client) (string, error) {
	loginData, err := awsauth.GenerateLoginData("", "", "", "")
	if err != nil {
		return "", err
	}
	if loginData == nil {
		return "", fmt.Errorf("got nil response from GenerateLoginData")
	}

	role, err := GetRole()
	if err != nil {
		return "", err
	}
	loginData["role"] = role

	path := "auth/aws/login"

	secret, err := v.Logical().Write(path, loginData)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", fmt.Errorf("empty response from credential provider")
	}
	if secret.Auth == nil {
		return "", fmt.Errorf("auth secret has no auth data")
	}

	token := secret.Auth.ClientToken
	return token, nil
}

func main() {
	vaultAddressFlag := flag.String("vault-address", "http://127.0.0.1:8200", "The address at which vault can be accessed")
	quorumAddressFlag := flag.String("quorum-address", "http://127.0.0.1:8545", "The address at which the quorum node can be reached")
	authTokenFlag := flag.String("auth-token", "", "An auth token to use instead of AWS authorization, for help with testing")
	keyDirFlag := flag.String("keystore", "/home/ubuntu/.ethereum/keystore", "The directory to use as a keystore")
	flag.Parse()

	// Vault client setup
	vaultCFG := vault.DefaultConfig()
	// TODO: Put this behind a command line flag so we can test
	vaultCFG.Address = *vaultAddressFlag

	var err error
	vaultClient, err := vault.NewClient(vaultCFG)
	if err != nil {
		log.Fatal(err)
	}

	var token string
	if *authTokenFlag != "" {
		token = *authTokenFlag
	} else {
		token, err = LoginAws(vaultClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	vaultClient.SetToken(token)

	// Quorum client setup
	quorumAddress := *quorumAddressFlag
	quorumClient, err := quorum.Dial(quorumAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Keystore setup
	gethKeyDir := *keyDirFlag
	gethKeystore := keystore.NewKeyStore(gethKeyDir, keystore.StandardScryptN, keystore.StandardScryptP)

	svc := transactionExecutorService{vaultClient: vaultClient, keystore: gethKeystore, quorumClient: &quorumClient}

	getVaultKeyHandler := httptransport.NewServer(
		makeGetVaultKeyEndpoint(svc),
		decodeGetVaultKeyRequest,
		encodeResponse,
	)

	generateKeyHandler := httptransport.NewServer(
		makeGenerateKeyEndpoint(svc),
		decodeGenerateKeyRequest,
		encodeResponse,
	)

	http.Handle("/get-vault-key", getVaultKeyHandler)
	http.Handle("/generate-key", generateKeyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
