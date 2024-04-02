package model

import "time"

type TransferTransaction struct {
	Timestamp      int64  `json:"timestamp"`
	TransactionID  int64  `json:"tx_id"`
	RecipientID    string `json:"recipient_id"`
	Amount         string `json:"amount"`
	TransactionFee string `json:"tx_fee"`
	Status         string `json:"status"`
	References     string `json:"references"`
}

type TransactionType string

const (
	TransactionTypeSend    TransactionType = "out"
	TransactionTypeReceive TransactionType = "in"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSubmitted TransactionStatus = "submitted"
	TransactionStatusSuccess   TransactionStatus = "success"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusExpired   TransactionStatus = "expired"
	TransactionStatusCancelled TransactionStatus = "cancelled" //nolint:all
)

type TransactionAction string

const (
	TransactionActionTransfer      TransactionAction = "transfer"
	TransactionActionAirdrop       TransactionAction = "airdrop"
	TransactionActionDeposit       TransactionAction = "deposit"
	TransactionActionWithdraw      TransactionAction = "withdraw"
	TransactionActionSwap          TransactionAction = "swap"
	TransactionActionVaultTransfer TransactionAction = "vault_transfer"
)

type TransactionPlatform string

const (
	TransactionPlatformDiscord TransactionPlatform = "discord"
)

type GetTransactionsRequest struct {
	Type         TransactionType       `json:"type"`
	Status       TransactionStatus     `json:"status"`
	ActionList   []TransactionAction   `json:"action_list"`
	TokenAddress string                `json:"token_address"`
	ProfileID    string                `json:"profile_id"`
	Platforms    []TransactionPlatform `json:"platforms"`
	ChainIDs     []string              `json:"chain_ids"`
	Page         int64                 `json:"page"`
	Size         int64                 `json:"size"`
	IsSender     *bool                 `json:"is_sender"`
	SortBy       string                `json:"sort_by"`
}

type GetTransactionsResponse struct {
	Data       []TransactionData `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

type TransactionData struct {
	Id                 string                 `json:"id"`
	FromProfileId      string                 `json:"from_profile_id"`
	OtherProfileId     string                 `json:"other_profile_id"`
	FromProfileSource  string                 `json:"from_profile_source"`
	OtherProfileSource string                 `json:"other_profile_source"`
	SourcePlatform     string                 `json:"source_platform"`
	Amount             string                 `json:"amount"`
	TokenId            string                 `json:"token_id"`
	ChainId            string                 `json:"chain_id"`
	InternalId         int64                  `json:"internal_id"`
	ExternalId         string                 `json:"external_id"`
	OnchainTxHash      string                 `json:"onchain_tx_hash"`
	Type               TransactionType        `json:"type"`
	Action             TransactionAction      `json:"action"`
	Status             TransactionStatus      `json:"status"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	ExpiredAt          *time.Time             `json:"expired_at"`
	SettledAt          *time.Time             `json:"settled_at"`
	Token              *Token                 `json:"token"`
	OriginalTxId       string                 `json:"original_tx_id"`
	OtherProfile       *Profile               `json:"other_profile"`
	FromProfile        *Profile               `json:"from_profile"`
	OtherProfiles      []Profile              `json:"other_profiles"`
	AmountEachProfiles []AmountEachProfiles   `json:"amount_each_profiles"`
	UsdAmount          float64                `json:"usd_amount"`
	Metadata           map[string]interface{} `json:"metadata"`

	// used in airdrop response
	OtherProfileIds []string `json:"other_profile_ids"`
	TotalAmount     string   `json:"total_amount"`

	// used in swap response
	FromTokenId string `json:"from_token_id"`
	ToTokenId   string `json:"to_token_id"`
	FromToken   *Token `json:"from_token,omitempty"`
	ToToken     *Token `json:"to_token,omitempty"`
	FromAmount  string `json:"from_amount"`
	ToAmount    string `json:"to_amount"`
}

type AmountEachProfiles struct {
	ProfileId string  `json:"profile_id"`
	Amount    string  `json:"amount"`
	UsdAmount float64 `json:"usd_amount"`
}
