package model

// TODO: add settings fixtures
type SystemSetting struct {
	BasicModel
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// available keys
var (
	SETTING_PAYMENT_TIMEOUT = "SETTING_PAYMENT_TIMEOUT"
)

func (SystemSetting) TableName() string {
	return "SystemSetting"
}
