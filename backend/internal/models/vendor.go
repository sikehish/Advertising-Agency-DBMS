package models

import (
	"time"
)

type VendorSupplierInformation struct {
    VendorID        int        `gorm:"primaryKey" json:"vendorID"`
    VendorName      string     `json:"vendorName"`
    ContactInfo     string     `json:"contactInfo"`
    ServicesProvided string    `json:"servicesProvided"`
    PaymentTerms    string     `json:"paymentTerms"`
    // CreatedAt       time.Time  `json:"createdAt"`
    // UpdatedAt       time.Time  `json:"updatedAt"`
    // DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}

type VendorInvoice struct {
    InvoiceID       int        `gorm:"primaryKey" json:"invoiceID"`
    PaymentID       int        `json:"paymentID"`
    VendorID        int        `json:"vendorID"`
    InvoiceDate     time.Time  `gorm:"type:date" json:"invoiceDate"`
    TotalAmount     float64    `json:"totalAmount"`
    PaymentStatus   string     `json:"paymentStatus"`
    // CreatedAt       time.Time  `json:"createdAt"`
    // UpdatedAt       time.Time  `json:"updatedAt"`
    // DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}
