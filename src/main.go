package main

import (
	"fmt"
	"morestrings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=yapart_store password=jfgSvSD@gM sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
