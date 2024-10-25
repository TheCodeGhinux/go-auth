package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

type User struct {
	ID          string         `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	Name        string         `gorm:"column:name; type:varchar(255)" json:"name"`
	Email       string         `gorm:"column:email; type:varchar(255)" json:"email"`
	IsVerified  bool           `gorm:"column:is_verified; type:bool" json:"is_verified"`
	Deactivated bool           `gorm:"column:deactivated; type:bool" json:"deactivated"`
	IsActive    bool           `gorm:"column:is_active; type:bool; default:false" json:"is_active"`
	Profile     Profile        `gorm:"foreignKey:Userid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
	Password    string         `gorm:"column:password; type:text; not null" json:"-"`
	CreatedAt   time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Role        string         `gorm:"column:role" json:"role"`
}

type CreateUserRequestModel struct {
	Email       string `json:"email" binding:"required,email" required:"Email is required"`
	Password    string `json:"password" binding:"required,min=8" required:"Password is required and must be at least 8 characters long"`
	FirstName   string `json:"first_name" binding:"required" required:"First name is required"`
	LastName    string `json:"last_name" binding:"required" required:"Last name is required"`
	PhoneNumber string `json:"phone_number" binding:"required" required:"Phone number is required"`
	UserName    string `json:"user_name"`
}

// CreateUser inserts a new user into the database
func CreateUser(user *User, db *gorm.DB) error {
	result := db.Create(user)
	return result.Error
}

func FindUserByID(id string, db *gorm.DB) (*User, error) {
	var user User
		result := db.Preload("Profile").Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	// Log the found user
	fmt.Printf("User found: %+v\n", user)
	return &user, nil
}

//	func FindUserByEmail(email string, user *User, db *gorm.DB) error {
//		result := db.Where("email = ?", email).First(user)
//		fmt.Printf("Existing user found: %+v\n", result)
//		return result.Error
//	}
func FindUserByEmail(email string, db *gorm.DB) (*User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// No user found
			return nil, nil
		}
		// Log and return unexpected errors
		fmt.Printf("Unexpected error: %v\n", result.Error)
		return nil, result.Error
	}

	// Log the found user
	fmt.Printf("User found: %+v\n", user)
	return &user, nil
}
