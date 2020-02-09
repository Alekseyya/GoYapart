package models

//Общий тип репозитория
type IRepository interface{
	Add()
	Delete()
	Update()
	Get()
}