package main

import (
	"fmt"
	_ "yapart/morestrings"
	"yapart/services/postgress"

	//import "github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	postgress.Setup()
}

func main() {
	// var b models.IBrand = &models.Brand{}
	// b.Add()
	// var m models.IMark = &models.Mark{}
	// m.Add()

	// var brands []models.Brand
	// db.Select("Id, name, picture_id").Find(&brands)
	// for _, v := range brands {
	// 	fmt.Println(v)
	// }

	//main

	fmt.Println("asdasd")
}
