package model

import (
	"time"

	"github.com/Barrokgl/go-qiwi-wallet-api"
)

//go:generate go-functiozzo -type Wallet -package model
type Wallet struct {
	BasicModel
	Phone        string    `json:"phone" gorm:"column:phone"`
	InUse        bool      `json:"inUse" gorm:"column:inUse"`
	Status       string    `json:"status" gorm:"column:status"`
	QiwiToken    string    `json:"qiwiToken" gorm:"column:qiwiToken"`
	TokenExpires time.Time `json:"tokenExpires" gorm:"column:tokenExpires"`

	ClientService   ClientService `json:"clientService,omitempty"`
	ClientServiceID int           `json:"clientServiceId" gorm:"column:clientServiceId"`

	Transactions []Transaction `json:"transactions,omitempty" gorm:"many2many:WalletTransaction"`

	Balances       []goqiwi.Sum `json:"balances" gorm:"-" sql:"-"`
	Blocked        bool         `json:"blocked" gorm:"-" sql:"-"`
	Identification string       `json:"identification" gorm:"-" sql:"-"`
}

func (Wallet) TableName() string {
	return "Wallet"
}
