package postgress

import (
	"log"
	"yapart/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=yapart_store password=jfgSvSD@gM sslmode=disable")
	if err != nil {
		log.Fatal("models.Setup err: %v", err)
	}
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func Migrate() {
	db.AutoMigrate(&models.Brand{}, &models.Cart{}, &models.CartLine{}, &models.Category{}, &models.Group{}, &models.Mark{},
		&models.Model{}, &models.Modification{}, &models.Order{}, &models.OrderItem{}, &models.Picture{},
		&models.Product{}, &models.ProductModification{}, &models.Section{})
}
