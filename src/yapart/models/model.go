package models

type Model struct {
	Id            int            `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name          string         `gorm:"column:Name"`
	MarkId        int            `gorm:"column:MarkId"`
	Pictures      []Picture      `gorm:"foreignkey:ModelId"`
	Modifications []Modification `gorm:"foreignkey:ModelId"`
}
