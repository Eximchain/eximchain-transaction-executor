# Test Cases

```sh
curl -XPOST -d'{}' localhost:8080/get-vault-key
curl -XPOST -d'{}' localhost:8080/generate-key
curl -XPOST -d'{}' localhost:8080/node-sync-progress
curl -XPOST -d'{"address":"$ADDRESS"}' localhost:8080/get-balance
curl -XPOST -d'{"from":"$FROM","to":$TO,"amount":$AMOUNT}' localhost:8080/execute-transaction
```
