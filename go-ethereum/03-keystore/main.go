package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// create wallet as encrypted json file
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatalf("Error to generate account: %v", err)
	// }

	// fmt.Println(account.Address)

	// read encrypted json file and decrypt it
	b, err := ioutil.ReadFile("./wallet/UTC--2022-12-23T08-27-33.530618700Z--fc39a077becac16817d957b102d2818282b4a85f")
	if err != nil {
		log.Fatalf("Error to read key file: %v", err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatalf("Error to decrypt key file: %v", err)
	}

	pvkToBytes := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(pvkToBytes))

	pbkToBytes := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println(hexutil.Encode(pbkToBytes))

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println(address.Hex())
}
