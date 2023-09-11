package easyv4

import (
	"log"

	"github.com/boltdb/bolt"
)

const (
	blockChainDbName = "blockChain.db"
	blockBucketName  = "blockBucket"
	lastHashKey      = "lastHashKey"
)

type BlockChain struct {
	db       *bolt.DB
	lastHash []byte
}

func (bc *BlockChain) AddBlock(data string) {
	block := NewBlock(data, bc.lastHash)
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucketName))
		if bucket == nil {
			log.Fatal("AddBlock " + blockBucketName + " is not found")
		}
		bucket.Put(block.Hash, block.ToBytes())
		bucket.Put([]byte(lastHashKey), block.Hash)
		bc.lastHash = block.Hash
		return nil
	})
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db: bc.db,
		// 最初指向末尾
		curHashPointer: bc.lastHash,
	}
}

func (bc *BlockChain) CloseDB() error {
	return bc.db.Close()
}

func NewBlockChain() *BlockChain {
	db, err := bolt.Open(blockChainDbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucketName))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucketName))
			if err != nil {
				log.Fatal("NewBlockChain CreateBucket " + blockBucketName + " is faild: " + err.Error())
			}
			// 创世块
			block := NewGenesisBlock()
			lastHash = block.Hash
			err = bucket.Put(block.Hash, block.ToBytes())
			if err != nil {
				log.Fatal("NewBlockChain bucket.Put(block.Hash, block.ToBytes())", err)
			}
			err = bucket.Put([]byte(lastHashKey), lastHash)
			if err != nil {
				log.Fatal("NewBlockChain bucket.Put([]byte(lastHashKey), lastHash)", err)
			}
		} else {
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})

	return &BlockChain{
		db:       db,
		lastHash: lastHash,
	}
}
