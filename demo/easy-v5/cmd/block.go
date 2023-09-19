package cmd

import (
	"fmt"

	easyv5 "github.com/767829413/fybq/demo/easy-v5"
	"github.com/spf13/cobra"
)

func blockInit() {
	printCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR_ADDRESS", "addBlock address (required)")
	printCmd.MarkPersistentFlagRequired("address")
	printCmd.Flag("address")
}

var printCmd = &cobra.Command{
	Use:   "printChain",
	Short: "Print block chain data",
	Run: func(cmd *cobra.Command, args []string) {
		printChain(easyv5.NewBlockChain(addr))
	},
}

func printChain(blockChain *easyv5.BlockChain) {
	fmt.Println("Start printing block data")
	bci := blockChain.NewIterator()
	var block *easyv5.Block
	for {
		block = bci.Next()
		if block == nil {
			break
		}
		fmt.Printf("Version number: %x\n", block.Version)
		fmt.Printf("Pre-block hash: %x\n", block.PrevHash)
		fmt.Printf("Merkle root: %x\n", block.MerkelRoot)
		fmt.Printf("Timestamp: %x\n", block.TimeStamp)
		fmt.Printf("Difficulty value: %x\n", block.Difficulty)
		fmt.Printf("Nonce: %x\n", block.Nonce)
		fmt.Printf("Current block hash: %x\n", block.Hash)
		if len(block.Transactions) != 0 {
			fmt.Printf("Data: %s\n", block.Transactions[0].TXInputs[0].Sig)
		}
	}
	blockChain.CloseDB()
}
