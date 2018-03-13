package model

import "time"

const (
	// statuses
	TR_STATUS_STARTED  = "started"
	TR_STATUS_ERROR    = "error"
	TR_STATUS_APPROVED = "approved"
	TR_STATUS_DONE     = "done"

	// types
	TR_TYPE_WALLET  = "wallet"
	TR_TYPE_CARD    = "card"
	TR_TYPE_PAYMENT = "payment"
	TR_TYPE_ACCOUNT = "account"
)

type Transaction struct {
	BasicModel
	Amount    float64   `json:"amount" gorm:"column:amount"`
	Type      string    `json:"type" gorm:"column:type"`
	To        string    `json:"to" gorm:"column:to"`
	StartTime time.Time `json:"startTime" gorm:"column:startTime"`
	EndTime   time.Time `json:"endTime" gorm:"column:endTime"`
	Status    string    `json:"status" gorm:"column:status"`
	Wallets   []Wallet  `gorm:"many2many:WalletTransaction"`
}

func (Transaction) TableName() string {
	return "Transaction"
}

func IsTransactionTypeTransferable(trType string) bool {
	switch true {
	case trType == TR_TYPE_CARD || trType == TR_TYPE_WALLET:
		return true
	default:
		return false
	}
}
