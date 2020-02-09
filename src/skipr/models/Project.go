package models

type Project struct {
	ID            int             `gorm:"primary_key, AUTO_INCREMENT;column:ID"`
	TSHoursRecord []TSHoursRecord `gorm:foreignkey:ProjectID`
	ShortName     string          `gorm:column:"ShortName"`
}
