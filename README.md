# Blockchain-Go

Coding for fun, GO GO GO! :beers: :beers:

![Bitcoin](https://miro.medium.com/max/3000/1*bFlCApzWW8EYVmSAnXcWYA.jpeg)

### I. Introduction
Building a simplified Blockchain in Go.

If you are not familiar with Blockchain concepts, go [here](https://youtu.be/TVlo66aOZE0) (A very good video in Chinese).

Generally, it is a distributed database of transaction records and is public to everyone. Everyone owns a copy
(sounds like replica log problem) and 
the system needs to make sure everyone owns the same copy(Consistency).

Each block contains several transactions and "pointers" to the previous block and the next block. 
This pointer is derived from hashing algorithm based on the information on the each block. 
Any change on the block will lead to the change of this hash value(pointers that points to another blocks), 
which indicates the breach of chain.

Therefore, no one can modify the transactions on each block privately. 
And the complex calculation of hash value is motivation for people to mine Bitcoin.

Just imagine, if we live in a world without trusty third-party transaction platform, such as Alipay, Paypal, 
how can we know it is safe to conduct the transactions? 

Blockchain provides the possibility that "broadcast" every transactions to all entities with E-signature(Private key).
Every participant have access to check whether this transaction is valid and this process guarantee the 
reliability of this system.

Above is my interpretation, hope this is clear enough to understand how it works.

### II. Implementation
####Function
So far, this blockchain supports mining, transaction and wallet management, I would to enable it to be a 
distributed system in the near future.

#### Language Version
- Golang 13.6
- Golang Modules: [BoltDB](https://github.com/boltdb/bolt.git)

#### File Structure
- main.go
- block.go
- blockchain.go
- cli.go
- proofofwork.go
- transaction.go
- utils.go
- wallet.go
- wallets.go
- base85.go

#### Hash
- Simply using SHA-256

#### Proof of Work Algorithm
- Hashcash

The principle is a brute force method, the more powerful of your calculation capability,
the higher possibilities you get the answer(hash value).

#### Database
- BoltDB

#### Digital Signature
- Elliptic Curve Cryptography (Without deep implementation, just use it)

#### Readable Public Key
- Base58 Algorithm

### III. Build
```go
$ GOROOT=/usr/local/go
$ GOPATH=$pwd
$ go build -i -o Blockchain_Go.exe .
or
$ go build Blockchain_Go/
```
This repository can be run in GoLand directly.

### IV. Run
```go
$ Blockchain_Go createwallet
$ Blockchain_Go createblockchain -address <wallet address>
$ Blockchain_Go send -from <wallet address1> -to <wallet address2>
$ Blockchain_Go printchain
$ Blockchain_Go addblock -data "Send 1 BTC to Jade"
$ Blockchain_Go getbalance -address <wallet address>
```

### V. Result Sample
```text
$ blockchain_go createwallet
Your new address: 1NzmVp1G9qHzzk5xqQE9pR3njNfXQK5Psy

$ blockchain_go createwallet
Your new address: 1N3NM8heUZnBZoGvCRtQNJBbamRc7B5AXi

$ blockchain_go createblockchain -address 1NzmVp1G9qHzzk5xqQE9pR3njNfXQK5Psy
0000003bdc62e4028da8ccc36704805fc10497b6644248301300bf588fe207ad
Done!

$ blockchain_go printchain
============ Block 0000003bdc62e4028da8ccc36704805fc10497b6644248301300bf588fe207ad ============
Prev. block:
PoW: false

$ blockchain_go getbalance -address 1NzmVp1G9qHzzk5xqQE9pR3njNfXQK5Psy
Balance of '1NzmVp1G9qHzzk5xqQE9pR3njNfXQK5Psy': 0

```
```text

```

### VI. Acknowledge
All credits go to [Jeiwan](https://jeiwan.net/posts/building-blockchain-in-go-part-1/)

![Golang](https://golang.org/lib/godoc/images/footer-gopher.jpg?style=centerme)