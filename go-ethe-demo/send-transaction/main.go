package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var SK = "0xaf5ead4413ff4b78bc94191a2926ae9ccbec86ce099d65aaf469e9eb1a0fa87f"
var ADDR = "0x6177843db3138ae69679A54b95cf345ED759450d"

func main() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Sepolia testnet instead of Goerli
	cl, err := ethclient.DialContext(ctx, "https://sepolia.infura.io/v3/40a17052b04b43ad884503a6fcf9e3bc")
	if err != nil {
		panic(err)
	}
	// Send a transaction
	err = sendTransaction(cl)
	if err != nil {
		panic(err)
	}

}
