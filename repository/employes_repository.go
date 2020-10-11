package repository

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/database"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"

	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func FindEmployees() ([]models.Employee, error) {

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
		//panic(err.Error())
	}

	return employees, nil
}

func FindEmployeesByOffice(code uint64) ([]*models.Employee, error) {
	var (
		EmployeeNumber                                              uint64
		Lastname, Firstname, Extension, Email, OfficeCode, JobTitle string
		ReportsTo                                                   utils.NullString
	)

	rows, err := database.DB.Query("SELECT e.employeeNumber, e.firstName, e.lastName, e.extension, e.email, e.officeCode, e.jobTitle, e.reportsTo FROM employees e INNER JOIN offices o ON e.officeCode = o.officeCode WHERE o.officeCode = ?;", code)

	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

	var employees []*models.Employee

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

		employee := &models.Employee{
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
		//panic(err.Error())
	}

	return employees, nil
}
