package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitializeDb connects to database
func InitializeDb() {
	dbDriver := "mysql"
	dbUser := "gobdd"
	dbPass := "gobdd"
	dbName := "image_gobdd"
	dbHost := "db" // container = ip du container
	dbPort := 3306

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err := sql.Open(dbDriver, dsn)
	fmt.Println(DB)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

}
