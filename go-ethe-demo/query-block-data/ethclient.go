package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEthClient creates a new ethclient.Client
func NewEthClient(url string) *ethclient.Client {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := ethclient.DialContext(ctx, url)
	if err != nil {
		panic(err)
	}
	return cl
}
