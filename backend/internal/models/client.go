package models

type Client struct {
	ClientID       int    `json:"clientId" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"index"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	BillingAddress string `json:"billingAddress"`
	AccountManager string `json:"accountManager"`
	// DeletedAt     gorm.DeletedAt `json:"-"` // Soft Delete field
}
