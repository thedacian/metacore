package main

//DB operations module

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DbInit sets up the database
func DbInit() {
	db, err := gorm.Open("sqlite3", "db.db")
	defer db.Close()
	if err != nil {
		handleError(err)
	}

	//Create a record in the DB for testing purposes
	db.Create(&User{Username: "admin", Password: "password"})

	//Create the User DB table using the provided model
	db.AutoMigrate(&User{})
}

//DbCheckUser checks if a supplied record is present in the DB and returns true
func DbCheckUser(username, password string) bool {
	db, err := gorm.Open("sqlite3", "db.db")
	defer db.Close()
	if err != nil {
		handleError(err)
	}

	// Instantiate var in memory to store the record
	var usr User

	if usr.Password != password {
		log.Print("DbCheckUser - Valid")
		return false
	} else {
		log.Print("DbCheckUser - Invalid")
		return true
	}
}
