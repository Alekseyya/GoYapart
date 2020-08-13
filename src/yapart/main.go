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
	fmt.Printf("assdfsd")
}
