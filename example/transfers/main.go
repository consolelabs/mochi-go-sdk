package main

import (
	"fmt"
	"os"

	"github.com/consolelabs/mochi-go-sdk/mochi"
	"github.com/consolelabs/mochi-go-sdk/mochi/config"
	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func main() {
	cfg := &config.Config{
		MochiPay: config.MochiPay{
			ApplicationID:   os.Getenv("MOCHI_PAY_APPLICATION_ID"),
			ApplicationName: os.Getenv("MOCHI_PAY_APPLICATION_NAME"),
			APIKey:          os.Getenv("MOCHI_PAY_API_KEY"),
		},
	}

	client := mochi.NewClient(cfg)
	txs, err := client.Transfer(&model.TransferRequest{
		RecipientIDs: nil,
		Amounts:      nil,
		TokenID:      "",
		References:   "",
		Description:  "",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Transfer Transactions:")
	for _, tx := range txs {
		fmt.Printf("Transaction ID: %s, Amount: %s, Status: %s\n", tx.TransactionID, tx.Amount, tx.Status)
	}
}
