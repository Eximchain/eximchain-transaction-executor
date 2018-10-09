# Basic Usage

Basic usage example when running in AWS from [the terraform configuration](https://github.com/Eximchain/terraform-aws-eximchain-tx-executor)

```sh
# Create a user
/opt/transaction-executor/go/bin/eximchain-transaction-executor user -email louis@eximchain.com -update

# Store the token
TOKEN=<Token from previous command>

# Make an RPC call
curl -XPOST -H "Authorization: $TOKEN" -d'{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' localhost:8080/rpc
```

# Example Commands

## Server

```sh
curl -XPOST -H "Authorization: $TOKEN" -d'{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' localhost:8080/rpc
```

## User

```sh
./eximchain user --email zuo.wang@enuma.io
./eximchain user --email zuo.wang@enuma.io --update
./eximchain user --list
./eximchain user --email zuo.wang@enuma.io --delete
```

## Endpoints

| endpoint            | rpc_method          |
| ------------------- | ------------------- |
| rpc                 | all                 |
