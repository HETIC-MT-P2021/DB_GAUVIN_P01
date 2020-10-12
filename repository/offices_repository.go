package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func FindOfficeByCode(code uint64) (*models.Office, error) {
	var (
		OfficeCode                                                                     uint64
		City, Phone, AddressLine1, AddressLine2, State, Country, PostalCode, Territory string
	)

	rows := database.DB.QueryRow("SELECT o.officeCode, o.city, o.phone, o.addressLine1, o.addressLine2, o.state, o.country, o.postalCode, o.territory FROM offices o WHERE o.officeCode = ?;", code)

	err := rows.Scan(
		&OfficeCode,
		&City,
		&Phone,
		&AddressLine1,
		&AddressLine2,
		&State,
		&Country,
		&PostalCode,
		&Territory,
	)

	office := models.Office{
		OfficeCode:   OfficeCode,
		City:         City,
		Phone:        Phone,
		AddressLine1: AddressLine1,
		AddressLine2: AddressLine2,
		State:        State,
		Country:      Country,
		PostalCode:   PostalCode,
		Territory:    Territory,
	}

	if err != nil {
		panic(err.Error())
	}

	return &office, nil
}
