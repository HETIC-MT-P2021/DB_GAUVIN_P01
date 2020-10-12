package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func FindCustomerByCode(code uint64) (*models.Customer, error) {
	var (
		CustomerNumber         uint64
		CustomerName           string
		ContactFirstName       string
		ContactLastName        string
		Phone                  string
		AddressLine1           string
		AddressLine2           utils.NullString
		City                   string
		State                  utils.NullString
		PostCode               utils.NullString
		Country                string
		SalesRepEmployeeNumber uint64
		CreditLimit            float64
		OrderNumber            uint64
		QuantityOrdered        uint64
	)

	rows := database.DB.QueryRow("SELECT c.customerNumber, c.customerName, c.contactLastName, c.contactFirstName, c.phone, c.addressLine1, c.addressLine2, c.city, c.state, c.postalCode, c.country, c.salesRepEmployeeNumber, c.creditLimit, o.orderNumber, od.quantityOrdered FROM customers as c INNER JOIN orders as o ON c.customerNumber = o.customerNumber INNER JOIN orderdetails as od ON od.orderNumber = o.orderNumber WHERE c.customerNumber = ?;", code)

	err := rows.Scan(
		&CustomerNumber,
		&CustomerName,
		&ContactFirstName,
		&ContactLastName,
		&Phone,
		&AddressLine1,
		&AddressLine2,
		&City,
		&State,
		&PostCode,
		&Country,
		&SalesRepEmployeeNumber,
		&CreditLimit,
		&OrderNumber,
		&QuantityOrdered,
	)

	customer := models.Customer{
		CustomerNumber:         CustomerNumber,
		CustomerName:           CustomerName,
		ContactFirstName:       ContactFirstName,
		ContactLastName:        ContactLastName,
		Phone:                  Phone,
		AddressLine1:           AddressLine1,
		AddressLine2:           AddressLine2.String,
		City:                   City,
		State:                  State.String,
		PostCode:               PostCode.String,
		Country:                Country,
		SalesRepEmployeeNumber: SalesRepEmployeeNumber,
		CreditLimit:            CreditLimit,
		OrderNumber:            OrderNumber,
		QuantityOrdered:        QuantityOrdered,
	}

	if err != nil {
		panic(err.Error())
	}

	return &customer, nil
}
