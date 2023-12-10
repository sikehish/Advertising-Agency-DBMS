package models

type ContactUs struct {
	Email   string `json:"email" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"index"`
	Message string `json:"message"`
}
