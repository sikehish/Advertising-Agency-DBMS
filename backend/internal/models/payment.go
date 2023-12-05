package models

import (
	"time"
)

type Payment struct {
    PaymentID            int        `gorm:"primaryKey" json:"paymentID"`
    InvoiceID            int        `json:"invoiceID"`
    EmployeeID           int        `json:"employeeID"`
    PaymentDate          time.Time  `gorm:"type:date" json:"paymentDate"`
    PaymentMethod        string     `json:"paymentMethod"`
    TransactionReference string     `json:"transactionReference"`
    // CreatedAt            time.Time  `json:"createdAt"`
    // UpdatedAt            time.Time  `json:"updatedAt"`
    // DeletedAt            gorm.DeletedAt  `gorm:"index" json:"-"`
}