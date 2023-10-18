# Mochi SDK for Go

The Mochi SDK for Go is a library that provides easy access to the Mochi API from your Go applications. It simplifies the process of making API requests, handling authentication, and processing responses. Use this SDK to integrate Mochi's functionality into your Go projects effortlessly.

## Features

- [x] GetApplicationBalances: Retrieve token balances for your Mochi application.
- [x] RequestPayment: Request a payment from a user.
- [x] Transfer: Transfer tokens from your application to a list of user.

## Installation

To use the Mochi SDK in your Go project, you can simply install it using:

```bash
go get github.com/consolelabs/mochi-sdk
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
### `GetApplicationBalances()`
```go
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

```

