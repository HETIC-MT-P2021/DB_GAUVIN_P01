package models

type Customer struc {
	Number                 uint64 	`json:"customerNumber"`
	Name                   string 	`json:"customerName"`
	FirstName              string 	`json:"customerFirstName"`
	LastName               string 	`json:"customerLastName"`
	Phone                  string 	`json:"phone"`
	AddressLine1           string 	`json:"addressLine1"`
	AddressLine2           string 	`json:"addressLine2"`
	City                   string 	`json:"city"`
	State                  string 	`json:"state"`
	PostCode               string 	`json:"postalCode"`
	Country                string 	`json:"country"`
	SalesRepEmployeeNumber uint64 	`json:"salesRepEmployeeNumber"`
	CreditLimit            float64 	`json:"creditLimit"`
}