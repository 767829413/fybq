package easyv3

import (
	"log"

	"github.com/767829413/fybq/util"
	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	db *bolt.DB
	// 游标,用于不断索引
	curHashPointer []byte
}

func (bci *BlockChainIterator) Next() *Block {
	var block Block
	bci.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucketName))
		if bucket == nil {
			log.Fatal("BlockChainIterator.Next " + blockBucketName + " is not found")
		}
		tmpBlockBytes := bucket.Get(bci.curHashPointer)
		if tmpBlockBytes != nil {
			util.BytesToStruct(tmpBlockBytes, &block)
		}
		return nil
	})
	if block.Hash == nil {
		bci.curHashPointer = nil
		return nil
	}
	bci.curHashPointer = block.PrevHash
	return &block
}
