package models

type Section struct {
	Id         int        `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name       string     `gorm:"column:Name"`
	Show       bool       `gorm:"column:Show"`
	Sort       int        `gorm:"column:Sort"`
	GroupId    int        `gorm:"column:GroupId"`
	Categories []Category `gorm:"foreignkey:CategoryId"`
}
