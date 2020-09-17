package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"

	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func SelectEmployee() *models.Employees {
	results, err := database.DB.Query("SELECT * FROM employees")
	if err != nil {
		panic(err.Error())
	}

	var (
		EmployeeNumber uint64
		Lastname       string
		Firstname      string
		Extension      string
		Email          string
		OfficeCode     string
		ReportsTo      utils.NullString
		JobTitle       string
	)

	for results.Next() {
		err = results.Scan(&EmployeeNumber, &Lastname, &Firstname, &Extension, &Email, &OfficeCode, &ReportsTo, &JobTitle)

		if err != nil {
			panic(err.Error())
		}
	}
	emp := models.Employees{
		EmployeeNumber: EmployeeNumber,
		Lastname:       Lastname,
		Firstname:      Firstname,
		Extension:      Extension,
		Email:          Email,
		OfficeCode:     OfficeCode,
		ReportsTo:      ReportsTo.String,
		JobTitle:       JobTitle,
	}
	return &emp
}
