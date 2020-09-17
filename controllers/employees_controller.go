package controllers

import (
	"fmt"
	"net/http"

	"github.com/jasongauvin/DB_GAUVIN_P01/repository"
)

func ShowEmployee(w http.ResponseWriter, r *http.Request) {
	employee := repository.SelectEmployee()
	fmt.Println(employee)
}
