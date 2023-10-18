package main

import (
	"fmt"

	"github.com/consolelabs/mochi-go-sdk/mochipay"
)

func main() {
	config := &mochipay.Config{
		ApplicationID:   "<application-id>",
		ApplicationName: "<application-name>",
		APIKey:          "<api-key>",
	}

	client := mochipay.NewClient(config)
	balances, err := client.GetApplicationBalances()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Balances:")
	for _, balance := range balances {
		fmt.Printf("Token ID: %s, Amount: %s\n", balance.TokenID, balance.Amount)
	}
}
