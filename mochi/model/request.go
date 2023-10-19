package model

type PaymentRequest struct {
	UserProfileID string `json:"user_profile_id"`
	TokenAmount   string `json:"token_amount"`
	TokenID       string `json:"token_id"`
	Description   string `json:"description"`
	CallbackURL   string `json:"callback_url"`
}

type TransferRequest struct {
	RecipientIDs []string `json:"recipient_ids"`
	Amounts      []string `json:"amounts"`
	TokenID      string   `json:"token_id"`
	References   string   `json:"references"`
	Description  string   `json:"description"`
}
