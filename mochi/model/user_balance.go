package model

import "time"

type UserTokenBalanceResponses struct {
	Data []UserTokenBalance `json:"data"`
}

type UserTokenBalance struct {
	ID        string    `json:"id"`
	ProfileID string    `json:"profile_id"`
	TokenID   string    `json:"token_id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
	UsdAmount float64   `json:"usd_amount"`
}
