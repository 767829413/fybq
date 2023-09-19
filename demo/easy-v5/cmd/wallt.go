package cmd

import (
	"fmt"

	easyv5 "github.com/767829413/fybq/demo/easy-v5"
	"github.com/spf13/cobra"
)

var newWalletCmd = &cobra.Command{
	Use:   "newWallet",
	Short: "Create a wallet",
	Run: func(cmd *cobra.Command, args []string) {
		newWallet()
	},
}

var listAddrCmd = &cobra.Command{
	Use:   "listAddr",
	Short: "List Addresses",
	Run: func(cmd *cobra.Command, args []string) {
		listAddress()
	},
}

func newWallet() {
	fmt.Println("Start creating a wallet")
	wallets := easyv5.GetWallets()
	wallets.CreateWallet()
	fmt.Println("Created Successfully")
}

func listAddress() {
	fmt.Println("Start listing addresses")
	wallets := easyv5.GetWallets()
	addrs := wallets.GetAllAddresses()
	for _, addr := range addrs {
		fmt.Printf("address: %s\n", addr)
	}
}
