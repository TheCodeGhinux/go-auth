package models

import "github.com/TheCodeGhinux/go-auth/pkg/repository/db"

func MigrateDb() {
	// db := db.Connect()
		if err := db.DB.Postgres.AutoMigrate(&User{}, &Profile{}); err != nil {
		// Handle the error appropriately; for example, log it or return it
		panic("Failed to migrate database: " + err.Error())
	}

	
}