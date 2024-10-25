package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/go-auth/pkg/controllers/user"
	"github.com/TheCodeGhinux/go-auth/pkg/middlewares"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine, ApiVersion string, db *db.Database) *gin.Engine {
	userController := controller.UserController{Db: db}

	userGroup := router.Group(fmt.Sprintf("%v/users", ApiVersion))
	{
		userGroup.GET("/:id", userController.FindUserById)
	}

	// Protected route group
	userProtected := router.Group(fmt.Sprintf("%v/users", ApiVersion))
	userProtected.Use(middlewares.UserAuth)
	{
		userProtected.GET("/", userController.GetUser)
	}

	return router

}
