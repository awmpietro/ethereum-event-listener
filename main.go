/**
 * @file main.go
 * @author: Arthur Mastropietro <arthur.mastropietro@gmail.com
 * @date 2018
 */

package main

import (
	"log"
	"math/big"

	"github.com/BurntSushi/toml"
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
	DbHost          string
}

// conf holds the filled Config struct
var conf *Config

// main function of the program
func main() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal("Erro ", err)
	}
	provider := new(Provider)
	provider.setUp()
	provider.listenToEvent()
	//Just for testnet
	// var nonce int64 = 6
	// p.auth.Nonce = big.NewInt(nonce)
}
