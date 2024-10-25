package auth

import (
	"net/http"

	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	service "github.com/TheCodeGhinux/go-auth/services/auth"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *db.Database
}


func (uc *UserController) RegisterUser(c *gin.Context) {

	user, err := service.RegisterUser(c, db.DB.Postgres)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if user == nil {
        return
    }

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) LoginUser(c *gin.Context) {

	message, statusCode, user, err := service.LoginUser(c, db.DB.Postgres)

	    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if user == nil {
        return
    }

		utils.RespondHandler(c, message, statusCode, user)

}