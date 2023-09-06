package easyv3

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
	Blocks   []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	block := NewBlock(data, bc.lastHash)
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucketName))
		if bucket == nil {
			log.Fatal(blockBucketName + "is not found")
		}
		bucket.Put(block.Hash, block.ToBytes())
		bucket.Put([]byte(lastHashKey), block.Hash)
		bc.lastHash = block.Hash
		return nil
	})
	bc.Blocks = append(bc.Blocks, block)
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
				log.Fatal(err)
			}
			// 创世块
			block := NewGenesisBlock()
			lastHash = block.Hash
			err = bucket.Put(block.Hash, block.ToBytes())
			if err != nil {
				log.Fatal("bucket.Put(block.Hash, block.ToBytes())", err)
			}
			err = bucket.Put([]byte(lastHashKey), lastHash)
			if err != nil {
				log.Fatal("bucket.Put([]byte(lastHashKey), lastHash)", err)
			}
		} else {
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})

	return &BlockChain{
		db:       db,
		lastHash: lastHash,
		Blocks:   []*Block{},
	}
}
