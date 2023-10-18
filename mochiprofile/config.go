package mochiprofile

const (
	DefaultPreviewBaseURL = "https://api.mochi-profile.console.so"
	DefaultProdBaseURL    = "https://api.mochi-profile.console.so"
)

type Config struct {
	BaseURL   string
	IsPreview bool
}
