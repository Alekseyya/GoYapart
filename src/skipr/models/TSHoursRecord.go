package models

//import "database/sql"
import (
	"database/sql"
	"fmt"
	"time"
)

//Трудозатраты
type TSHoursRecord struct {
	// BaseModel                 BaseModel
	ID           int       `gorm:"primary_key,AUTO_INCREMENT;column:ID;not null;"`
	ProjectID    int       `gorm:"column:ProjectID"`
	EmployeeID   int       `gorm:"column:EmployeeID"`
	RecordDate   time.Time `gorm:"column:RecordDate"`
	Hours        float32   `gorm:"column:Hours"`
	Description  string    `gorm:"column:Description" json:"description"`
	RecordStatus int       `gorm:"column:RecordStatus"`
	RecordSource int       `gorm:"column:RecordSource"`
	//ParentTSAutoHoursRecordID int               `gorm:"column:ParentTSAutoHoursRecordID"`
	TSAutoHoursRecord TSAutoHoursRecord `gorm:"foreignkey:ParentTSAutoHoursRecordID"`
	//ParentVacationRecordID  int               `gorm:"column:ParentVacationRecordID"`
	VacationRecord          VacationRecord `gorm:"foreignkey:ParentVacationRecordID"`
	ExternalSourceElementID string         `gorm:"column:ExternalSourceElementID"`
	ExternalSourceListID    string         `gorm:"column:ExternalSourceListID"`
	PMComment               string         `gorm:"column:PMComment"`
	//base model
	ItemID        int          `gorm:"column:ItemID"`
	IsVersion     bool         `gorm:"column:IsVersion"`
	VersionNumber int          `gorm:"column:VersionNumber"`
	Created       time.Time    `gorm:"column:Created"`
	Modified      time.Time    `gorm:"column:Modified"`
	Author        string       `gorm:"column:Author"`
	AuthorSID     string       `gorm:"column:AuthorSID"`
	Editor        string       `gorm:"column:Editor"`
	EditorSID     string       `gorm:"column:EditorSID"`
	IsDeleted     bool         `gorm:"column:IsDeleted"`
	DeletedDate   sql.NullTime `gorm:"column:DeletedDate"`
	DeletedBy     string       `gorm:"column:DeletedBy"`
	DeletedBySID  string       `gorm:"column:DeletedBySID"`
}

func (tshoursrecord *TSHoursRecord) TableName() string {
	return "tshoursrecord"
}

func (tshoursrecord *TSHoursRecord) ToString() string {
	return fmt.Sprintf("id: %d", tshoursrecord.ID)
}
