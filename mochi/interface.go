package mochi

import "github.com/consolelabs/mochi-go-sdk/mochi/model"

type APIClient interface {
	GetByDiscordID(id string) (*model.Profile, error)

	GetAppBalance() ([]model.TokenBalance, error)
	RequestPayment(req *model.PaymentRequest) error
	Transfer(req *model.TransferRequest) ([]model.Transaction, error)

	GetTokens(req *model.GetTokenRequest) ([]model.Token, error)
	GetChains() ([]model.Chain, error)
}
