package models

import (
	"errors"
	"fmt"
	. "yapart/services/postgress"
)

type IBrand interface {
	IRepository
}

//Брэнд
type Brand struct {
	Id                   int                   `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;" json:"id"`
	Name                 string                `gorm:"column:Name" json:"name"`
	PictureId            int                   `gorm:"column:PictureId"`
	Products             []Product             `gorm:"foreignkey:BrandId"`
	CategoryId           int                   `gorm:"column:CategoryId"`
	Pictures             []Picture             `gorm:"foreignkey:BrandId"`
	ProductModifications []ProductModification `gorm:"foreignkey:BrandId"`
}

func (b *Brand) Get() ([]*Brand, error){
	var brands []*Brand
	if err := YapartDB.Find(&brands).Error; err != nil {
		return brands, err
	}
	return nil, errors.New("Таблица пуста")
}

func (b *Brand) Add() {
	fmt.Println("Brand")
}

func (b *Brand) Delete() {
	fmt.Println("AAAAAAAAAAA")
}

func (b *Brand) Update() {

}

func Transform() {
	fmt.Println("Transform")
}

func Transform2() {
	fmt.Println("Transform2")
}
