package main

import (
	"fmt"
	_ "yapart/morestrings"
	"yapart/services/postgress"

	//import "github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() {
	postgress.Setup()
}

func main() {
	Init()
	fmt.Printf("assdfsd")
}
