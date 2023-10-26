# Mochi SDK for Go

The Mochi SDK for Go is a library that provides easy access to the Mochi API from your Go applications. It simplifies the process of making API requests, handling authentication, and processing responses. Use this SDK to integrate Mochi's functionality into your Go projects effortlessly.

## Features

### Mochi Profile
- [x] GetByDiscordID: Retrieve a Mochi profile by Discord ID.

### Mochi Pay
- [x] GetAppBalance: Retrieve token balances for your Mochi application.
- [x] RequestPayment: Request a payment from a user.
- [x] Transfer: Transfer tokens from your application to a list of user.
- [x] GetChains: Get the list of supported chain
- [x] GetTokens: Get the list of token filtered by Chain and Symbol.

## Installation

To use the Mochi SDK in your Go project, you can simply install it using:

```bash
go get github.com/consolelabs/mochi-go-sdk
```


## Authorization
From MochiPay, you will receive an application ID, application name, and API key. You can use these to create a new MochiPay client. The client will be used to make requests to the Mochi API.
```go
    config := &mochipay.Config{
	ApplicationID:   "<application-id>",
	ApplicationName: "<application-name>",
	APIKey:          "<api-key>",
    }
```

## Examples
Here's a simple example of how to use the YourAPI SDK:
### Mochi Pay
```go
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


```

### Mochi Profile
```go
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


```