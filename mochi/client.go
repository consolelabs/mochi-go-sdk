package mochi

import "github.com/consolelabs/mochi-go-sdk/mochi/config"

type Client struct {
	cfg *config.Config
}

func NewClient(cfg *config.Config) APIClient {
	if cfg.MochiPay.BaseURL == "" {
		if cfg.MochiPay.IsPreview {
			cfg.MochiPay.BaseURL = config.DefaultMochiPayPreviewBaseURL
		} else {
			cfg.MochiPay.BaseURL = config.DefaultMochiPayProdBaseURL
		}
	}

	if cfg.MochiProfile.BaseURL == "" {
		cfg.MochiProfile.BaseURL = config.DefaultMochiProfileBaseURL
	}

	return &Client{
		cfg: cfg,
	}
}
