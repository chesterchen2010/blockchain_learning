package main

import (
	"log"

	"github.com/boltdb/bolt"
)

//区块链遍历器
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

/*遍历器*/

//接收者为遍历器的方法，返回下一个区块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bu.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}
