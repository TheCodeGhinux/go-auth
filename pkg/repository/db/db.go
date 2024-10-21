package db

import (
	"github.com/TheCodeGhinux/go-auth/utils"
	"gorm.io/gorm"
)

type Database struct {
	Postgres *gorm.DB
}

var (
	DB     *Database = &Database{}
	Logger *utils.Logger
)

func Connect() *Database {
	return DB
}