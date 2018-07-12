package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func makeRequest(data string) string {
	svc := transactionExecutorService{}
	handler := MakeRPCHandler(svc, "http://localhost:7545")
	rpcServer := httptest.NewServer(handler)
	defer rpcServer.Close()

	resp, _ := http.Post(rpcServer.URL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(data)))
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func ExampleWeb3ClientVersion() {
	data := `{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"EthereumJS TestRPC/v2.1.0/ethereum-js"}
}

func ExampleWeb3Sha3() {
	data := `{"jsonrpc":"2.0","method":"web3_sha3","params":["0x68656c6c6f20776f726c64"],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"0x47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad"}
}

func ExampleNetVersion() {
	data := `{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"5777"}
}

func ExampleNetPeerCount() {
	data := `{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":0}
}

func ExampleNetListening() {
	data := `{"jsonrpc":"2.0","method":"net_listening","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":true}
}

func ExampleEthProtocolVersion() {
	data := `{"jsonrpc":"2.0","method":"eth_protocolVersion","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"63"}
}

func ExampleEthSyncing() {
	data := `{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":false}
}

func ExampleShhVersion() {
	data := `{"jsonrpc":"2.0","method":"shh_version","params":[],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"2"}
}

func ExampleEthGetBalance() {
	data := `{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x9193d626a1A3668AAdeaFF4fda44A3a52A784021", "latest"],"id":1}`
	fmt.Println(makeRequest(data))
	// Output:
	// {"jsonrpc":"2.0","result":"0x56bc75e2d63100000"}
}
