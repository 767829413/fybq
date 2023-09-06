package main

import (
	"fmt"

	easyv3 "github.com/767829413/fybq/demo/easy-v3"
	// easyv1 "github.com/767829413/fybq/demo/easy-v1"
)

func main() {
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// db.Update(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("b1"))
	// 	if bucket == nil {
	// 		bucket, err = tx.CreateBucket([]byte("b1"))
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// 	bucket.Put([]byte("okokoko"), []byte("hhhhhhhhhhh"))
	// 	bucket.Put([]byte("ccccccc"), []byte("fffffffffff"))
	// 	return nil
	// })
	
	// db.View(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("b1"))
	// 	if bucket == nil {
	// 		log.Fatal("bucket is nil")
	// 	}
	// 	fmt.Printf("%s\n",bucket.Get([]byte("okokoko")))
	// 	fmt.Printf("%s\n",bucket.Get([]byte("ccccccc")))
	// 	return nil
	// })
	BlockChain := easyv3.NewBlockChain()
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
