# Blockchain-Go

Coding for fun, GO GO GO! :beers: :beers:

![Golang](https://golang.org/lib/godoc/images/footer-gopher.jpg?style=centerme)

### Introduction
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

Hope this is clear enough to understand how it works.

### Implementation

#### Hash
- Simply using SHA-256

#### Proof of Work Algorithm
- Hashcash

### Pre-Request
- Golang 13.6

### Build
```go
$ GOROOT=/usr/local/go
$ GOPATH=$pwd
$ go build -i -o Blockchain_Go.exe .
or
$ go build Blockchain_Go/
```
### Run
```go
$ ./Blockchain_Go.exe
or
$ Blockchain_Go
```

### Acknowledge
All credits go to [Jeiwan](https://jeiwan.net/posts/building-blockchain-in-go-part-1/)