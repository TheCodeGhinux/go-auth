package greeting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Greeting(r *gin.Engine) {
	r.GET("/greeting", greet)
	r.GET("/", greet)
}


	func greet(c *gin.Context) {
		appName := viper.GetString("App.name")
		appDesc := viper.GetString("App.desc")
		c.JSON(http.StatusOK, gin.H{
			"message":         "pong",
			"app name":        "Welcome to " + appName,
			"app description": appDesc,
		})
	}