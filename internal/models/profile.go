package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        string         `gorm:"type:uuid;primary_key" json:"profile_id"`
	FirstName string         `gorm:"column:first_name; type:text; not null" json:"first_name"`
	LastName  string         `gorm:"column:last_name; type:text;not null" json:"last_name"`
	FullName  string         `gorm:"column:full_name; type:text;" json:"full_name"`
	UserName  string         `gorm:"column:user_name; type:text;" json:"user_name"`
	Phone     string         `gorm:"type:varchar(255)" json:"phone"`
	AvatarURL string         `gorm:"type:varchar(255)" json:"avatar_url"`
	Userid    string         `gorm:"type:uuid;" json:"user_id"`
	CreatedAt time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func FindUserUsername(user_name string, db *gorm.DB) (*Profile, error) {
	var profile Profile
	result := db.Where("user_name = ?", user_name).First(&profile)

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
	fmt.Printf("User found: %+v\n", profile)
	return &profile, nil
}
