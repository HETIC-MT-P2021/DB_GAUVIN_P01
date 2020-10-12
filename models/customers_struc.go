package models

type Customer struct {
	CustomerNumber         uint64  `json:"customerNumber"`
	CustomerName           string  `json:"customerName"`
	ContactFirstName       string  `json:"contactFirstName"`
	ContactLastName        string  `json:"contactLastName"`
	Phone                  string  `json:"phone"`
	AddressLine1           string  `json:"addressLine1"`
	AddressLine2           string  `json:"addressLine2"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	PostCode               string  `json:"postalCode"`
	Country                string  `json:"country"`
	SalesRepEmployeeNumber uint64  `json:"salesRepEmployeeNumber"`
	CreditLimit            float64 `json:"creditLimit"`
	OrderNumber            uint64  `json:"orderNumber"`
	QuantityOrdered        uint64  `json:"quantityOrdered"`
}
