# Deploy and Interact with Smart Contract in Go

## 1. Download go-ethereum (geth)

- check both geth and development tools and download them

[download link](https://geth.ethereum.org/downloads)

## 2. Compile Smart Contract using Docker

1. download `ethereum/solc:stable` docker image

```bash
$ docker pull ethereum/solc:stable
```

2. compile `todo.sol` (used powershell on windows)

```shell
$ docker run -v "$(pwd):/tmp" ethereum/solc:stable -o /tmp/build --abi --bin /tmp/contract/todo.sol
> Compiler run successful. Artifact(s) can be found in directory "/tmp/build".
```

## 3. Generate Go package from Smart Contract Abi

1. make `gen` directory

```bash
$ mkdir gen
```

2. run `abigen` development tool to generate go package from abi

```bash
$ abigen --bin=build/Todo.bin --abi=build/Todo.abi -pkg=todo --out=gen/todo.go
```

## 4. Deploy Smart Contract

```bash
$ go run ./06-deploy-contract/
------------------------------
Contract Address: 0xF76a652F4a980B27933edeee1063Ec78610bD452
Tx Hash: 0xe8ed88abb198557ea8e781c44b177bc4a88af520e18c1c9d8b1f7879c071c0e3
------------------------------
```

## 5. Interact with Smart Contract

```bash
$ go run ./07-interact-with-contract/
```