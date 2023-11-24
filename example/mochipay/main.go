package main

import (
	"fmt"
	"os"

	"github.com/consolelabs/mochi-go-sdk/mochi"
	"github.com/consolelabs/mochi-go-sdk/mochi/config"
)

func main() {
	config := &config.Config{
		MochiPay: config.MochiPay{
			ApplicationID:   os.Getenv("MOCHI_PAY_APPLICATION_ID"),
			ApplicationName: os.Getenv("MOCHI_PAY_APPLICATION_NAME"),
			APIKey:          os.Getenv("MOCHI_PAY_API_KEY"),
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
