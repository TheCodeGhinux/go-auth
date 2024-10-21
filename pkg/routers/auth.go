package routers

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.Engine, ApiVersion string,) *gin.Engine {

	// authGroup := router.Group(fmt.Sprintf("%v/auth", ApiVersion))
	{
		// authGroup.POST("/login", controllers.Login)
		// authGroup.POST("/register", controllers.Register)
	}
	return router

}