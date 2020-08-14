package models

type TSAutoHoursRecord struct {
	ID int `gorm:"primary_key;auto_increment:true;column:ID"`
	//TSHoursRecord TSHoursRecord //`gorm:"foreignkey:ParentTSAutoHoursRecordID"`
}

func (TSAutoHoursRecord) TableName() string {
	return "tsautohoursrecord"
}
