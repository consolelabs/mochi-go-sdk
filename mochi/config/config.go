package config

const (
	DefaultMochiPayPreviewBaseURL = "https://api-preview.mochi-pay.console.so"
	DefaultMochiPayProdBaseURL    = "https://api.mochi-pay.console.so"
	DefaultMochiProfileBaseURL    = "https://api.mochi-profile.console.so"
)

type Config struct {
	MochiPay     MochiPay
	MochiProfile MochiProfile
}

type MochiPay struct {
	BaseURL         string
	ApplicationID   string
	ApplicationName string
	APIKey          string
	IsPreview       bool
}

type MochiProfile struct {
	BaseURL string
}
