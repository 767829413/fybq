package main

import (
	easyv3 "github.com/767829413/fybq/demo/easy-v3"
	// easyv1 "github.com/767829413/fybq/demo/easy-v1"
)

func main() {
	easyv3.Execute()
	// BlockChain := easyv3.NewBlockChain()
	// BlockChain.AddBlock("小智向皮卡丘转账100SB")
	// BlockChain.AddBlock("小智向小火龙转账200SB")
	// BlockChain.AddBlock("小刚向小智转账1000SB")
	// bci := BlockChain.NewIterator()
	// var block *easyv3.Block
	// for {
	// 	block = bci.Next()
	// 	if block == nil {
	// 		break
	// 	}
	// 	fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
	// 	fmt.Printf("当区块哈希值: %x\n", block.Hash)
	// 	fmt.Printf("数据: %s\n", block.Data)
	// }
	// BlockChain.CloseDB()
}
