package mochiprofile

import "fmt"

type PaymentRequest struct {
	UserProfileID string `json:"user_profile_id"`
	TokenAmount   string `json:"token_amount"`
	TokenID       string `json:"token_id"`
	Description   string `json:"description"`
	CallbackURL   string `json:"callback_url"`
}

func (p PaymentRequest) ToMessageHeader() string {
	return fmt.Sprintf("Transfer(%v,%v,%v)", p.UserProfileID, p.TokenID, p.TokenAmount)
}
