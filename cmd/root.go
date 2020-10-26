package cmd

import (
	"fmt"
	"os"
	"sort"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/spf13/cobra"
)

var (
	am *sdk.AccountManager = sdk.NewAccountManager("keystore")

	rootCmd = &cobra.Command{
		Use:   "conflux-hw",
		Short: "Conflux hardware wallet tool",
		Run: func(cmd *cobra.Command, args []string) {
			accounts := listAccountsAsc()
			if len(accounts) == 0 {
				fmt.Println("No account found!")
				return
			}

			for i, addr := range accounts {
				fmt.Printf("[%v]\t%v\n", i, addr)
			}

			fmt.Printf("Totally %v accounts found.\n", len(accounts))
		},
	}
)

func listAccountsAsc() []string {
	var accounts []string

	for _, addr := range am.List() {
		accounts = append(accounts, addr.String())
	}

	sort.Strings(accounts)

	return accounts
}

// Execute is the command line entrypoint.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
