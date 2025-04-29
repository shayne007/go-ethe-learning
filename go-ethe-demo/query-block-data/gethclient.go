package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewGethClient creates a new gethclient.Client
func NewGethClient(url string) *gethclient.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}
	return gethclient.New(client)
}

// QueryNodeStats queries the node statistics
func QueryNodeStats(url string) error {
	client := NewGethClient(url)
	// Note: MemStats is only available on nodes that expose this RPC method
	stats, err := client.GetNodeInfo(context.Background())
	if err != nil {
		fmt.Println("Error getting node info:", err)
		return err
	}
	fmt.Printf("Node info: %v\n", stats)
	return nil
}
