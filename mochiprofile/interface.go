package mochiprofile

type APIClient interface {
	GetByDiscordID(id string) (*ProfileResponse, error)
}
