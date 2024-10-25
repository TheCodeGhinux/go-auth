package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TheCodeGhinux/go-auth/internal/models"
	"github.com/TheCodeGhinux/go-auth/pkg/middlewares"
	"github.com/TheCodeGhinux/go-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	g          = galidator.New()
	customizer = g.Validator(models.CreateUserRequestModel{})
)

func RegisterUser(c *gin.Context, db *gorm.DB) (*models.User, error) {
	body := &models.CreateUserRequestModel{}

	// Bind and validate input
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, customizer.DecryptErrors(err))
		return nil, nil
	}

	// Hash the password
	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to hash password")
		return nil, nil
	}

	username, err := generateUniqueUsername(body.FirstName, body.Email, db)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return nil, nil
	}

	user := createUser(body, hashedPassword, username)

	// Check if user already exists
	if err := checkUserExistence(user.Email, db); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to check user existence")
		return nil, nil
	}

	// Insert the new user into the database
	if err := models.CreateUser(user, db); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Failed to create user")
		return nil, nil
	}

	return user, nil
}

func LoginUser(c *gin.Context, db *gorm.DB) (string, int, gin.H, error) {
	body := &models.LoginUserRequestModel{}

	// Bind and validate input
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, customizer.DecryptErrors(err))
		return "GGH", 200, nil, nil
	}

	// Find the user by email
	user, _ := models.FindUserByEmail(body.Email, db)

	if user == nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return "", 0, nil, nil
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid email or password")
		return "", 0, nil, nil
	}

	// Generate a new JWT token
	token, err := middlewares.GenerateToken(user)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate token")
		return "", 0, nil, nil 
	}

	middlewares.SetCookie(c, token.AccessToken)

	responseData := gin.H{
		"access_token": token.AccessToken,
		"user":  user,
	}
	return "Login successful", http.StatusOK, responseData, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func generateUniqueUsername(firstName, email string, db *gorm.DB) (string, error) {
	username := strings.Split(email, "@")[0]

	// Regenerate username if it exists
	for {
		existingUsername, err := models.FindUserUsername(username, db)
		if err != nil {
			return "", err
		}

		if existingUsername == nil {
			// Username is unique; return it
			return username, nil
		}
		// Regenerate username if it exists
		username = utils.GenerateUsername(firstName)
	}

	// return "", nil
}

func createUser(body *models.CreateUserRequestModel, password, username string) *models.User {
	return &models.User{
		ID:        utils.GenerateUUID(),
		Email:     body.Email,
		Name:      body.FirstName + " " + body.LastName,
		Password:  password,
		IsActive:  false,
		CreatedAt: time.Now(),
		Profile: models.Profile{
			ID:        utils.GenerateUUID(),
			FirstName: body.FirstName,
			LastName:  body.LastName,
			UserName:  username,
			Phone:     body.PhoneNumber,
			CreatedAt: time.Now(),
		},
	}
}

func checkUserExistence(email string, db *gorm.DB) error {
	existingUser, err := models.FindUserByEmail(email, db)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("user already exists")
	}
	return nil
}

