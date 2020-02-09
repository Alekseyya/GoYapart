package models

type Modification struct {
	Id                   int                   `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	ModelId              int                   `gorm:"column:ModelId"`
	Year                 int                   `gorm:"column:Year"`
	Sort                 int                   `gorm:"column:Sort"`
	Pictures             []Picture             `gorm:"foreignkey:ModificationId"`
	ProductModifications []ProductModification `gorm:"foreignkey:ProductModificationId"`
}
