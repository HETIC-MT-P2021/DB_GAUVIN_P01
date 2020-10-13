package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func FindOrderByCode(code uint64) (*[]models.OrderProduct, error) {
	orderAndProduct := models.OrderProduct{}

	rows, err := database.DB.Query("SELECT o.orderNumber, o.orderDate, o.requiredDate, o.shippedDate, o.status, o.comments, o.customerNumber, od.orderNumber, od.productCode, od.quantityOrdered, od.priceEach, od.orderLineNumber, od.quantityOrdered * od.priceEach AS orderPrice, p.productCode, p.productName, p.productLine, p.productScale, p.productVendor, p.productDescription, p.quantityInStock, p.buyPrice, p.MSRP FROM orders as o INNER JOIN orderdetails as od ON o.orderNumber = od.orderNumber INNER JOIN products as p ON od.productCode = p.productCode WHERE o.orderNumber = ?;", code)

	var orderAndProducts []models.OrderProduct

	for rows.Next() {

		err = rows.Scan(
			&orderAndProduct.OrderNumber,
			&orderAndProduct.OrderDate,
			&orderAndProduct.RequiredDate,
			&orderAndProduct.ShippedDate,
			&orderAndProduct.Status,
			&orderAndProduct.Comments,
			&orderAndProduct.CustomerNumber,
			&orderAndProduct.OrderDetailNumber,
			&orderAndProduct.ProductDetailCode,
			&orderAndProduct.QuantityOrdered,
			&orderAndProduct.PriceEach,
			&orderAndProduct.OrderLineNumber,
			&orderAndProduct.OrderPrice,
			&orderAndProduct.ProductCode,
			&orderAndProduct.ProductName,
			&orderAndProduct.ProductLine,
			&orderAndProduct.ProductScale,
			&orderAndProduct.ProductVendor,
			&orderAndProduct.ProductDescription,
			&orderAndProduct.QuantityInStock,
			&orderAndProduct.BuyPrice,
			&orderAndProduct.MSRP,
		)

		orderAndProducts = append(orderAndProducts, orderAndProduct)
	}

	if err != nil {
		panic(err.Error())
	}

	return &orderAndProducts, nil
}
