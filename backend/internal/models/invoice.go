package models

import (
	"time"
)

type Invoice struct {
    InvoiceID      int        `gorm:"primaryKey" json:"invoiceID"`
    ClientID       int        `json:"clientID"`
    CampaignID     int        `json:"campaignID"`
    InvoiceDate    time.Time  `gorm:"type:date" json:"invoiceDate"`
    PaymentStatus  string     `json:"paymentStatus"`
    TotalAmount    float64    `json:"totalAmount"`
    // CreatedAt      time.Time  `json:"createdAt"`
    // UpdatedAt      time.Time  `json:"updatedAt"`
    // DeletedAt      gorm.DeletedAt  `gorm:"index" json:"-"`
}