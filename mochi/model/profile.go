package model

import "time"

type Profile struct {
	ID                 string              `json:"id"`
	CreatedAt          *time.Time          `json:"created_at"`
	UpdatedAt          *time.Time          `json:"updated_at"`
	AssociatedAccounts []AssociatedAccount `json:"associated_accounts"`
	ProfileName        string              `json:"profile_name"`
	Avatar             string              `json:"avatar"`
	Pnl                string              `json:"pnl"`
}

type AssociatedAccount struct {
	ID                 string     `json:"id"`
	ProfileID          string     `json:"profile_id"`
	Platform           string     `json:"platform"`
	PlatformIdentifier string     `json:"platform_identifier"`
	PlatformMetadata   struct{}   `json:"platform_metadata"`
	DiscordID          string     `json:"discord_id"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	TotalAmount        string     `json:"total_amount"`
	Pnl                string     `json:"pnl"`
}
