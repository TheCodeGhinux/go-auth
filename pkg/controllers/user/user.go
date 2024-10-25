package user

import (
	"net/http"

	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	service "github.com/TheCodeGhinux/go-auth/services/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *db.Database
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