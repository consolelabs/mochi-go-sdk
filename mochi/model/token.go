package model

import "time"

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
