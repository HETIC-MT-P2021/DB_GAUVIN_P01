package models

type Employees struct {
	EmployeeNumber uint64 `json:"employeeNumber"`
	Lastname       string `json:"lastName"`
	Firstname      string `json:"firstName"`
	Extension      string `json:"extention"`
	Email          string `json:"email"`
	OfficeCode     string `json:"officeCode"`
	ReportsTo      string `json:"reportsCode"`
	JobTitle       string `json:"jobTitle"`
}
