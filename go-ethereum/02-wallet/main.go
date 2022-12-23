package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// generate ECDSA private key
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Error to generate private key: %v", err)
	}

	// convert ECDSA private key to readable format
	pvkToBytes := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pvkToBytes))

	// get public key paired with private key
	pbkToBytes := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(pbkToBytes))

	// generate wallet address from public key
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
