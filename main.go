package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/happyeric77/golangSolidityBoilerplate/ContractDeploy"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")
	client, err := ethclient.Dial(fmt.Sprintf("https://kovan.infura.io/v3/%s", os.Getenv("INFURA_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	contractInstance, ca, txa, err := ContractDeploy.Deploy(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	_ = contractInstance
	_ = ca
	_ = txa
}
