package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	todo "go-ethereum-tutorial/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
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
	// read encrypted json file
	b, err := ioutil.ReadFile("./wallet/accountA/UTC--2022-12-23T08-52-06.653324900Z--e5a7e0a934851e1f61625953817fdc3b0c3411e1")
	if err != nil {
		log.Fatalf("Error to read key file: %v", err)
	}

	// decrypt json file and get private key of account A
	key, err := keystore.DecryptKey(b, "passwordA")
	if err != nil {
		log.Fatalf("Error to decrypt key file: %v", err)
	}

	// connect to ethereum goerli network
	client, err := ethclient.DialContext(context.Background(), infuraGoerliURL)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// get address and nonce of account A
	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}

	// get suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// get network id
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create tx signer
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 3000000
	auth.Nonce = big.NewInt(int64(nonce))

	// deploy todo smart contract
	contractAddr, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------")
	fmt.Println("Contract Address:", contractAddr.Hex())
	fmt.Println("Tx Hash:", tx.Hash().Hex())
	fmt.Println("------------------------------")
}
