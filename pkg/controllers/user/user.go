package user

import (
	"net/http"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	service "github.com/TheCodeGhinux/go-auth/services/user"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *db.Database
}

func (uc *UserController) FindUserById(c *gin.Context) {

	id := c.Param("id")
	message, statusCode, user, err := service.GetUserById(c, id, db.DB.Postgres)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		return
	}

	utils.RespondHandler(c, message, statusCode, user)
}

func (uc *UserController) GetUser(c *gin.Context) {

	user, _ := c.Get("user")

	userId := user.(*models.User).ID
	message, statusCode, user, err := service.GetUserById(c, userId, db.DB.Postgres)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		return
	}

	utils.RespondHandler(c, message, statusCode, user)

}
