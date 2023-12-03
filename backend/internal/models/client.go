package models

import "gorm.io/gorm"

type Client struct {
    gorm.Model
    Name          string `json:"name"`
    Email         string `json:"email"`
    Phone         string `json:"phone"`
    BillingAddress string `json:"billing_address"`
    AccountManager string `json:"account_manager"`
}
