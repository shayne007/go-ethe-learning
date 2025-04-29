package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	const BLOCK_CHAIN_PLATFORM_URL = "https://sepolia.infura.io/v3/40a17052b04b43ad884503a6fcf9e3bc"
	cl := NewEthClient(BLOCK_CHAIN_PLATFORM_URL)

	// get the latest block number
	header, err := cl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := header.Number
	fmt.Println(blockNumber) // 3467113
	// get the block by number, if nil, it will get the latest block
	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Uint64())     // 3467113
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 855bba72d94902cc385042661498a415979b7b6ee9ba4b9168f17a1f933c151b
	fmt.Println(len(block.Transactions()))   // 144
	count, err := cl.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count) // 144

	fmt.Println("start to query node stats: ")
	QueryNodeStats(BLOCK_CHAIN_PLATFORM_URL)
}
