package models

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

type Employee struct {
	EmployeeNumber uint64           `json:"employeeNumber"`
	LastName       string           `json:"lastName"`
	FirstName      string           `json:"firstName"`
	Extension      string           `json:"extention"`
	Email          string           `json:"email"`
	OfficeCode     string           `json:"officeCode"`
	JobTitle       string           `json:"jobTitle"`
	ReportsTo      utils.NullString `json:"reportsCode"`
}
