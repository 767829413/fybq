package cmd

import (
	"fmt"

	easyv5 "github.com/767829413/fybq/demo/easy-v5"
	"github.com/spf13/cobra"
)

func transactionInit() {
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
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "FROM makes a single transaction to TO while MINER mines and writes to DATA",
	Run: func(cmd *cobra.Command, args []string) {
		send(easyv5.NewBlockChain(addr))
	},
}

var getBalanceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "FROM makes a transaction to TO",
	Run: func(cmd *cobra.Command, args []string) {
		getBalance(easyv5.NewBlockChain(addr))
	},
}

func send(blockChain *easyv5.BlockChain) {
	fmt.Println("Commencement of transfers")
	// 1. 创建挖矿交易
	coinBase := easyv5.NewCoinbase(miner, data)
	// 2. 创建普通交易
	txs := easyv5.NewTransaction(from, to, amount, blockChain)
	if txs == nil {
		fmt.Println("send failed,closing of the transaction")
		return
	}
	// 3. 添加区块
	blockChain.AddBlock([]*easyv5.Transaction{coinBase, txs})
	fmt.Println("Closing of the transaction")
}

func getBalance(blockChain *easyv5.BlockChain) {
	fmt.Println("Start printing the balance of the specified address")
	utxos := blockChain.FindUTXOs(addr)
	var balance = 0.0
	for _, utxo := range utxos {
		balance += utxo.Value
	}
	fmt.Printf("The balance of %s is: %f\n", addr, balance)
}
