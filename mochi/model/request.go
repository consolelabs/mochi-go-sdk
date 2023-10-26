package model

// PaymentRequest is used to request payment
type PaymentRequest struct {
	UserProfileID string `json:"user_profile_id"`
	TokenAmount   string `json:"token_amount"`
	TokenID       string `json:"token_id"`
	Description   string `json:"description"`
	CallbackURL   string `json:"callback_url"`
}

// TransferRequest is used to transfer tokens
type TransferRequest struct {
	RecipientIDs []string `json:"recipient_ids"`
	Amounts      []string `json:"amounts"`
	TokenID      string   `json:"token_id"`
	References   string   `json:"references"`
	Description  string   `json:"description"`
}

// GetTokenRequest is used to get tokens by chain id and symbol
type GetTokenRequest struct {
	ChainID string `json:"chain_id"`
	Symbol  string `json:"symbol"`
}
