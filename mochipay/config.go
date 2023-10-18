package mochipay

const (
	DefaultPreviewBaseURL = "https://api-preview.mochi-pay.console.so"
	DefaultProdBaseURL    = "https://api.mochi-pay.console.so"
)

type Config struct {
	BaseURL         string
	ApplicationID   string
	ApplicationName string
	APIKey          string
	IsPreview       bool
}
