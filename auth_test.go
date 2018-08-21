package main

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

func testRpc(t *testing.T, token string) *http.Response {
	jsonStr := []byte(`{"jsonrpc":"2.0","id":2,"method":"eth_syncing","params":[]}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/rpc", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("NewRequest %v", err)
	}
	req.Header.Add("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	return resp
}

func TestHttpAuthXfail(t *testing.T) {
	resp := testRpc(t, "asdf")
	defer resp.Body.Close()
	if resp.Status != "401 Unauthorized" {
		t.Fatalf("response Status: %v, expected 401 Unauthorized", resp.Status)
	}
}

func TestHttpAuth(t *testing.T) {
	// Grab the first user token from the server over IPC
	args := []string{"--list"}
	output := sendIPC(args)
	fields := strings.Fields(output)
	if len(fields) < 2 {
		t.Fatal("No users in database; please add at least one token.")
	}

	token := fields[1]
	t.Logf("Using token %s", token)
	resp := testRpc(t, token)
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Fatalf("response Status: %v, expected 200 OK", resp.Status)
	}
}
