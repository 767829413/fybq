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
	Short: "FROM makes a transaction to TO",
	Run: func(cmd *cobra.Command, args []string) {
		getBalance(NewBlockChain(addr))
	},
}

var from string
var to string
var amount float64
var miner string
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "FROM makes a single transaction to TO while MINER mines and writes to DATA",
	Run: func(cmd *cobra.Command, args []string) {
		send(NewBlockChain(addr))
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&data, "data", "d", "YOUR_CONTENT_DATA", "addBlock data (required)")
	addCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR_ADDRESS", "addBlock address (required)")
	addCmd.MarkPersistentFlagRequired("data")
	addCmd.MarkPersistentFlagRequired("address")
	addCmd.Flag("data")
	addCmd.Flag("address")

	printCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR_ADDRESS", "addBlock address (required)")
	printCmd.MarkPersistentFlagRequired("address")
	printCmd.Flag("address")

	getBalanceCmd.PersistentFlags().StringVarP(&addr, "address", "r", "YOUR_ADDRESS", "addBlock address (required)")
	getBalanceCmd.MarkPersistentFlagRequired("address")
	getBalanceCmd.Flag("address")

	sendCmd.PersistentFlags().StringVarP(&from, "from", "f", "FROM_ADDRESS", "send from (required)")
	sendCmd.PersistentFlags().StringVarP(&to, "to", "t", "TO_ADDRESS", "send to (required)")
	sendCmd.PersistentFlags().Float64VarP(&amount, "amount", "a", 0.0, "send amount (required)")
	sendCmd.PersistentFlags().StringVarP(&miner, "miner", "m", "MINER_ADDRESS", "send miner (required)")
	sendCmd.PersistentFlags().StringVarP(&data, "data", "d", "YOUR_CONTENT_DATA", "send data (required)")
	sendCmd.MarkPersistentFlagRequired("from")
	sendCmd.MarkPersistentFlagRequired("to")
	sendCmd.MarkPersistentFlagRequired("amount")
	sendCmd.MarkPersistentFlagRequired("miner")
	sendCmd.MarkPersistentFlagRequired("data")
	sendCmd.Flag("from")
	sendCmd.Flag("to")
	sendCmd.Flag("amount")
	sendCmd.Flag("miner")
	sendCmd.Flag("data")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(getBalanceCmd)
	rootCmd.AddCommand(sendCmd)
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
	fmt.Printf("The balance of %s is: %f\n", addr, balance)
}

func send(blockChain *BlockChain) {
	fmt.Println("Commencement of transfers")
	// 1. 创建挖矿交易
	coinBase := NewCoinbase(miner, data)
	// 2. 创建普通交易
	txs := NewTransaction(from, to, amount, blockChain)
	if txs == nil {
		fmt.Println("send failed,closing of the transaction")
		return
	}
	// 3. 添加区块
	blockChain.AddBlock([]*Transaction{coinBase, txs})
	fmt.Println("Closing of the transaction")
}
