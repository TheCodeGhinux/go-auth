package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/pkg/config"
	"github.com/TheCodeGhinux/go-auth/pkg/repository/db"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuth(c *gin.Context) {
	configs := config.LoadConfig()

	// Get token from cookie
	tokenString, err := c.Cookie("access_token")

	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "User unauthorized, please sign in to continue")
		c.Abort()
		return
	}

	if tokenString == "" {
		utils.SendError(c, http.StatusUnauthorized, "User unauthorized, please sign in to continue")
		c.Abort()
		return
	}

	// Decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Token.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token has expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.SendError(c, http.StatusUnauthorized, "Token has expired, please sign in again")
			c.Abort()
			return
		}

		userId, ok := claims["userId"].(string)
		if !ok {
			utils.SendError(c, http.StatusUnauthorized, "Invalid token data, please sign in to continue")
			c.Abort()
			return
		}

		// Find user with token userId
		user, err := models.FindUserByID(userId, db.DB.Postgres)

		if err != nil || user == nil {
			utils.SendError(c, http.StatusUnauthorized, "User not found, please sign in to continue")
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)
		c.Next()
	} else {
		utils.SendError(c, http.StatusUnauthorized, "Invalid token, please sign in to continue")
		c.Abort()
		return
	}
}
