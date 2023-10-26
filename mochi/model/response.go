package model

type TokenBalanceResponses struct {
	Data []TokenBalance `json:"data"`
}

type TransactionResponse struct {
	Data []Transaction `json:"data"`
}

type TokensResponse struct {
	Data []Token `json:"data"`
}

type ChainsResponse struct {
	Data []Chain `json:"data"`
}
