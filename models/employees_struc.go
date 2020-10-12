package models

type Employee struct {
	EmployeeNumber uint64 `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string `json:"extention"`
	Email          string `json:"email"`
	OfficeCode     string `json:"officeCode"`
	JobTitle       string `json:"jobTitle"`
	ReportsTo      string `json:"reportsCode"`
}
