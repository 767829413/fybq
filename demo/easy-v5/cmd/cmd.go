package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	data   string
	addr   string
	from   string
	to     string
	amount float64
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

func init() {
	blockInit()
	transactionInit()

	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(getBalanceCmd)
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(newWalletCmd)
	rootCmd.AddCommand(listAddrCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
