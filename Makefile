all: *.go
	go build

server: *.go
	go run common.go auth.go db.go main.go rpc.go server.go service.go transport.go user.go server -auth-token test -quorum-address http://127.0.0.1:7545 -keystore ./keystore

fmt: *.go
	gofmt -w *.go

test: *.go
	rm -rf eximchain_test.db
	go test -v -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm c.out coverage.html
