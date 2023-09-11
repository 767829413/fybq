package easyv4

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var blockChain *BlockChain

var rootCmd = &cobra.Command{
	Use:  "mbc",
	Long: `mbc is a very easy block chain`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Please create your block!!
		Get the action with --help or -h.
		`)
	},
}

var data string
var addCmd = &cobra.Command{
	Use:   "addBlock",
	Short: "Add a block,--data or -d DATA",
	Run: func(cmd *cobra.Command, args []string) {
		addBlock()
	},
}

var printCmd = &cobra.Command{
	Use:   "printChain",
	Short: "print block chain data",
	Run: func(cmd *cobra.Command, args []string) {
		printChain()
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&data, "data", "d", "YOUR CONTENT DATA", "addBlock data (required)")
	addCmd.MarkPersistentFlagRequired("data")
	addCmd.Flag("data")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(printCmd)
	blockChain = NewBlockChain()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addBlock() {
	fmt.Println("Start adding block data")
	blockChain.AddBlock(data)
	fmt.Println("Add successfully")
}

func printChain() {
	fmt.Println("Start printing block data")
	bci := blockChain.NewIterator()
	var block *Block
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
		fmt.Printf("Data: %s\n", block.Data)
	}
	blockChain.CloseDB()
}
