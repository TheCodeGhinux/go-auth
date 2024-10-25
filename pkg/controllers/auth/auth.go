package auth

import (
	"net/http"

	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	service "github.com/TheCodeGhinux/go-auth/services/auth"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *db.Database
}

func CreateUser(c *gin.Context) {

}

func (uc *UserController) FindUserById(c *gin.Context) {

	id := c.Param("id")
	user, err := service.GetUser(c, id, db.DB.Postgres)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		return
	}

	c.JSON(http.StatusOK, user)
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
