package models

type Client struct {
	ClientID       int    `json:"clientID" gorm:"primaryKey;column:ClientID"`
	Name           string `json:"name" gorm:"index"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	BillingAddress string `json:"billing_address"`
	AccountManager string `json:"account_manager"`
	// DeletedAt     gorm.DeletedAt `json:"-"` // Soft Delete field
}
