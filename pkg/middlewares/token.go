package middlewares

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenDto struct {
	AccessToken string `json:"access_token"`
}

func GenerateToken(user *models.User) (*TokenDto, error) {
	configs := config.LoadConfig()

	expiryTime, err := strconv.Atoi(configs.Token.Duration)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return nil, err
	}
	fmt.Println("Expiry time in minutes:", expiryTime)

	// Set up user claims
	userClaims := jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Minute * time.Duration(expiryTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(configs.Token.Secret))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return nil, err // Return an error if signing fails
	}

	// Return the access token in a TokenDto
	return &TokenDto{AccessToken: tokenString}, nil

}

func SetCookie(c *gin.Context, tokenString string) {
	c.SetCookie("access_token", tokenString, 3600, "/", "", false, true)
}

// func VerifyToken(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("your-secret-key"), nil
// 	})
// }
