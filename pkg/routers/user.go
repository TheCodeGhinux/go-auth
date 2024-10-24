package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/go-auth/pkg/controllers/auth"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine, ApiVersion string, db *db.Database) *gin.Engine {
	userController := controller.UserController{Db: db}

	userGroup := router.Group(fmt.Sprintf("%v/auth", ApiVersion))
	{
		userGroup.GET("users/:id", userController.FindUserById)
		// authGroup.POST("/register", controllers.Register)
	}
	return router

}
