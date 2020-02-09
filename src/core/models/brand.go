package model

import "github.com/jinzhu/gorm"

type Brand struct{
	Id int `gorm:"primary_key" json:"id"`
	name string `json:"name"`
	picture_id int `gorm:"FOREIGNKEY"` 
}

