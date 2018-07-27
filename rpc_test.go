package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eximchain/go-ethereum/accounts"
	"github.com/eximchain/go-ethereum/accounts/keystore"
)

type Response struct {
	Id      string      `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
}

func makeRequest(data string) Response {
	svc := transactionExecutorService{
		keystore:      keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP),
		quorumAddress: "http://localhost:8545",
		accountCache:  make(map[string]accounts.Account),
	}

	handler := MakeRPCHandler(svc)
	rpcServer := httptest.NewServer(handler)
	defer rpcServer.Close()

	resp, _ := http.Post(rpcServer.URL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(data)))
	body, _ := ioutil.ReadAll(resp.Body)
	var res Response
	json.Unmarshal([]byte(body), &res)
	return res
}

func NewTestAccount() string {
	data := `{"jsonrpc":"2.0","method":"personal_newAccount","params":[],"id":1}`
	res := makeRequest(data)

	return res.Result.(string)
}

func TestPersonalNewAccount(t *testing.T) {
	account := NewTestAccount()

	if account == "" {
		t.Error("personal_newAccount error")
	}
}

/*
func TestEthAccounts(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_accounts error")
	}
}
*/

func TestEthSendTransaction(t *testing.T) {
	account1 := NewTestAccount()
	account2 := NewTestAccount()

	data := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{
  "from": "` + account1 + `",
  "to": "` + account2 + `",
  "gas": "0x76c0", 
  "gasPrice": "0x9184e72a000", 
  "value": "0x9184e72a"
}],"id":1}`
	res := makeRequest(data)

	log.Println(res)

	if res.Result == nil {
		t.Error("eth_sendTransaction error")
	}
}

func TestWeb3ClientVersion(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("web3_clientVersion error")
	}
}

func TestWeb3Sha3(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"web3_sha3","params":["0x68656c6c6f20776f726c64"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("web3_sha3 error")
	}
}

func TestNetVersion(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("net_version error")
	}
}

func TestNetPeerCount(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("net_peerCount error")
	}
}

func TestNetListening(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"net_listening","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("net_listening error")
	}
}

func TestEthProtocolVersion(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_protocolVersion","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_protocolVersion error")
	}
}

func TestEthSyncing(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_syncing error")
	}
}

func TestEthCoinbase(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_coinbase","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_coinbase error")
	}
}

func TestEthMining(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_mining","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_mining error")
	}
}

func TestEthHashrate(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_hashrate","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_hashrate error")
	}
}

func TestEthGasPrice(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_gasPrice","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_gasPrice error")
	}
}

func TestEthBlockNumber(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_blockNumber error")
	}
}

func TestEthGetBalance(t *testing.T) {
	account := NewTestAccount()
	data := `{"jsonrpc":"2.0","method":"eth_getBalance","params":["` + account + `","latest"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getBalance error")
	}
}

func TestEthGetStorageAt(t *testing.T) {
	account := NewTestAccount()
	data := `{"jsonrpc":"2.0","method":"eth_getStorageAt","params":["` + account + `","0x0","latest"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getStorageAt error")
	}
}

func TestEthGetTransactionCount(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["0x9193d626a1A3668AAdeaFF4fda44A3a52A784021","latest"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getTransactionCount error")
	}
}

func TestEthGetBlockTransactionCountByHash(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByHash","params":["0x1124a47ba6ed3cdb24604270cd2c893492948c29576c68f4ec3d6c919bdf82b0"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getBlockTransactionCountByHash error")
	}
}

func TestEthGetBlockTransactionCountByNumber(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByNumber","params":["0x0"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getBlockTransactionCountByNumber error")
	}
}

func TestEthGetUncleCountByBlockHash(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getUncleCountByBlockHash","params":["0x1124a47ba6ed3cdb24604270cd2c893492948c29576c68f4ec3d6c919bdf82b0"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getUncleCountByBlockHash error")
	}
}

func TestEthGetUncleCountByBlockNumber(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getUncleCountByBlockNumber","params":["0x0"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getUncleCountByBlockNumber error")
	}
}

func TestEthGetCode(t *testing.T) {
	account := NewTestAccount()
	data := `{"jsonrpc":"2.0","method":"eth_getCode","params":["` + account + `","latest"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getCode error")
	}
}

func TestEthSign(t *testing.T) {
	account := NewTestAccount()
	data := `{"jsonrpc":"2.0","method":"eth_sign","params":["` + account + `","0xdeadbeaf"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_sign error")
	}
}

func TestEthCall(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_call","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_call error")
	}
}

func TestEthEstimateGas(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_estimateGas","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_estimateGas error")
	}
}

func TestEthGetBlockByHash(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0x1124a47ba6ed3cdb24604270cd2c893492948c29576c68f4ec3d6c919bdf82b0",true],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getBlockByHash error")
	}
}

func TestEthGetBlockByNumber(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0",true],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getBlockByNumber error")
	}
}

func TestEthGetTransactionByHash(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getTransactionByHash error")
	}
}

func TestEthGetTransactionByBlockHashAndIndex(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getTransactionByBlockHashAndIndex","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getTransactionByBlockHashAndIndex error")
	}
}

func TestEthGetTransactionByBlockNumberAndIndex(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getTransactionByBlockNumberAndIndex","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getTransactionByBlockNumberAndIndex error")
	}
}

func TestEthGetTransactionReceipt(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getTransactionReceipt error")
	}
}

func TestEthGetUncleByBlockHashAndIndex(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getUncleByBlockHashAndIndex","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getUncleByBlockHashAndIndex error")
	}
}

func TestEthGetUncleByBlockNumberAndIndex(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getUncleByBlockNumberAndIndex","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getUncleByBlockNumberAndIndex error")
	}
}

func TestEthNewFilter(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_newFilter","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_newFilter error")
	}
}

func TestEthNewBlockFilter(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_newBlockFilter","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_newBlockFilter error")
	}
}

func TestEthNewPendingTransactionFilter(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_newPendingTransactionFilter","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_newPendingTransactionFilter error")
	}
}

func TestEthUninstallFilter(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_uninstallFilter","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_uninstallFilter error")
	}
}

func TestEthGetFilterChanges(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getFilterChanges","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getFilterChanges error")
	}
}

func TestEthGetFilterLogs(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getFilterLogs","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getFilterLogs error")
	}
}

func TestEthGetLogs(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getLogs error")
	}
}

func TestEthGetWork(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_getWork","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_getWork error")
	}
}

func TestEthSubmitWork(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_submitWork","params":[],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_submitWork error")
	}
}

func TestEthSubmitHashrate(t *testing.T) {
	data := `{"jsonrpc":"2.0","method":"eth_submitHashrate","params":["0x0000000000000000000000000000000000000000000000000000000000500000","0x59daa26581d0acd1fce254fb7e85952f4c09d0915afd33d3886cd914bc7d283c"],"id":1}`
	res := makeRequest(data)

	if res.Result == nil {
		t.Error("eth_submitHashrate error")
	}
}
