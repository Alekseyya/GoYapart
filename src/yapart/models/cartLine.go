package models

type CartLine struct {
	Id          int     `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	CartId      int     `gorm:"column:CartId"`
	Article     string  `gorm:"column:Article"`
	Description string  `gorm:"column:Description"`
	Price       float32 `gorm:"column:Price"`
	Quantity    int     `gorm:"column:Quantity"`
}
