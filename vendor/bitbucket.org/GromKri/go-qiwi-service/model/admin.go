package model

type Admin struct {
	BasicModel
	Login    string `json:"login" required:"true" gorm:"not_null;unique"`
	Password string `json:"password,omitempty" required:"true" gorm:"not_null;unique"`
}

func (Admin) TableName() string {
	return "Admin"
}
