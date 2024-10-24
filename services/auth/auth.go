package auth

import (
	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context, id string, db *gorm.DB) (models.User, error){
	
	var userRepo models.User
	
	if err := db.First(&userRepo, id).Error; err != nil {
		return models.User{}, err
	}	
		return userRepo, nil
}