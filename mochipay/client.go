package mochipay

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) APIClient {
	if cfg.BaseURL == "" {
		if cfg.IsPreview {
			cfg.BaseURL = DefaultPreviewBaseURL
		} else {
			cfg.BaseURL = DefaultProdBaseURL
		}
	}

	return &Client{
		cfg: cfg,
	}
}
