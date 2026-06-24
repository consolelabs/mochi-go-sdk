# Mochi SDK for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/consolelabs/mochi-go-sdk/mochi.svg)](https://pkg.go.dev/github.com/consolelabs/mochi-go-sdk/mochi)

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

## Documentation

NOTICE: This library and the Mochi API are unfinished. Because of that there may be major changes to library in the future.

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


<!-- consolidation-hardening: dev-docs -->
## Development & docs

This repo was reindexed in the Console Labs org-consolidation hardening pass (2026-06).

- `CLAUDE.md` , guidance for AI agents + humans (stack, conventions, commands).
- `docs/ARCHITECTURE.md` , what's here and how it fits together.
- `docs/SECURITY-AUDIT-2026-06-25.md` , secret-scan + dependency baseline.
- CI: `.github/workflows/security.yml` runs gitleaks + a dependency audit on every PR.

Build / test:

```
go build ./...
go test ./...
```

Secrets come from env / the platform, never hardcoded.
