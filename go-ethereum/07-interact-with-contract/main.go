package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	todo "go-ethereum-tutorial/gen"
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
	b, err := os.ReadFile("./wallet/accountA/UTC--2022-12-23T08-52-06.653324900Z--e5a7e0a934851e1f61625953817fdc3b0c3411e1")
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

	contractAddr := common.HexToAddress("0xF76a652F4a980B27933edeee1063Ec78610bD452")

	// create Todo instance
	t, err := todo.NewTodo(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// create tx signer
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 3000000
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	/*add todo*/
	// tx, err := t.Add(auth, "My Second Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*update todo*/
	// tx, err := t.Update(auth, big.NewInt(0), "Update First Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*toggle todo*/
	// tx, err := t.Toggle(auth, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*remove todo*/
	// tx, err := t.Remove(auth, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Tx Hash:", tx.Hash().Hex())

	// get todo list
	tasks, err := t.List(&bind.CallOpts{
		From: addr,
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, task := range tasks {
		fmt.Println("------------------------------")
		fmt.Printf("Task #%d\n", i)
		fmt.Println("Content:", task.Content)
		fmt.Println("Done:", task.Status)
	}
	fmt.Println("------------------------------")
}
