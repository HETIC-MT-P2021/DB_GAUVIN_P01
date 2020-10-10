package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"

	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func SelectEmployees() ([]models.Employee, error) {

	var (
		EmployeeNumber                                              uint64
		Lastname, Firstname, Extension, Email, OfficeCode, JobTitle string
		ReportsTo                                                   utils.NullString
	)

	rows, err := database.DB.Query("SELECT * FROM employees")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		err = rows.Scan(
			&EmployeeNumber,
			&Lastname,
			&Firstname,
			&Extension,
			&Email,
			&OfficeCode,
			&JobTitle,
			&ReportsTo)

		employee := models.Employee{
			EmployeeNumber: EmployeeNumber,
			Lastname:       Lastname,
			Firstname:      Firstname,
			Extension:      Extension,
			Email:          Email,
			OfficeCode:     OfficeCode,
			JobTitle:       JobTitle,
			ReportsTo:      ReportsTo.String,
		}

		employees = append(employees, employee)
	}
	if err != nil {
		panic(err.Error())
	}

	return employees, err
}
