package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// matchAddressCmd represents the matchAddress command
var matchAddressCmd = &cobra.Command{
	Use:   "matchAddress",
	Short: "Checks two Ethereum addresses to see if they match.",
	Long: `matchAddress will check to Ethereum addresses to see if they match, 
	for use when sending between wallets/exchanges.

	ex: You withdraw tokens from Binance. Binance sends a confirmation email with the withdrawal
	address. Paste the address from the email as the first argument to matchAddress, then copy/paste
	your withdrawal address from your wallet. MatchAddress will confirm you are withdrawing to the correct address.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		valid := common.IsHexAddress(args[0]) && common.IsHexAddress(args[1])
		fmt.Println("matchAddress called")
		if !valid {
			fmt.Println("On or both addresses are invalid!")
		} else {
			if !match(args[0], args[1]) {
				fmt.Println("Addresses do not match")
			} else {
				fmt.Println("Addresses match.")
			}
		}
	},
}

func match(a1, a2 string) bool {
	return a1 == a2
}

func init() {
	rootCmd.AddCommand(matchAddressCmd)
}
