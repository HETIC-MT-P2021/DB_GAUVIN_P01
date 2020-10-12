package models

type Product struct {
	ProductCode        string  `json:"productCode"`
	ProductName        string  `json:"productName"`
	ProductLine        string  `json:"productLine"`
	ProductScale       string  `json:"productScale"`
	ProductVendor      string  `json:"productVendor"`
	ProductDescription string  `json:"productDescription"`
	QuantityInStock    uint64   `json:"quantityInStock"`
	BuyPrice           float64 `json:"buyPrice"`
	MSRP               float64 `json:"MSRP"`
}
