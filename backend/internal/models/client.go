package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name          string
	Email         string
	Phone         string
	BillingAddress string
	AccountManager string
}
