package models

import (
	"github.com/jinzhu/gorm"
)

//Картинки
type Picture struct {
	gorm.Model
	Id             int    `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name           string `gorm:"column:Name"`
	Path           string `gorm:"column:Path"`
	ProductId      int    `gorm:"column:ProductId"`
	MarkId         int    `gorm:"colunm:MarkId"`
	ModelId        int    `gorm:"column:ModelId"`
	ModificationId int    `gorm:"column:ModificationId"`
	BrandId        int    `gorm:"column:BrandId"`
}
