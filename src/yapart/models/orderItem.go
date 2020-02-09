package models

import (
	"github.com/jinzhu/gorm"
)

type OrderEnum int

//https://support.bigcommerce.com/s/article/Order-Statuses
const (
	Pending                   OrderEnum = iota
	AwaitingPayment           OrderEnum = iota
	Awaitingfulfillment       OrderEnum = iota
	AwaitingShipment          OrderEnum = iota
	AwaitingPickup            OrderEnum = iota
	AwaitingShipped           OrderEnum = iota
	Completed                 OrderEnum = iota
	Shipped                   OrderEnum = iota
	Cancelled                 OrderEnum = iota
	Declined                  OrderEnum = iota
	Refunded                  OrderEnum = iota
	Disputed                  OrderEnum = iota
	ManualVerificationRequire OrderEnum = iota
	PartiallyRefunded         OrderEnum = iota
)

type OrderItem struct {
	gorm.Model
	Id                 int       `gorm:"primary_key,AUTO_INCREMENT;column:Id;not null;"`
	OrderId            int       `gorm:"column:OrderId"`
	Article            string    `gorm:"column:Article"`
	Description        string    `gorm:"column:Description"`
	PriceWithDiscount  float32   `gorm:"column:PriceWithDiscount"`
	Quantity           int       `gorm:"column:Quantity"`
	Comment            string    `gorm:"column:Comment"`
	IsClosed           bool      `gorm:"column:IsClosed"`
	OrderStatus        OrderEnum `gorm:"column:OrderStatus;type:integer"`
	OrderStatusComment string    `gorm:"column:OrderStatusComment"`
}
