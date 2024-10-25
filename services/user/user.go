package user

import (
	"fmt"
	"net/http"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserById(c *gin.Context, id string, db *gorm.DB) (string, int, *models.User, error) {
	if id == "" {
		fmt.Printf("Id found: %+v\n", id)
		utils.SendError(c, http.StatusBadRequest, "Invalid user ID, please provide a valid user ID")
		return "", 0, nil, nil
	}

	user, _ := models.FindUserByID(id, db)
	if user == nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return "", 0, nil, nil
	}

	return "User fetched successfully", http.StatusOK, user, nil
}