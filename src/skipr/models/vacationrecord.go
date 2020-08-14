package models

type VacationRecord struct {
	ID        int
	ShortName string
}

func (VacationRecord) TableName() string {
	return "vacationrecord"
}
