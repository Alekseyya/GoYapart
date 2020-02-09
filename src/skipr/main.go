package main

import (
	"fmt"
	"skipr/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/RPCS?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// var tsHoursRecord []models.TSHoursRecord
	// db.Where(&models.TSHoursRecord{ID: 2536}).First(&tsHoursRecord)
	// for _, v := range tsHoursRecord {
	// 	fmt.Println(v)
	// }
	var tsHoursRecord = models.TSHoursRecord{ProjectID: 230, EmployeeID: 21,
		RecordDate: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC), Hours: 2.5,
		Description: "Проверка Golang", Created: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC),
		RecordStatus: 1, RecordSource: 1,
		IsVersion:     false,
		VersionNumber: 0,
		IsDeleted:     false, Modified: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC)}

	// var errors = db.Create(&tsHoursRecord).GetErrors()
	// fmt.Println(errors)

	for index := 0; index < 20; index++ {
		if errors := db.Create(&tsHoursRecord).GetErrors(); errors != nil {
			for _, e := range errors {
				fmt.Println(e)
			}
		}
		fmt.Println(index)
	}
}
