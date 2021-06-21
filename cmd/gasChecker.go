/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//TODO: Fix this command. Its calculating the gwei wrong somehow and not notifying.
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-toast/toast"
	"github.com/hrharder/go-gas"
	"github.com/spf13/cobra"
)

var userPrice int64

// gasCheckerCmd represents the gasChecker command
var gasCheckerCmd = &cobra.Command{
	Use:   "gasChecker",
	Short: "Notifies the user for average Ethereum gas prices under a certain price.",
	Long: `The ethNotify gasChecker command notifies the user when average gas price 
	below a certain threshold. Use the --price flag to set the threshold price. eg; 
	
	"ethNotify gasChecker --price=100" //checks for average gas prices under 100 gwei.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Checking for average gas prices under %v\n", userPrice)
		notification := &toast.Notification{
			AppID: "GasChecker",
			Title: "Average gas price",
			Icon:  "C:/dev/go/gaschecker/icon.png",
		}

		prices := make(chan string)
		//Receive the price from Ethgasstation api, get only the first 3 digits.
		go func() {
			for {
				averagePrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
				if err != nil {
					log.Fatal(err)
				}
				price := averagePrice.String()[:3]
				if price[2] == '0' && price[0] != 1 {
					price = price[:2]
				}

				prices <- price
				time.Sleep(4 * time.Minute)

			}
		}()
		//Compare the current gas price to the users flag and notify when necessary.
		for price := range prices {
			intPrice, err := strconv.Atoi(price)
			if err != nil {
				log.Fatal(err)
			}
			notification.Message = fmt.Sprintf("The current average gas price is %v Gwei.", intPrice)
			fmt.Println(notification.Message)
			if UserPrice >= intPrice {
				err := notification.Push()
				if err != nil {
					log.Fatal(err)
				}
			}

		}
		fmt.Println("gasChecker called")
	},
}

func init() {
	rootCmd.AddCommand(gasCheckerCmd)
	// gasCheckerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	gasCheckerCmd.Flags().Int64VarP(&userPrice, "price", "p", 0, "The price to check for (in gwei)")
}
