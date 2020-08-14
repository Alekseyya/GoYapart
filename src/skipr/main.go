package main

import (
	"fmt"
	"skipr/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "postgres"
	password      = "jfgSvSD@gM"
	database_name = "rpcs_go"
)

func AddTSHoursRecord(dbMySQL *gorm.DB) {
	var tsHoursRecord = models.TSHoursRecord{ProjectID: 230, EmployeeID: 21,
		RecordDate: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC), Hours: 2.5,
		Description: "Проверка Golang", Created: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC),
		RecordStatus: 1, RecordSource: 1,
		IsVersion:     false,
		VersionNumber: 0,
		IsDeleted:     false, Modified: time.Date(2020, time.January, 23, 0, 0, 0, 0, time.UTC)}

	for index := 0; index < 20; index++ {
		if errors := dbMySQL.Create(&tsHoursRecord).GetErrors(); errors != nil {
			for _, e := range errors {
				fmt.Println(e)
			}
		}
		fmt.Println(index)
	}
}

func GetAllTSHoursRecord(dbMySQL *gorm.DB) interface{} {
	var tsRecords = []models.TSHoursRecord{}
	dbMySQL.Select("*").Find(&tsRecords)
	return tsRecords
}

func GetAllEmployee(dbMySQL *gorm.DB) interface{} {
	var employees = []models.Employee{}
	dbMySQL.Select("*").Find(&employees)
	return employees
}

func FirstMigrateChema(dbPg *gorm.DB) {
	dbPg.AutoMigrate(&models.Employee{}, &models.Project{}, &models.TSAutoHoursRecord{}, &models.TSHoursRecord{}, &models.VacationRecord{})
}

func Migrate(dbMySQL *gorm.DB, dbPg *gorm.DB) {
	for _, dataTSHoursRecord := range GetAllTSHoursRecord(dbMySQL).([]models.TSHoursRecord) {
		dbPg.Create(dataTSHoursRecord)
	}
	for _, dataEmployee := range GetAllEmployee(dbMySQL).([]models.Employee) {
		dbPg.Create(dataEmployee)
	}
}

func AddForeignKeys(entry interface{}, db *gorm.DB) int {
	switch entry.(type) {
	case models.TSHoursRecord:
		for _, err := range db.Model(&models.TSHoursRecord{}).AddForeignKey("EmployeeID", "Employee", "RESTRICT", "RESTRICT").GetErrors() {
			fmt.Println(err)
		}

	case models.Employee:
		//return len(v.ids)
	}
	return -1
}

func main() {
	var dbMySQL, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/RPCS?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	var dbPg, errPg = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=rpcs_go password=jfgSvSD@gM sslmode=disable")
	if errPg != nil {
		panic(errPg)
	}
	//AddForeignKeys(models.TSHoursRecord{}, dbPg)
	FirstMigrateChema(dbPg)
	//Migrate(dbMySQL, dbPg)

	//for _, dataEmployee := range GetAllEmployee(dbMySQL).([]models.Employee){
	//	fmt.Println(dataEmployee)
	//}

	//fmt.Println(dbMySQL.HasTable(&models.TSHoursRecord{}))
	//fmt.Println(dbPg.HasTable(&models.TSHoursRecord{}))
	//fmt.Println(dbMySQL.HasTable(&models.Employee{}))
	//fmt.Println(dbPg.HasTable(&models.Employee{}))

	//for _, err := range dbPg.Model(&models.Employee{}).AddForeignKey("Employee_ID", "tshoursrecord(ID)",
	//"RESTRICT", "RESTRICT").GetErrors(){
	//	fmt.Println(err)
	//}

	//var employees = []models.Employee{}
	//dbMySQL.Select("*").Find(&employees)
	//
	//for _, dbError := range dbMySQL.Select("*").Find(&employees).GetErrors() {
	//	fmt.Println(dbError)
	//}

	defer dbMySQL.Close()
	defer dbPg.Close()
}
