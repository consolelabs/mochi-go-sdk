package main

import (
	"fmt"

	"github.com/consolelabs/mochi-go-sdk/mochi"
	"github.com/consolelabs/mochi-go-sdk/mochi/config"
)

func main() {
	config := &config.Config{}

	client := mochi.NewClient(config)
	result, err := client.GetByDiscordID("797042642600722473")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Profile:", result)
}
