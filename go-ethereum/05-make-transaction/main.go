package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	fmt.Println("Balance of A before transaction:", balanceA)
	fmt.Println("Balance of B before transaction:", balanceB)

	nonceA, err := client.PendingNonceAt(context.Background(), accountA)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(100000000000000000)                      // 0.1 ether
	gasPrice, err := client.SuggestGasPrice(context.Background()) // get suggested gas price
	if err != nil {
		log.Fatal(err)
	}

	// create transaction that sends 0.1 ether from account A to account B
	tx := types.NewTransaction(nonceA, accountB, amount, 21000, gasPrice, nil)

	// get network id
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

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

	// sign on transaction with private key of account A
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// send transaction
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent:", tx.Hash().Hex())

	time.Sleep(30 * time.Second)

	// get balance of account A
	balanceA, err = client.BalanceAt(context.Background(), accountA, nil)
	if err != nil {
		log.Fatal(err)
	}

	// get balance of account B
	balanceB, err = client.BalanceAt(context.Background(), accountB, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance of A after transaction:", balanceA)
	fmt.Println("Balance of B after transaction:", balanceB)
}

/*
$ go run ./05-make-transaction/
Balance of A before transaction: 200000000000000000
Balance of B before transaction: 0
tx sent: 0xc29f5ce22087689394c06b65e484c094d2a329c9c906daf54c252fb19dbc1ab8
Balance of A after transaction: 99999202101535000
Balance of B after transaction: 100000000000000000
*/
