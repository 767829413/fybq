package easyv4

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
var addr string
var addCmd = &cobra.Command{
	Use:   "addBlock",
	Short: "Add a block,--data or -d DATA",
	Run: func(cmd *cobra.Command, args []string) {
		addBlock(NewBlockChain(addr))
	},
}

var printCmd = &cobra.Command{
	Use:   "printChain",
	Short: "Print block chain data",
	Run: func(cmd *cobra.Command, args []string) {
		printChain(NewBlockChain(addr))
	},
}

var getBalanceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "Get the balance of the specified address",
	Run: func(cmd *cobra.Command, args []string) {
		getBalance(NewBlockChain(addr))
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&data, "data", "d", "YOUR CONTENT DATA", "addBlock data (required)")
	addCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR ADDRESS", "addBlock address (required)")
	addCmd.MarkPersistentFlagRequired("data")
	addCmd.MarkPersistentFlagRequired("address")
	addCmd.Flag("data")
	addCmd.Flag("address")
	getBalanceCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR ADDRESS", "addBlock address (required)")
	getBalanceCmd.MarkPersistentFlagRequired("address")
	getBalanceCmd.Flag("address")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(getBalanceCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addBlock(blockChain *BlockChain) {
	fmt.Println("Start adding block data")
	txs := []*Transaction{}
	blockChain.AddBlock(txs)
	fmt.Println("Add successfully")
}

func printChain(blockChain *BlockChain) {
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
		if len(block.Transactions) != 0 {
			fmt.Printf("Data: %s\n", block.Transactions[0].TXInputs[0].Sig)
		}
	}
	blockChain.CloseDB()
}

func getBalance(blockChain *BlockChain) {
	fmt.Println("Start printing the balance of the specified address")
	utxos := blockChain.FindUTXOs(addr)
	var balance = 0.0
	for _, utxo := range utxos {
		balance += utxo.Value
	}
	fmt.Printf("The balance of %x is: %f\n", addr, balance)
}
