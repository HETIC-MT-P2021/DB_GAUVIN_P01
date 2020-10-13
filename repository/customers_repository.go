package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func FindCustomerByCode(code uint64) (*models.Customer, error) {
	customer := models.Customer{}

	rows := database.DB.QueryRow("SELECT c.customerNumber, c.customerName, c.contactLastName, c.contactFirstName, c.phone, c.addressLine1, c.addressLine2, c.city, c.state, c.postalCode, c.country, c.salesRepEmployeeNumber, c.creditLimit, o.orderNumber, od.quantityOrdered FROM customers as c INNER JOIN orders as o ON c.customerNumber = o.customerNumber INNER JOIN orderdetails as od ON od.orderNumber = o.orderNumber WHERE c.customerNumber = ?;", code)

	err := rows.Scan(
		&customer.CustomerNumber,
		&customer.CustomerName,
		&customer.ContactFirstName,
		&customer.ContactLastName,
		&customer.Phone,
		&customer.AddressLine1,
		&customer.AddressLine2,
		&customer.City,
		&customer.State,
		&customer.PostCode,
		&customer.Country,
		&customer.SalesRepEmployeeNumber,
		&customer.CreditLimit,
		&customer.OrderNumber,
		&customer.QuantityOrdered,
	)

	if err != nil {
		panic(err.Error())
	}

	return &customer, nil
}
