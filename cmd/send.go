package cmd

import (
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/spf13/cobra"
)

var (
	url string
	raw string

	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send signed transaction",
		Run: func(cmd *cobra.Command, args []string) {
			send()
		},
	}
)

func init() {
	sendCmd.PersistentFlags().StringVar(&url, "url", "http://main.confluxrpc.org", "Conflux RPC url")
	sendCmd.PersistentFlags().StringVar(&raw, "raw", "", "Raw transaction in HEX format")
	sendCmd.MarkPersistentFlagRequired("raw")

	rootCmd.AddCommand(sendCmd)
}

func send() {
	client, err := sdk.NewClient(url)
	if err != nil {
		fmt.Println("Failed to create client:", err.Error())
		return
	}

	if strings.HasPrefix(raw, "0x") {
		raw = raw[2:]
	}

	rawData, err := hex.DecodeString(raw)
	if err != nil {
		fmt.Println("Failed to decode raw data in HEX format:", err.Error())
		return
	}

	txHash, err := client.SendRawTransaction(rawData)
	if err != nil {
		fmt.Println("Failed to send raw transaction:", err.Error())
		return
	}

	fmt.Println(txHash)
}
