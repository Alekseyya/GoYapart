package models

import "fmt"

type IMark interface {
	IRepository
}

type Mark struct {
	Id       int       `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name     string    `gorm:"column:Name" json:"name"`
	Show     bool      `gorm:"column:Show" json:"show"`
	Pictures []Picture `gorm:"foreignkey:MarkId"`
	Models   []Model   `gorm:"foreignkey:ModelId"`
}

func (m *Mark) Add() {
	fmt.Println("Marks")
}
func (m *Mark) Delete() {

}

func (m *Mark) Update() {

}

func (m *Mark) Get() {

}
