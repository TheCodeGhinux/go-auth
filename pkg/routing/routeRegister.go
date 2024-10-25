package routing

import (
	"github.com/TheCodeGhinux/go-auth/pkg/controllers/greeting"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"github.com/TheCodeGhinux/go-auth/pkg/routers"
	"github.com/gin-gonic/gin"
)

func RouteRegister(router *gin.Engine) {

	apiVersion := "api/v1"
	greeting.Greeting(router)
	routers.Auth(router, apiVersion, db.DB)
	routers.User(router, apiVersion, db.DB)

}
