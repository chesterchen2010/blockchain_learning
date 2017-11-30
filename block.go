package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//区块结构体
type Block struct {
	Timestamp int64
	// Data []byte
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

//函数，创建一个新Block
// func NewBlock(data string, prevBlockHash []byte) *Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	// block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	// block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.RunPow()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//函数，创建创世区块
// func NewGenesisBlock() *Block {
func NewGenesisBlock(coinbase *Transaction) *Block {
	// return NewBlock("Genesis Block", []byte{})
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID) //把[]byte加入到[][]byte
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

/*正反序列化*/

//序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
