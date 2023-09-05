package easyv1

import (
	"crypto/sha256"
)

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		PrevHash: prevBlockHash,
		Data:     []byte(data),
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	info := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", nil)
}
