package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Sepolia testnet instead of Goerli
	cl, err := ethclient.DialContext(ctx, "https://sepolia.infura.io/v3/40a17052b04b43ad884503a6fcf9e3bc")
	if err != nil {
		panic(err)
	}

	blockNumber, err := cl.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	blockNumberBig := big.NewInt(int64(blockNumber))

	eventSignatureBytes := []byte("Transfer(address,address,uint256)")
	eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)

	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(10)),
		ToBlock:   blockNumberBig,
		Topics: [][]common.Hash{
			{eventSignaturehash},
		},
	}

	logs, err := cl.FilterLogs(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
	}

}
