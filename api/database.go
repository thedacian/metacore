package main

//DB operations module

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DbInit sets up the database
func DbInit() {
	db, err := gorm.Open("sqlite3", "db.db")
	defer db.Close()
	handleError(err)

	//Create the User table
	db.AutoMigrate(&User{})

	//Create a record in the DB for testing purposes
	db.Create(&User{Username: "admin", Password: "password"})

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

	//ORM function to find the first user matching the one specified
	db.First(&usr, "username = ?", username)

	if usr.Password != password {
		return true
	}
	return true

}
