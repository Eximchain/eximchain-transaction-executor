package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"

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

func LoginAws(v *vault.Client) error {
	loginData, err := awsauth.GenerateLoginData("", "", "", "")
	if err != nil {
		return err
	}
	if loginData == nil {
		return fmt.Errorf("got nil response from GenerateLoginData")
	}

	role, err := GetRole()
	if err != nil {
		return err
	}
	loginData["role"] = role

	path := "auth/aws/login"

	secret, err := v.Logical().Write(path, loginData)
	if err != nil {
		return err
	}
	if secret == nil {
		return fmt.Errorf("empty response from credential provider")
	}
	if secret.Auth == nil {
		return fmt.Errorf("auth secret has no auth data")
	}

	token := secret.Auth.ClientToken
	v.SetToken(token)

	return nil
}

func main() {
	vaultAddressFlag := flag.String("vault-address", "http://127.0.0.1:8200", "The address at which vault can be accessed")
	flag.Parse()

	vaultCFG := vault.DefaultConfig()
	// TODO: Put this behind a command line flag so we can test
	vaultCFG.Address = *vaultAddressFlag

	var err error
	vaultClient, err := vault.NewClient(vaultCFG)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Put this behind a command line flag so we can test
	if true {
		err := LoginAws(vaultClient)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Magic token for dev vault server
		vaultClient.SetToken("63b9d575-d4d3-bb0c-8ac0-d2562d142f6c")
	}

	svc := transactionExecutorService{vaultClient: vaultClient}

	getKeyHandler := httptransport.NewServer(
		makeGetKeyEndpoint(svc),
		decodeGetKeyRequest,
		encodeResponse,
	)

	http.Handle("/get-key", getKeyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
