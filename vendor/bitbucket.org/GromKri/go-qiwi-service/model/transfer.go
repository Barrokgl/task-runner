package model

type Transfer struct {
	BasicModel
	Phone  string  `json:"phone" gorm:"column:phone"`
	Amount float64 `json:"amount" gorm:"column:amount"`
	Status string  `json:"status" gorm:"column:status"`
	Error  string  `json:"error" gorm:"column:error"`

	Transaction   Transaction `json:"transaction,omitempty"`
	TransactionID int64       `json:"transactionId" gorm:"column:transactionId"`

	Wallet   Wallet `json:"wallet,omitempty"`
	WalletID int    `json:"walletId" gorm:"column:walletId"`
}

func (Transfer) TableName() string {
	return "Transfer"
}
