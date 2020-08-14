package models

import (
	"database/sql"
	"fmt"
)

type Employee struct {
	ID int `gorm:"primary_key;auto_increment:true;column:ID"`
	//DepartmentID     int             `gorm:"column:DepartmentID;foreignkey:DepartmentID"`
	//EmployeePositionID     int             `gorm:"column:EmployeePositionID;foreignkey:EmployeePositionID"`
	//EmployeePositionOfficialID     int             `gorm:"column:EmployeePositionOfficialID;foreignkey:EmployeePositionOfficialID"`
	//OrganisationID     int             `gorm:"column:OrganisationID;foreignkey:EmployeePositionOfficialID"`
	//EmployeeLocationID     int             `gorm:"column:OrganisationID;foreignkey:EmployeeLocationID"`
	LastName                  string          `gorm:"column:LastName"`
	FirstName                 string          `gorm:"column:FirstName"`
	MidName                   string          `gorm:"column:MidName"`
	BirthdayDate              sql.NullTime    `gorm:"column:BirthdayDate"`
	EnrollmentDate            sql.NullTime    `gorm:"column:EnrollmentDate"`
	ProbationEndDate          sql.NullTime    `gorm:"column:ProbationEndDate"`
	DismissalDate             sql.NullTime    `gorm:"column:DismissalDate"`
	Email                     string          `gorm:"column:Email"`
	ADLogin                   string          `gorm:"column:ADLogin"`
	DismissalReason           string          `gorm:"column:DismissalReason"`
	EmployeePositionTitle     string          `gorm:"column:EmployeePositionTitle"`
	OfficeName                string          `gorm:"column:OfficeName"`
	WorkPhoneNumber           string          `gorm:"column:WorkPhoneNumber"`
	PersonalMobilePhoneNumber string          `gorm:"column:PersonalMobilePhoneNumber"`
	TSHoursRecord             []TSHoursRecord `gorm:"foreignkey:EmployeeID"`
}

//
//[Display(Name = "Мобильный телефон общедоступный")]
//public string PublicMobilePhoneNumber { get; set; }
//
//[Display(Name = "Skype")]
//public string SkypeLogin { get; set; }
//
//[Display(Name = "Специализация")]
//[DataType(DataType.MultilineText)]
//public string Specialization { get; set; }
//
//[Display(Name = "Примечание")]
//[DataType(DataType.MultilineText)]
//public string Comments { get; set; }
//
//[Display(Name = "Грейд")]
//public int? EmployeeGradID { get; set; }
//public virtual EmployeeGrad EmployeeGrad { get; set; }
//
//[Display(Name = "ADEmployeeID")]
//public string ADEmployeeID { get; set; }
//
//[Display(Name = "Признак вакансии")]
//public bool IsVacancy { get; set; }
//
//[Display(Name = "ID вакансии")]
//public string VacancyID { get; set; }
//
//[Display(Name = "Данные медстраховки (ДМС и/или ОМС)")]
//public string MedicalInsuranceInfo { get; set; }
//
//[Display(Name = "Фактический домашний адрес")]
//public string HomeAddress { get; set; }
//
//[Display(Name = "ФИО контактного лица для связи в случае чрезвычайных ситуаций")]
//public string EmergencyContactName { get; set; }
//
//[Display(Name = "Мобильный телефон контактного лица")]
//public string EmergencyContactMobilePhoneNumber { get; set; }
//
//[Display(Name = "Серия и номер паспорта РФ")]
//public string PassportNumber { get; set; }
//
//[Display(Name = "ФИО в загранпаспорте")]
//public string InternationalPassportName { get; set; }
//
//[Display(Name = "Серия и номер загранпаспорта")]
//public string InternationalPassportNumber { get; set; }
//
//[DataType(DataType.Date)]
//[Display(Name = "Срок действия загранпаспорта")]
//public DateTime? InternationalPassportDueDate { get; set; }
//
//[Display(Name = "Серия и номер иностранного паспорта")]
//public string ForeignPassportNumber { get; set; }
//
//[DataType(DataType.Date)]
//[Display(Name = "Срок действия иностранного паспорта")]
//public DateTime? ForeignPassportDueDate { get; set; }
//
//[Display(Name = "Номер карты авиакомпании")]
//public string AirlineCardInfo { get; set; }

func (employee Employee) TableName() string {
	return "employee"
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("id: %d", employee.ID)
}
