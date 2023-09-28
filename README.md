# Blockchain Quickstart

* [加密相关](./ch1-Cryptography%20Related/README.md)
* [比特币初识](./ch2-Bitcoin%20for%20Beginners/README.md)
* [区块结构](./ch3-Block%20Structure/README.md)

## DEMO Showcase

```bash
cd ch4-BTC\ Dealings
go build -o mbc main.go
./mbc -h
mbc is a very easy block chain

Usage:
  mbc [flags]
  mbc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  getBalance  FROM makes a transaction to TO
  help        Help about any command
  listAddr    List Addresses
  newWallet   Create a wallet
  printChain  Print block chain data
  send        FROM makes a single transaction to TO while MINER mines and writes to DATA

Flags:
  -h, --help   help for mbc

Use "mbc [command] --help" for more information about a command.
```
