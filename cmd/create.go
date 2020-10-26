package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

var (
	numAccounts uint

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new accounts",
		Run: func(cmd *cobra.Command, args []string) {
			createNewAccounts()
		},
	}
)

func init() {
	createCmd.PersistentFlags().UintVar(&numAccounts, "num", 1, "Number of accounts to create")

	rootCmd.AddCommand(createCmd)
}

func inputAndConfirmPassword() (string, error) {
	fmt.Println("Create accounts, please input password for key file!")

	passwd1 := mustInputPassword("Enter password: ")
	passwd2 := mustInputPassword("Confirm password: ")

	if passwd1 != passwd2 {
		return "", errors.New("password mismatch")
	}

	return passwd1, nil
}

func createNewAccounts() {
	password, err := inputAndConfirmPassword()
	if err != nil {
		fmt.Println("Failed to create account:", err.Error())
		return
	}

	for i := uint(0); i < numAccounts; i++ {
		account, err := am.Create(password)
		if err != nil {
			fmt.Println("Failed to create account:", err.Error())
			return
		}

		fmt.Println(account.String())
	}

	fmt.Printf("Totally %v new accounts created.\n", numAccounts)
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
