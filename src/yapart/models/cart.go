package models

type Cart struct {
	Id        int        `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	UserId    int        `gorm:"column:UserId"`
	CartLines []CartLine `gorm:"foreignkey:CartId"`
}
