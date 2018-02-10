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
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Config has the configuration data needed provided by toml file
type Config struct {
	ProviderURL     string
	ContractAddress string
	PrivateKey      string
	AbiPath         string
	EventName       string
	BlockNumber     *big.Int
	DbName          string
	DbCollection    string
}

// Provider has the variables needed to communicate with the Ethereum RPC
type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	contractClient  *mycontract.Mycontract
	auth            *bind.TransactOpts
}

// ContractEvent maps from the events data in a struct
type ContractEvent struct {
	Name  string
	Count *big.Int
}

// EventCounter maps the Mongo collection used to keep track of the events
type EventCounter struct {
	Counter int
}

// conf holds the filled Config struct
var conf *Config
var session *mgo.Session

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

// checkEventLog check if all events emitted are in track in Mongo. If no, do a routine to update and
// forward non tracked events
func checkEventLog(logs []types.Log) {
	c := session.DB(conf.DbCollection).C(conf.DbName) //Db=eventsCounter, Collection/Table=counter
	var eventCounter EventCounter
	err := c.Find(bson.M{}).One(&eventCounter)
	if err != nil {
		log.Fatal("MongoDB Find: ", err)
	}
	if eventCounter.Counter != len(logs) {
		// Events are desynchronized, do something
		left := len(logs) - eventCounter.Counter
		for i := (len(logs) - left); i < len(logs); i++ {
			forwardEvents(logs[i])
		}
		updateCounter(len(logs))
	}
	return
}

// updateCounter updates the events counter in Mongo
func updateCounter(current int) {
	c := session.DB(conf.DbCollection).C(conf.DbName)
	var eventCounter EventCounter
	eventCounter.Counter = current
	err := c.Update(bson.M{}, &eventCounter)
	if err != nil {
		fmt.Println("Can't update MongoDB")
		return
	}
}

//incrementCounter increments the events counter in Mongo
func incrementCounter() {
	c := session.DB(conf.DbCollection).C(conf.DbName)
	var eventCounter EventCounter
	err := c.Find(bson.M{}).One(&eventCounter)
	if err != nil {
		log.Fatal(err)
	}
	eventCounter.Counter++
	updateCounter(eventCounter.Counter)
}

// forwaredEvents will forward the emitted events as they arrive
func forwardEvents(log types.Log) {
	fmt.Println("Mais um")
	//contractAbi := getContractAbi()
	/*
		Forward events somewhere
		Example of what to do with the event log data
		contractEvent := ContractEvent{}
		err = contractAbi.Unpack(&contractEvent, conf.EventName, log.Data)
		if err != nil {
			fmt.Println("Failed to unpack:", err)
		}
		fmt.Println("Contract Address:", log.Address.Hex())
		fmt.Println("Name:", contractEvent.Name)
		fmt.Println("Counter:", contractEvent.Count)
	*/
}

//getContractAbi returs the contract ABI
func getContractAbi() abi.ABI {
	abiPath, _ := filepath.Abs(conf.AbiPath)
	file, err := ioutil.ReadFile(abiPath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}
	return contractAbi
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
	/*
		tx, err := provider.contractClient.Greet(provider.auth, "Hi Gopher! Event was triggered")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Transaction TX: 0x%x\n", tx.Hash())
	*/

	// Listening to an event
	query := ethereum.FilterQuery{
		Addresses: []common.Address{provider.contractAddress},
		FromBlock: conf.BlockNumber,
	}
	ctx := context.Background()

	var eventCh = make(chan types.Log)

	sub, err := provider.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		log.Println("Subscribe Failed: ", err)
		return
	}
	// Check events logs
	logs, err := provider.client.FilterLogs(ctx, query)
	if err != nil {
		fmt.Println("Filter Logs: ", err)
	}
	session, err = mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Failed to start Mongo session:", err)
	}
	checkEventLog(logs)
	session.Close()

	// The program keeps in a loop listening to the event which is stored in eventCh channel
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err) // Error
		case log := <-eventCh:
			incrementCounter()
			forwardEvents(log)
		}
	}
}
