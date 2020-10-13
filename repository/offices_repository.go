package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func FindOfficeByCode(code uint64) (*models.Office, error) {
	office := models.Office{}

	rows := database.DB.QueryRow("SELECT o.officeCode, o.city, o.phone, o.addressLine1, o.addressLine2, o.state, o.country, o.postalCode, o.territory FROM offices o WHERE o.officeCode = ?;", code)

	err := rows.Scan(
		&office.OfficeCode,
		&office.City,
		&office.Phone,
		&office.AddressLine1,
		&office.AddressLine2,
		&office.State,
		&office.Country,
		&office.PostalCode,
		&office.Territory,
	)

	if err != nil {
		panic(err.Error())
	}

	return &office, nil
}
