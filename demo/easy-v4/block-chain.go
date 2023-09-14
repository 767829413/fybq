package easyv4

import (
	"fmt"
	"log"

	"github.com/767829413/fybq/util"
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

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	block := NewBlock(txs, bc.lastHash)
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucketName))
		if bucket == nil {
			log.Fatal("AddBlock " + blockBucketName + " is not found")
		}
		bucket.Put(block.Hash, util.StructToBytes(block))
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

func (bc *BlockChain) FindTransactionUTXOs(addr string, amount float64) (map[string][]int64, float64) {
	// 找到的合理的UTXO集合
	var UTXO = make(map[string][]int64)
	// 包含的总额
	var calc float64
	f := func(a ...any) bool {
		tx := a[0].(*Transaction)
		out := a[1].(*TXOutput)
		i := a[2].(int)
		spendUTXOs := a[3].(map[string][]int64)
		var flag = true

		ff := func() bool {
			// 统计总额
			calc += out.Value
			UTXO[string(tx.TXID)] = append(UTXO[string(tx.TXID)], int64(i))
			// 判断总额是否满足
			return calc >= amount
		}
		// 通过addr匹配区分
		if out.PubKeyHash == addr {
			if indexArr, ok := spendUTXOs[string(tx.TXID)]; !ok {
				if ff() {
					return true
				}
			} else {
				for _, idx := range indexArr {
					if idx == int64(i) {
						flag = false
					}
				}
				if flag {
					if ff() {
						return true
					}
				}
			}
		}
		return false
	}
	bc.IterativeTransactions(f)
	return UTXO, calc
}

// 找到指定地址的所有UTXO
func (bc *BlockChain) FindUTXOs(addr string) []*TXOutput {
	var UTXO []*TXOutput
	f := func(a ...any) bool {
		tx := a[0].(*Transaction)
		out := a[1].(*TXOutput)
		i := a[2].(int)
		spendUTXOs := a[3].(map[string][]int64)
		var flag = true
		// 通过addr匹配区分
		if out.PubKeyHash == addr {
			if indexArr, ok := spendUTXOs[string(tx.TXID)]; !ok {
				UTXO = append(UTXO, out)
			} else {
				for _, idx := range indexArr {
					if idx == int64(i) {
						flag = false
					}
				}
				if flag {
					UTXO = append(UTXO, out)
				}
			}
		}
		return false
	}
	bc.IterativeTransactions(f)
	return UTXO
}

func (bc *BlockChain) IterativeTransactions(f func(a ...any) bool) {
	var spendUTXOs = make(map[string][]int64)
	bci := bc.NewIterator()
	// 遍历区块
	for {
		block := bci.Next()
		if block == nil {
			fmt.Println("Block traversal complete!")
			break
		}
		if len(block.Transactions) == 0 {
			continue
		}
		// 遍历交易
		for _, tx := range block.Transactions {
			// 遍历Output,找到相关的utxo(在添加前需要检查是否消耗)
			for i, out := range tx.TXOutputs {
				flag := f(tx, out, i, spendUTXOs)
				if flag {
					return
				}
			}
			if !tx.IsCoinbase() {
				// 遍历Input,找到花费的utxo
				for _, in := range tx.TXInputs {
					// 寻找对应地址的input
					if in.Sig == addr {
						idx := string(in.QTXID)
						indexArr := spendUTXOs[idx]
						indexArr = append(indexArr, in.Index)
						spendUTXOs[idx] = indexArr
					}
				}
			}
		}
	}
}

func NewBlockChain(addr string) *BlockChain {
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
			block := NewGenesisBlock(addr)
			lastHash = block.Hash
			err = bucket.Put(block.Hash, util.StructToBytes(block))
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
