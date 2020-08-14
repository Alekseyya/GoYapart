package models

type Project struct {
	ID            int             `gorm:"primary_key;auto_increment:true;column:ID"`
	TSHoursRecord []TSHoursRecord `gorm:"foreignkey:ProjectID"`
	ShortName     string          `gorm:column:"ShortName"`
}

func (Project) TableName() string {
	return "project"
}
