package models

//Категория
type Category struct {
	Id        int       `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	Name      string    `gorm:"column:Name"`
	Show      bool      `gorm:"column:Show"`
	SectionId int       `gorm:"column:SectionId"`
	Products  []Product `gorm:"column:CategoryId"`
}
