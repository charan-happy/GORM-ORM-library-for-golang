package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

// opening connection to Database
func DBConnect() {
	// Data Source Name configuration
	dsn := "host=localhost user=postgres password=admin dbname=GoDB port=5432 sslmode=disable"
	// opening connection to PostGreSQL Database
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection to Database Established Successfully!")
	}

	sql, err := Db.DB()

	if err != nil {
		panic(err.Error())
	}
	if err := sql.Ping(); err != nil {
		panic(err.Error())
	}
}

// closing connection to Database
func DBClose() {
	//closing connection to PostGreSQL Database
	postgreDB, _ := Db.DB()
	err := postgreDB.Close()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection to Database Closed Successfully!")
	}
}

func main() {
	DBConnect()
	DBClose()
}
