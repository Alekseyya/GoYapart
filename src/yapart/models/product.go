package models

import (
	"github.com/jinzhu/gorm"
)

//TODO добавить версионность
//Вопрос версионности будет решаться через gorm.model - записи не будут удаляться навсегда

//Товары
type Product struct {
	gorm.Model
	Id                   int                   `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Article              string                `gorm:"column:Article"`
	ShortArticle         string                `gorm:"column:Description"`
	Price                float32               `gorm:"column:Price"`
	DaysDelivery         int                   `gorm:"column:DaysDelivery"`
	Popular              bool                  `gorm:"column:Popular"`
	Characteristic       string                `gorm:"column:Characteristic"`
	Discount             bool                  `gorm:"column:Discount"`
	BrandId              int                   `gorm:"column:BrandId"`
	CategoryId           int                   `gorm:"column:CategoryId"`
	Pictures             []Picture             `gorm:"foreignkey:ProductId"`
	ProductModifications []ProductModification `gorm:"foreignkey:ProductId"`
}
