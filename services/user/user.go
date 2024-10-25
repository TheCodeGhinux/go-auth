package user

import (
	"fmt"
	"net/http"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context, id string, db *gorm.DB) (*models.User, error) {
	if id == "" {
		fmt.Printf("Id found: %+v\n", id)
		utils.SendError(c, http.StatusBadRequest, "Invalid user ID, please provide a valid user ID")
		return nil, nil
	}

	user, _ := models.FindUserByID(id, db)
	if user == nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return nil, nil
	}

	return user, nil
}