package geth

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewRPCClient() (*ethclient.Client, error) {
	endpoint := os.Getenv("GETH_HTTP_RPC_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("empty env NODE_HTTP_RPC_ENDPOINT")
	}
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewWebSocketClient() (*ethclient.Client, error) {
	endpoint := os.Getenv("GETH_WEBSOCKET_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("empty env GETH_WEBSOCKET_ENDPOINT")
	}
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}
