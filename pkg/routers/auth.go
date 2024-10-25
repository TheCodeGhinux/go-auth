package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/go-auth/pkg/controllers/auth"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.Engine, ApiVersion string, db *db.Database) *gin.Engine  {
	authController := controller.UserController{Db: db}

	authGroup := router.Group(fmt.Sprintf("%v/auth", ApiVersion))
	{
		authGroup.POST("/register", authController.RegisterUser)
	}
	return router

}