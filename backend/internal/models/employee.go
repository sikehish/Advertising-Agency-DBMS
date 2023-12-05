package models

type Employee struct {
	EmployeeID int     `gorm:"primaryKey" json:"employeeID"`
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Salary     float64 `json:"salary"`
	Department string  `json:"department"`
	// CreatedAt  time.Time  `json:"createdAt"`
	// UpdatedAt  time.Time  `json:"updatedAt"`
	// DeletedAt  gorm.DeletedAt  `gorm:"index" json:"-"`
}
