package main

import (
	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/pkg/config"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db/postgres"
	"github.com/TheCodeGhinux/go-auth/pkg/routing"
	// "github.com/TheCodeGhinux/go-auth/utils"
)

func main() {
	// Initialize the logger
	// logger := utils.InitLogger()

	configs := config.LoadConfig()

	postgres.ConnectDb(configs.DB)
	models.MigrateDb()
	routing.Route()
}
