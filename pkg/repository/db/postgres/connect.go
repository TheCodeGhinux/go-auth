package postgres

import (
	"fmt"
	"log"

	"github.com/TheCodeGhinux/go-auth/config"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(config config.DB) *gorm.DB {

	dbCfg := config
	// utils.LoggerInstance.Info("connecting to database")
	fmt.Println("Connecting to PostgreSQL database...")

	connectedDB := connectToDb(dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Dbname, dbCfg.Port, "disable", "UTC")

	//  utils.LoggerInstance.Info("connected to database")
	fmt.Println("Connected to PostgreSQL database")

    // Assign the connected PostgreSQL database to the global DB variable    
		db.DB.Postgres = connectedDB

	return connectedDB
}

func connectToDb(host, user, password, dbname, port, sslmode, timezone string) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection to %v db failed with: %v", dbname, err)
		panic(err)
	}

	// logger.Info("Connected to %v db", dbname)
	fmt.Sprintf("Connected to %v db", dbname)
	// db = db.Debug() // Uncomment if you want to enable debug mode
	return db
}
