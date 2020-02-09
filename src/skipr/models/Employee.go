package models

type Employee struct {
	ID            int             `gorm:"primary_key, AUTO_INCREMENT;column:ID"`
	LastName      string          `gorm:column:"LastName"`
	FirstName     string          `gorm:column:FirstName`
	TSHoursRecord []TSHoursRecord `gorm:foreignkey:EmployeeID`
}
