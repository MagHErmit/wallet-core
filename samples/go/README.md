# Sample Go Integration for [Wallet-Core](https://github.com/trustwallet/wallet-core)

## 🔖 Overview

This folder contains a small **Go** sample integration with
[Wallet Core](https://github.com/trustwallet/wallet-core) library (part of [Trust Wallet](https://trustwallet.com)),
using [cgo](https://golang.org/cmd/cgo/).

## Tx Example 
```json lines
POST http://localhost:8080/api/v1/sign_transaction
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "JSONServer.SignTx",
  "params": [
    {
    "gate":"bitcoin",
    "tx":{
      "utxo":[
        {
          "out_point": {
          "hash": "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
          "index": 0,
          "sequence": 4294967295
          },
          "amount":625000000
        }
      ],
      "to_Address":"1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
      "Change_Address":"1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
      "byte_Fee":1,
      "Amount":1000000
    }
  }
  ]
}
```
## Response example
```json lines
http://localhost:8080/api/v1/sign_transaction

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 12 Nov 2022 05:20:50 GMT
Content-Length: 490

{
  "result": "01000000000101fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f0000000000ffffffff0240420f00000000001976a914769bdff96a02f9135a1d19b749db6a78fe07dc9088ac6d7b3125000000001976a9149e089b6889e032d46e3b915a3392edfd616fb1c488ac0247304402207b54b90868bb77d2ae07175419c54953906f0e7ecb36fc37f558a0ab54bbf3b7022037171f62887936a9924949e766a40a7e19aedabeccd3d255ba5efd427143091101210288be7586c41a0498c1f931a0aaf08c15811ee2651a5fe0fa213167dcaba59ae800000000",
  "error": null,
  "id": 1
}
```


## 📜 Documentation

See the official [Trust Wallet developer documentation here](https://developer.trustwallet.com).

See especially Wallet Core
[Integration Guide](https://developer.trustwallet.com/wallet-core/integration-guide),
and [Build Instructions](https://developer.trustwallet.com/wallet-core/building).

## 🛠 Prerequisites

`macOS` or `Docker`

## ⚙️ Building and Running
###  macOS
#### Prerequisites on macOS
* CMake `brew install cmake`
* Boost `brew install boost`
* Xcode
* Xcode command line tools: `xcode-select --install`
* Other tools: `brew install git ninja autoconf automake libtool xcodegen clang-format`
* GoLang: [download](https://go.dev/dl/)
* Protobuf: `brew install protobuf protoc-gen-go`

#### Full Build

1. Clone the wallet-core repo and go inside:
```shell
git clone https://github.com/trustwallet/wallet-core.git

cd wallet-core
```
2. The full build can be triggered with one top-level script:
```shell
./bootstrap.sh
```

### 🐳 Docker
1. Run `docker run -it trustwallet/wallet-core`
The librabry is already built in this image  (Build instructions [here](building.md))  Note: may not be the most recent version.

2. Install go: `apt-get update && apt-get install golang` 
(or download from here [go1.16.12](https://go.dev/dl/go1.16.12.linux-amd64.tar.gz), configure `GOROOT` and append `GOROOT/bin` to `PATH`).

### 🏃🏽‍♂️ **Run** (macOS & Docker)
1. Go to the **samples/go** folder within wallet core repo:

```shell
cd wallet-core/samples/go
```

2. Compile it by `go build -o main`.  Relavant source file is `main.go`.

3. Run `./main` and you will see the output below: 

```shell
==> calling wallet core from go
==> mnemonic is valid:  true
==> Bitcoin...
```
4. *(optional)* You might want to copy and run `main` outside of the docker container, make sure you have `libc++1` and `libc++abi1` installed in your host Ubuntu.

5. *(optional)* If you want to make transaction on other networks you need to compile `src/proto` proto files and to do that, just run the `./compile.sh` . you can also modify it based on your project.