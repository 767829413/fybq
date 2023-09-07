package easyv3

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/767829413/fybq/util"
)

type ProofOfWorkload struct {
	block *Block
	// 目标值
	target *big.Int
}

func NewProofOfWorkload(block *Block) *ProofOfWorkload {
	pow := &ProofOfWorkload{block: block}
	// 指定难度值,目前手写
	tmpTargetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	tmpBigInt := &big.Int{}
	// 指定16进制
	tmpBigInt.SetString(tmpTargetStr, 16)
	pow.target = tmpBigInt
	return pow
}

func (pow *ProofOfWorkload) Run() ([]byte, uint64) {
	var (
		nonce uint64
		hash  []byte
	)
	block := pow.block
	for {
		// 1. 数据拼装
		info := bytes.Join([][]byte{
			util.Uint64ToBytes(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			util.Uint64ToBytes(block.TimeStamp),
			util.Uint64ToBytes(block.Difficulty),
			util.Uint64ToBytes(nonce),
			block.Data,
		}, nil)
		// 2. 哈希运算
		h := sha256.Sum256(info)
		hash = h[:]
		tmpBigInt := big.Int{}
		tmpBigInt.SetBytes(hash)
		// 3. 比较pow中的target进行比较

		// 3-b. 没找到,继续找,随机数+1
		if tmpBigInt.Cmp(pow.target) == -1 {
			// 3-a. 找到了,退出返回
			fmt.Printf("Mining success! hash: %x, nonce: %d\n", hash, nonce)
			return hash, nonce
		} else {
			nonce++
		}
	}
}
