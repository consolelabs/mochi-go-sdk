package model

type TokenBalanceResponses struct {
	Data []TokenBalance `json:"data"`
}

type TransactionResponse struct {
	Data []Transaction `json:"data"`
}
