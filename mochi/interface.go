package mochi

import "github.com/consolelabs/mochi-go-sdk/mochi/model"

type APIClient interface {
	GetAppBalance() ([]model.TokenBalance, error)
	RequestPayment(req *model.PaymentRequest) error
	Transfer(req *model.TransferRequest) ([]model.Transaction, error)

	GetByDiscordID(id string) (*model.Profile, error)
}
