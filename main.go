/**
 * @file main.go
 * @author: Arthur Mastropietro <arthur.mastropietro@gmail.com
 * @date 2018
 */

package main

import (
	"context"
	"fmt"

	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/awmpietro/ethereum-event-listener/mycontract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Config has the configuration data needed provided by toml file
type Config struct {
	ProviderURL     string
	ContractAddress string
	PrivateKey      string
	AbiPath         string
	EventName       string
}

// Provider has the variables needed to communicate with the Ethereum RPC
type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	contractClient  *mycontract.Mycontract
	auth            *bind.TransactOpts
}

// conf holds the filled Config struct
var conf *Config

// setUp is responsible for initializing the needed vars
func (p *Provider) setUp() {
	pClient, err := ethclient.Dial(conf.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}
	p.client = pClient
	pPrivateKey, err := crypto.HexToECDSA(conf.PrivateKey) // Private key
	if err != nil {
		log.Fatal(err)
	}
	p.privateKey = pPrivateKey
	p.contractAddress = common.HexToAddress(conf.ContractAddress) //Contract Address
	pContractClient, err := mycontract.NewMycontract(p.contractAddress, p.client)
	if err != nil {
		log.Fatal(err)
	}
	p.contractClient = pContractClient
	p.auth = bind.NewKeyedTransactor(p.privateKey)
}

// main function of the program
func main() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal("Erro ", err)
	}
	provider := new(Provider)
	provider.setUp()
	// var nonce int64 = 6
	// p.auth.Nonce = big.NewInt(nonce)

	// Contract Method which triggers an Event for testing
	tx, err := provider.contractClient.Greet(provider.auth, "Hi Gopher! Event was triggered")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction TX: 0x%x\n", tx.Hash())

	// Listening to an event
	query := ethereum.FilterQuery{
		Addresses: []common.Address{provider.contractAddress},
	}
	var eventCh = make(chan types.Log)
	ctx := context.Background()
	sub, err := provider.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		log.Println("Subscribe Failed: ", err)
		return
	}
	abiPath, _ := filepath.Abs(conf.AbiPath)
	file, err := ioutil.ReadFile(abiPath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}
	// The program keeps in a loop listening to the event which is stored in eventCh channel
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err) // Error
		case log := <-eventCh:
			var contractEvent struct {
				Name  string
				Count *big.Int
			}
			err = contractAbi.Unpack(&contractEvent, conf.EventName, log.Data)
			if err != nil {
				fmt.Println("Failed to unpack:", err)
			}
			// Example of what to do with the event log data
			fmt.Println("Contract Address:", log.Address.Hex())
			fmt.Println("Name:", contractEvent.Name)
			fmt.Println("Counter:", contractEvent.Count)
		}
	}
}
