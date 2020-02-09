package models

type Group struct {
	Id   int    `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name string `gorm:"column:Name"`
}
