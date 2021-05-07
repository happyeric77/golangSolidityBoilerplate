package ContractDeploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/happyeric77/golangSolidityBoilerplate/contracts"
)

func Deploy(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*contracts.Contracts, common.Address, common.Hash, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type(contractDeploy.Deploy): publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Depoly contract by: %x\nNonce: %v\n", fromAddress, nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gas Price: %v\n", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	addr, tx, contract, err := contracts.DeployContracts(auth, client)
	if err != nil {
		return nil, common.HexToAddress("0x0"), common.HexToHash("0x0"), err
	}
	fmt.Printf("Deployed CA: %x\nTx address: %v\n", addr, tx.Hash())
	return contract, addr, tx.Hash(), nil
}
