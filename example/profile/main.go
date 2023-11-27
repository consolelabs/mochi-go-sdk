package main

import (
	"fmt"

	"github.com/consolelabs/mochi-go-sdk/mochi"
	"github.com/consolelabs/mochi-go-sdk/mochi/config"
)

func main() {
	cfg := &config.Config{}

	client := mochi.NewClient(cfg)
	result, err := client.GetByDiscordID("discord_id")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Profile:", result)
}
