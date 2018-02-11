package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ContractEvent *ContractEvent
	EventCounter  *EventCounter
	Session       *mgo.Session
	ContractAbi   abi.ABI
}

// ContractEvent maps an event data in a struct
type ContractEvent struct {
	Name  string
	Count *big.Int
}

// EventCounter maps the Mongo collection used to keep track of the events
type EventCounter struct {
	Counter int
}

// checkEventLog check if all events emitted are in track in Mongo. If no, do a routine to update and
// forward non tracked events
func (e *Event) checkEventLog(logs []types.Log) {
	c := e.Session.DB(conf.DbName).C(conf.DbCollection) //Db=eventsCounter, Collection/Table=counter
	err := c.Find(bson.M{}).One(&e.EventCounter)
	if err != nil {
		log.Fatal("MongoDB Find: ", err)
	}
	if e.EventCounter.Counter != len(logs) {
		// Events are desynchronized, do something
		left := len(logs) - e.EventCounter.Counter
		for i := (len(logs) - left); i < len(logs); i++ {
			e.forwardEvents(logs[i])
		}
		e.updateCounter(len(logs))
	}
	return
}

// updateCounter updates the events counter in Mongo
func (e *Event) updateCounter(current int) {
	c := e.Session.DB(conf.DbCollection).C(conf.DbName)
	e.EventCounter.Counter = current
	err := c.Update(bson.M{}, &e.EventCounter)
	if err != nil {
		fmt.Println("Can't update MongoDB")
		return
	}
}

//incrementCounter increments the events counter in Mongo
func (e *Event) incrementCounter() {
	c := e.Session.DB(conf.DbName).C(conf.DbCollection)
	err := c.Find(bson.M{}).One(&e.EventCounter)
	if err != nil {
		log.Fatal(err)
	}
	e.EventCounter.Counter++
	e.updateCounter(e.EventCounter.Counter)
}

// forwaredEvents will forward the emitted events as they arrive
func (e *Event) forwardEvents(log types.Log) {
	fmt.Println("Mais um")
	//e.getContracAbi()
	/*
		Forward events somewhere
		Example of what to do with the event log data
		err = e.contractAbi.Unpack(&e.contractEvent, conf.EventName, log.Data)
		if err != nil {
			fmt.Println("Failed to unpack:", err)
		}
		fmt.Println("Contract Address:", log.Address.Hex())
		fmt.Println("Name:", e.contractEvent.Name)
		fmt.Println("Counter:", e.contractEvent.Count)
	*/
}

//getContractAbi returs the contract ABI
func (e *Event) getContractAbi() {
	abiPath, _ := filepath.Abs(conf.AbiPath)
	file, err := ioutil.ReadFile(abiPath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
	}
	e.ContractAbi, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}
}
