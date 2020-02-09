package models

type ProductModification struct {
	Id             int `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	ProductId      int `gorm:"column:ProductId"`
	ModificationId int `gorm:"column:ModificationId"`
}
