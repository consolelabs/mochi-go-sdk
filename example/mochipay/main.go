package main

import (
	"fmt"

	"github.com/consolelabs/mochi-go-sdk/mochi"
	"github.com/consolelabs/mochi-go-sdk/mochi/config"
)

func main() {
	config := &config.Config{
		MochiPay: config.MochiPay{
			ApplicationID:   "<application-id>",
			ApplicationName: "<application-name>",
			APIKey:          "<api-key>",
		},
	}

	client := mochi.NewClient(config)
	balances, err := client.GetAppBalance()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Balances:")
	for _, balance := range balances {
		fmt.Printf("Token ID: %s, Amount: %s\n", balance.TokenID, balance.Amount)
	}
}
