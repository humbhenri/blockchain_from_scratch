# Blockchain from Scratch

WARNING: WORK IN PROGRESS !!!

## TODO
  * [x] Block model
  * [x] Calculate block hash
  * [x] Block validation
  * [x] Mine blocks
  * [x] Save blockchain to filesystem
  * [x] Create block explorer visualization
  * [ ] Sync blockchain data between two nodes
  * [ ] Implement consensus algorithm between nodes
  * [ ] Create wallet to store keys
  * [ ] Create transaction and signatures

## How to run
```
go run main.go
```

## Testing and interacting with the blockchain

Add some data in the chain
``` shell
./udp.sh ADD_DATA 21312
```


A test script that exercises interactions between nodes is provided in `test.sh`.

## References
https://builtin.com/blockchain/create-your-own-blockchain
https://github.com/defunkydrummer/legochain.git
https://medium.com/@bradford_hamilton/write-your-own-blockchain-and-pow-algorithm-using-crystal-d53d5d9d0c52
https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc
https://medium.com/programmers-blockchain/create-simple-blockchain-java-tutorial-from-scratch-6eeed3cb03fa
