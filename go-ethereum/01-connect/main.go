package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var infuraURL string
var ganacheURL string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error to load environment variables: %v", err)
	}

	infuraURL = os.Getenv("INFURA_MAINNET_URL")
	if infuraURL == "" {
		log.Fatal("Error to load environment variables")
	}

	ganacheURL = os.Getenv("GANACHE_URL")
	if ganacheURL == "" {
		log.Fatal("Error to load environment variables")
	}
}

func main() {
	// connect to ethereum client
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	// get latest block on network
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println("Block Number:", block.Number())

	hexAddr := "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d" // address of CryptoKitties Smart Contract
	address := common.HexToAddress(hexAddr)

	// get latest balance of address
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance: %v", err)
	}

	fmt.Println("Balance:", balance)

	// convert big number to ether
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance to Ether:", balanceEther)
}
