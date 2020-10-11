package models

type Office struct {
	OfficeCode   uint64      `json:"officeCode"`
	City         string      `json:"city"`
	Phone        string      `json:"phone"`
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 string      `json:"addressLine2"`
	State        string      `json:"state"`
	Country      string      `json:"country"`
	PostalCode   string      `json:"postalCode"`
	Territory    string      `json:"territory"`
}
