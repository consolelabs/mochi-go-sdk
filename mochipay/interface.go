package mochipay

type APIClient interface {
	GetApplicationBalances() ([]TokenBalance, error)
	RequestPayment(req *PaymentRequest) error
	Transfer(req *TransferRequest) ([]TransactionResult, error)
}
