package main

type jsonRpcRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Id      string   `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type jsonRpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  string `json:"result"`
}

type RPCParams = []string

type RPCTransactionParams = []RPCTransaction

type RPCTransaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	Nonce    string `json:"nonce"`
}
