package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/happyeric77/golangSolidityBoilerplate/contractCtl"
	"github.com/joho/godotenv"
)

func loadChainInfo() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial(fmt.Sprintf("https://kovan.infura.io/v3/%s", os.Getenv("INFURA_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, nil, err
	}
	return client, privateKey, nil
}

func main() {
	godotenv.Load("./.env")

	client, privateKey, err := loadChainInfo()
	_ = client
	_ = privateKey

	/**************************
	** Deploy Smart Contract **
	**************************/
	// TODO: To implement javascript's await-like feature, need to add go-channel mechanism.

	// contractInstance, ca, txa, err := contractCtl.Deploy(client, privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = contractInstance
	// _ = ca
	// _ = txa

	/*********************************
	** Load existing Smart Contract **
	**********************************/

	contractInstance, err := contractCtl.Load(client)
	if err != nil {
		log.Fatal(err)
	}
	// _ = contractInstance

	/*******************************
	** Write tx to Smart Contract **
	********************************/

	// TODO: To implement javascript's await-like feature, need to add go-channel mechanism.
	// txOpts, err := contractCtl.TxOps(client, privateKey, 0)

	// txHash, err := contractInstance.SetItem(txOpts, "Chaange on-Chain smart contract sataus")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Tx Hash: %x\n", txHash.Hash())

	/*****************************************
	** Call Smart Contract Read-only method **
	*****************************************/

	itemStatus, err := contractInstance.Item(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Item value: %s\n", itemStatus)
}
