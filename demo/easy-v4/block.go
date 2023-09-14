package easyv4

import (
	"bytes"
	"crypto/sha256"
	"time"

	"github.com/767829413/fybq/util"
)

type Block struct {
	// 版本号
	Version uint64
	// 前区块链哈希
	PrevHash []byte
	// 墨克根
	MerkelRoot []byte
	// 时间戳
	TimeStamp uint64
	// 难度值
	Difficulty uint64
	// 随机数
	Nonce uint64
	// 当前区块哈希
	Hash []byte
	// 交易数据
	Transactions []*Transaction
}

func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Version:      00,
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
		PrevHash:     prevBlockHash,
		Transactions: txs,
	}
	block.MerkelRoot = block.MakeMerkelRoot()
	pow := NewProofOfWorkload(block)
	// 查找目标随机数不断进行哈希运算
	hash, nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	// block.SetHash()
	return block
}

func (block *Block) SetHash() {
	// 只需要对区块头进行哈希
	info := bytes.Join([][]byte{
		util.Uint64ToBytes(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		util.Uint64ToBytes(block.TimeStamp),
		util.Uint64ToBytes(block.Difficulty),
		util.Uint64ToBytes(block.Nonce),
	}, nil)
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

// 创建Merkel Root
// 直接拼接非二叉树实现
func (block *Block) MakeMerkelRoot() []byte {
	return nil
}

func NewGenesisBlock(addr string) *Block {
	coinbase := NewCoinbase(addr, "Genesis Block")
	return NewBlock([]*Transaction{coinbase}, nil)
}
