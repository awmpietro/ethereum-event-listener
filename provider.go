package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/awmpietro/ethereum-event-listener/mycontract"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	mgo "gopkg.in/mgo.v2"
)

// Provider has the variables needed to communicate with the Ethereum RPC
type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	contractClient  *mycontract.Mycontract
	auth            *bind.TransactOpts
}

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

// Contract method used for testing to trigger an event
func (p *Provider) greet() {
	tx, err := p.contractClient.Greet(p.auth, "Hi Gopher! Event was triggered")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction TX: 0x%x\n", tx.Hash())
}

// listenToEvent is the function responsible for keep listening an event
func (p *Provider) listenToEvent() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{p.contractAddress},
		FromBlock: conf.BlockNumber,
	}
	ctx := context.Background()

	var eventCh = make(chan types.Log)

	sub, err := p.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		log.Println("Subscribe Failed: ", err)
		return
	}
	// Check events logs
	logs, err := p.client.FilterLogs(ctx, query)
	if err != nil {
		fmt.Println("Filter Logs: ", err)
	}

	event := new(Event)
	event.Session, err = mgo.Dial(conf.DbHost)
	if err != nil {
		fmt.Println("Failed to start Mongo session:", err)
	}
	event.checkEventLog(logs)
	event.Session.Close()

	// The program keeps in a loop listening to the event which is stored in eventCh channel
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err) // Error
		case log := <-eventCh:
			event.incrementCounter()
			event.forwardEvents(log)
		}
	}
}
