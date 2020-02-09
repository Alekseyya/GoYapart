package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type ShippingEnum int
type PaymentEnum int
type CityEnum int

const (
	//Самолет
	RussianAir ShippingEnum = iota
	//Почта
	PostOffice ShippingEnum = iota
)

const (
	DebitCard PaymentEnum = iota
	Cash      PaymentEnum = iota
)

const (
	Moscow          CityEnum = iota
	SaintPetersburg CityEnum = iota
)

type Order struct {
	gorm.Model
	Id             int          `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	ShippingMethod ShippingEnum `gorm:"column:ShippingMethod;type:integer"`
	PaymentMethod  PaymentEnum  `gorm:"column:PaymentMethod;type:integer"`
	City           CityEnum     `gorm:"column:City;type:integer"`
	Phone          string       `gorm:"column:Phone"`
	ShippedDate    sql.NullTime `gorm:"column:ShippedDate"`
	Comment        string       `gorm:"column:Comment"`
	IsSend         bool         `gorm:"column:IsSend"`
	IsClosed       bool         `gorm:"column:IsClosed"`
	OrderItems     []OrderItem  `gorm:"foreignkey:OrderId"`
}
