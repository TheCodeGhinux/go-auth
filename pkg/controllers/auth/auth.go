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

	id := c.Param("userId")
	user, err := service.GetUser(c, id, db.DB.Postgres)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
