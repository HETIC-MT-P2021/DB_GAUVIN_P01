package repository

import (
	"time"

	"github.com/jasongauvin/DB_GAUVIN_P01/utils"

	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func FindOrderByCode(code uint64) (*[]models.OrderProduct, error) {
	var (
		OrderNumber        uint64
		OrderDate          time.Time
		RequiredDate       time.Time
		ShippedDate        time.Time
		Status             string
		Comments           utils.NullString
		CustomerNumber     uint64
		OrderDetailNumber  uint64
		ProductDetailCode  string
		QuantityOrdered    uint64
		PriceEach          float64
		OrderLineNumber    uint8
		OrderPrice         float64
		ProductCode        string
		ProductName        string
		ProductLine        string
		ProductScale       string
		ProductVendor      string
		ProductDescription string
		QuantityInStock    uint64
		BuyPrice           float64
		MSRP               float64
	)

	rows, err := database.DB.Query("SELECT o.orderNumber, o.orderDate, o.requiredDate, o.shippedDate, o.status, o.comments, o.customerNumber, od.orderNumber, od.productCode, od.quantityOrdered, od.priceEach, od.orderLineNumber, od.quantityOrdered * od.priceEach AS orderPrice, p.productCode, p.productName, p.productLine, p.productScale, p.productVendor, p.productDescription, p.quantityInStock, p.buyPrice, p.MSRP FROM orders as o INNER JOIN orderdetails as od ON o.orderNumber = od.orderNumber INNER JOIN products as p ON od.productCode = p.productCode WHERE o.orderNumber = ?;", code)

	var orderAndProducts []models.OrderProduct

	for rows.Next() {
		err = rows.Scan(
			&OrderNumber,
			&OrderDate,
			&RequiredDate,
			&ShippedDate,
			&Status,
			&Comments,
			&CustomerNumber,
			&OrderDetailNumber,
			&ProductDetailCode,
			&QuantityOrdered,
			&PriceEach,
			&OrderLineNumber,
			&OrderPrice,
			&ProductCode,
			&ProductName,
			&ProductLine,
			&ProductScale,
			&ProductVendor,
			&ProductDescription,
			&QuantityInStock,
			&BuyPrice,
			&MSRP,
		)

		orderAndProduct := models.OrderProduct{
			models.Order{
				OrderNumber:       OrderNumber,
				OrderDate:         OrderDate,
				RequiredDate:      RequiredDate,
				ShippedDate:       ShippedDate,
				Status:            Status,
				Comments:          Comments.String,
				CustomerNumber:    CustomerNumber,
				OrderDetailNumber: OrderDetailNumber,
				ProductDetailCode: ProductDetailCode,
				QuantityOrdered:   QuantityOrdered,
				PriceEach:         PriceEach,
				OrderLineNumber:   OrderLineNumber,
				OrderPrice:        OrderPrice,
			},
			models.Product{
				ProductCode:        ProductCode,
				ProductName:        ProductName,
				ProductLine:        ProductLine,
				ProductScale:       ProductScale,
				ProductVendor:      ProductVendor,
				ProductDescription: ProductDescription,
				QuantityInStock:    QuantityInStock,
				BuyPrice:           BuyPrice,
				MSRP:               MSRP,
			},
		}
		orderAndProducts = append(orderAndProducts, orderAndProduct)
	}

	if err != nil {
		panic(err.Error())
	}

	return &orderAndProducts, nil
}
