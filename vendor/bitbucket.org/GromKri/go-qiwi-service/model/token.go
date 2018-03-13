package model

type Token struct {
	BasicModel
	Token         string `json:"token" gorm:"column:token"`
	Type          string `json:"type" gorm:"column:type"`
	OwnerID       int64  `json:"ownerId" gorm:"column:ownerId"`
	JTI           string `json:"jti" gorm:"column:jti"`
	SignSecret    string `json:"signSecret" gorm:"column:signSecret"`
	RequestSecret string `json:"requestSecret" gorm:"column:requestSecret"`
	Blacklisted   bool   `json:"blacklisted" gorm:"column:blacklisted"`
}

var (
	TOKEN_TYPE_ADMIN   = "admin"
	TOKEN_TYPE_SERVICE = "SERVICE"
)

func (Token) TableName() string {
	return "Token"
}
