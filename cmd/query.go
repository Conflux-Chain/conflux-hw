package cmd

import (
	"fmt"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/spf13/cobra"
)

var (
	account string

	queryCmd = &cobra.Command{
		Use:   "query",
		Short: "Query on-chain information for specific account",
		Run: func(cmd *cobra.Command, args []string) {
			query()
		},
	}
)

func init() {
	queryCmd.PersistentFlags().StringVar(&url, "url", "http://main.confluxrpc.org", "Conflux RPC url")
	queryCmd.PersistentFlags().StringVar(&account, "account", "", "Account address in HEX format")
	queryCmd.MarkPersistentFlagRequired("account")

	rootCmd.AddCommand(queryCmd)
}

func query() {
	client, err := sdk.NewClient(url)
	if err != nil {
		fmt.Println("Failed to create client:", err.Error())
		return
	}

	epoch, err := client.GetEpochNumber()
	if err != nil {
		fmt.Println("Failed to get epoch number:", err.Error())
		return
	}
	fmt.Println("Epoch number:", epoch)

	address := *types.NewAddress(account)

	balance, err := client.GetBalance(address)
	if err != nil {
		fmt.Println("Failed to get balance:", err.Error())
		return
	}
	fmt.Println("Account balance:", balance)

	nonce, err := client.GetNextNonce(address, nil)
	if err != nil {
		fmt.Println("Failed to get nonce:", err.Error())
		return
	}
	fmt.Println("Account nonce:", nonce)
}
