package model

type ClientService struct {
	BasicModel
	Name     string `json:"name" gorm:"size:255"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
}

func (ClientService) TableName() string {
	return "ClientService"
}
