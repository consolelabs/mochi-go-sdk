package model

type RequestPayment struct {
	UserProfileID string `json:"user_profile_id"`
	TokenAmount   string `json:"token_amount"`
	TokenID       string `json:"token_id"`
	Description   string `json:"description"`
	CallbackURL   string `json:"callback_url"`
}
