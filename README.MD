# Ethereum Event Listener

An example of a program written in Go, which listens to a Smart Contract's event


## 1 - Solidity contract to ABI
```bash
solc --abi contracts/MyContract.sol -o contracts
```

## 2 - ABI to Go package
```bash
abigen --abi contracts/MyContract.abi --pkg mycontract --out mycontract/contract.go
```
## 3 - Fill your data in config.toml