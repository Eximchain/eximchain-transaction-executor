# server

```sh
curl -XPOST -H 'Authorization: $TOKEN' -d'{}' localhost:8080/get-vault-key
curl -XPOST -H 'Authorization: $TOKEN' -d'{}' localhost:8080/generate-key
curl -XPOST -H 'Authorization: $TOKEN' -d'{}' localhost:8080/node-sync-progress
curl -XPOST -H 'Authorization: $TOKEN' -d'{"address":"$ADDRESS"}' localhost:8080/get-balance
curl -XPOST -H 'Authorization: $TOKEN' -d'{"from":"$FROM","to":$TO,"amount":$AMOUNT}' localhost:8080/execute-transaction
```

## user

```sh
./eximchain user --email zuo.wang@enuma.io
./eximchain user --email zuo.wang@enuma.io --update
./eximchain user --list
./eximchain user --email zuo.wang@enuma.io --delete
```

| endpoint            | rpc_method          |
| ------------------- | ------------------- |
| get-vault-key       | eth_accounts        |
| generate-key        | personal_newAccount |
| node-sync-progress  | eth_syncing         |
| get-balance         | eth_getBalacne      |
| execute-transaction | eth_sendTransaction |
