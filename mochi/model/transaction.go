package model

type Transaction struct {
	Timestamp      int64  `json:"timestamp"`
	TransactionID  int64  `json:"tx_id"`
	RecipientID    string `json:"recipient_id"`
	Amount         string `json:"amount"`
	TransactionFee string `json:"tx_fee"`
	Status         string `json:"status"`
	References     string `json:"references"`
}
