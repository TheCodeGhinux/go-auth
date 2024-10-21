package routing

import (
	"fmt"
	"log"

	"github.com/TheCodeGhinux/go-auth/pkg/config"
	"github.com/gin-gonic/gin"
)

func Route() {

	configs := config.LoadConfig()
	r := gin.Default()

	RouteRegister(r)

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		log.Fatal("Error starting server in routing: ", err)
		return
	}

	
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

}