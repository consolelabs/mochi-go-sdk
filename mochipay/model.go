package mochipay

import "time"

type ErrorMessage struct {
	Msg string `json:"msg"`
}

type TokenBalanceResponses struct {
	Data []TokenBalance `json:"data"`
}

type TokenBalance struct {
	ID        string    `json:"id"`
	ProfileID string    `json:"profile_id"`
	TokenID   string    `json:"token_id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

type Token struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Decimal     int     `json:"decimal"`
	ChainID     string  `json:"chain_id"`
	Native      bool    `json:"native"`
	Address     string  `json:"address"`
	Icon        string  `json:"icon"`
	CoinGeckoID string  `json:"coin_gecko_id"`
	Price       float64 `json:"price"`
	Chain       Chain   `json:"chain"`
}

type Chain struct {
	ID       string `json:"id"`
	ChainID  string `json:"chain_id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rpc      string `json:"rpc"`
	Explorer string `json:"explorer"`
	Icon     string `json:"icon"`
	IsEvm    bool   `json:"is_evm"`
}

type PaymentRequest struct {
	UserProfileID string `json:"user_profile_id"`
	TokenAmount   string `json:"token_amount"`
	TokenID       string `json:"token_id"`
	Description   string `json:"description"`
	CallbackURL   string `json:"callback_url"`
}

type TransactionResult struct {
	Timestamp      int64  `json:"timestamp"`
	TransactionID  int64  `json:"tx_id"`
	RecipientID    string `json:"recipient_id"`
	Amount         string `json:"amount"`
	TransactionFee string `json:"tx_fee"`
	Status         string `json:"status"`
	References     string `json:"references"`
}

type TransactionResponse struct {
	Data []TransactionResult `json:"data"`
}

type TransferRequest struct {
	RecipientIDs []string `json:"recipient_ids"`
	Amounts      []string `json:"amounts"`
	TokenID      string   `json:"token_id"`
	References   string   `json:"references"`
	Description  string   `json:"description"`
}
