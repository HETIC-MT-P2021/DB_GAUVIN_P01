package models

import (
	"time"

	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

type Order struct {
	OrderNumber       uint64           `json:"orderNumber"`
	OrderDate         time.Time        `json:"orderDate"`
	RequiredDate      time.Time        `json:"requiredDate"`
	ShippedDate       time.Time        `json:"shippedDate"`
	Status            string           `json:"status"`
	Comments          utils.NullString `json:"comments"`
	CustomerNumber    uint64           `json:"customerNumber"`
	OrderDetailNumber uint64           `json:"orderDetailNumber"`
	ProductDetailCode string           `json:"productCode"`
	QuantityOrdered   uint64           `json:"quantityOrdered"`
	PriceEach         float64          `json:"priceEach"`
	OrderLineNumber   uint8            `json:"orderLineNumber"`
	OrderPrice        float64          `json:"orderPrice"`
}

type OrderProduct struct {
	Order
	Product
}
