package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var infuraGoerliURL string
var ganacheURL string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error to load environment variables: %v", err)
	}

	infuraGoerliURL = os.Getenv("INFURA_GOERLI_URL")
	if infuraGoerliURL == "" {
		log.Fatal("Error to load environment variables")
	}

	ganacheURL = os.Getenv("GANACHE_URL")
	if ganacheURL == "" {
		log.Fatal("Error to load environment variables")
	}
}

func main() {
	// generate account A and B

	// ks := keystore.NewKeyStore("./wallet/accountA", keystore.StandardScryptN, keystore.StandardScryptP)
	// passwordA := "passwordA"
	// _, err := ks.NewAccount(passwordA)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ks = keystore.NewKeyStore("./wallet/accountB", keystore.StandardScryptN, keystore.StandardScryptP)
	// passwordB := "passwordB"
	// _, err = ks.NewAccount(passwordB)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	accountA := common.HexToAddress("e5a7e0a934851e1f61625953817fdc3b0c3411e1")
	accountB := common.HexToAddress("6da3d3e0edbc38ad00d82102ebcdade65772b4fb")

	// connect to ethereum goerli network
	client, err := ethclient.DialContext(context.Background(), infuraGoerliURL)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// get balance of account A
	balanceA, err := client.BalanceAt(context.Background(), accountA, nil)
	if err != nil {
		log.Fatal(err)
	}

	// get balance of account B
	balanceB, err := client.BalanceAt(context.Background(), accountB, nil)
	if err != nil {
		log.Fatal(err)
	}

	// accountA got 0.2 ethers from goerli faucet
	fmt.Println("Balance of A:", balanceA)
	fmt.Println("Balance of B:", balanceB)
}
