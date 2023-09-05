package main

import (
	"fmt"

	easyv1 "github.com/767829413/fybq/demo/easy-v1"
)

func main() {
	BlockChain := easyv1.NewBlockChain()
	BlockChain.AddBlock("小智向皮卡丘转账100SB")
	BlockChain.AddBlock("小智向小火龙转账200SB")
	BlockChain.AddBlock("小刚向小智转账1000SB")
	for i, block := range BlockChain.Blocks {
		fmt.Printf("当前区块高度: %d\n", i)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("当区块哈希值: %x\n", block.Hash)
		fmt.Printf("数据: %s\n", block.Data)
	}
}
