package cmd

import (
	"fmt"
	"os"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

var (
	password string
	am       *sdk.AccountManager = sdk.NewAccountManager("keystore")

	rootCmd = &cobra.Command{
		Use:   "conflux-hw",
		Short: "Conflux hardware wallet tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Account:", mustGetOrCreateAccount())
		},
	}
)

func mustGetOrCreateAccount() types.Address {
	if account, err := am.GetDefault(); err == nil {
		return *account
	}

	fmt.Println("Create an account, please input password for key file!")

	passwd1 := mustInputPassword("Enter password: ")
	passwd2 := mustInputPassword("Confirm password: ")

	if passwd1 != passwd2 {
		fmt.Println("Password mismatch!")
		os.Exit(1)
	}

	password = passwd1

	account, err := am.Create(password)
	if err != nil {
		fmt.Println("Failed to create account:", err.Error())
		os.Exit(1)
	}

	return account
}

func mustInputPassword(prompt string) string {
	fmt.Print(prompt)

	passwd, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println("Failed to read password:", err.Error())
		os.Exit(1)
	}

	return string(passwd)
}

// Execute is the command line entrypoint.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
