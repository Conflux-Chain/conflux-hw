package cmd

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/spf13/cobra"
)

const (
	defaultGasPrice int64 = 10
	defaultGasLimit int64 = 21000
)

var (
	nonce    uint32
	to       string
	valueStr string
	epoch    uint64
	chain    uint

	signCmd = &cobra.Command{
		Use:   "sign",
		Short: "Sign transaction to send",
		Run: func(cmd *cobra.Command, args []string) {
			sign()
		},
	}
)

func init() {
	signCmd.PersistentFlags().Uint32Var(&nonce, "nonce", 0, "Transaction nonce")
	signCmd.MarkPersistentFlagRequired("nonce")
	signCmd.PersistentFlags().StringVar(&to, "to", "", "To address in HEX format")
	signCmd.MarkPersistentFlagRequired("to")
	signCmd.PersistentFlags().StringVar(&valueStr, "value", "", "Value to transfer in drip")
	signCmd.MarkPersistentFlagRequired("value")
	signCmd.PersistentFlags().Uint64Var(&epoch, "epoch", 0, "Transaction epoch height")
	signCmd.MarkPersistentFlagRequired("epoch")
	signCmd.PersistentFlags().UintVar(&chain, "chain", 2, "Conflux chain ID")
	signCmd.PersistentFlags().StringVar(&password, "password", "", "Password to decrypt key file")

	rootCmd.AddCommand(signCmd)
}

func sign() {
	from := mustGetOrCreateAccount()
	value, ok := new(big.Int).SetString(valueStr, 10)
	if !ok {
		fmt.Println("invalid value")
		return
	}

	tx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:         &from,
			Nonce:        types.NewBigInt(int64(nonce)),
			GasPrice:     types.NewBigInt(defaultGasPrice),
			Gas:          types.NewBigInt(defaultGasLimit),
			Value:        types.NewBigIntByRaw(value),
			StorageLimit: types.NewUint64(0),
			EpochHeight:  types.NewUint64(epoch),
			ChainID:      types.NewUint(chain),
		},
		To: types.NewAddress(to),
	}

	if len(password) == 0 {
		password = mustInputPassword("Enter password: ")
	}

	encoded, err := am.SignAndEcodeTransactionWithPassphrase(tx, password)
	if err != nil {
		fmt.Println("Failed to sign transaction:", err.Error())
		return
	}

	fmt.Println("0x" + hex.EncodeToString(encoded))
}
