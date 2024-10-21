package main

import (
	"github.com/TheCodeGhinux/go-auth/pkg/config"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db/postgres"
	"github.com/TheCodeGhinux/go-auth/pkg/routing"
	// "github.com/TheCodeGhinux/go-auth/utils"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })	
	// r.Run()
		// Initialize the logger
	// logger := utils.InitLogger()

		configs := config.LoadConfig()

	postgres.ConnectDb( configs.DB)
	routing.Route()
}
