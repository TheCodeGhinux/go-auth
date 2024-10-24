package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	Name        string         `gorm:"column:name; type:varchar(255)" json:"name"`
	Email       string         `gorm:"column:email; type:varchar(255)" json:"email"`
	IsVerified  bool           `gorm:"column:is_verified; type:bool" json:"is_verified"`
	Deactivated bool           `gorm:"column:deactivated; type:bool" json:"deactivated"`
	IsActive    bool           `gorm:"column:is_active; type:bool; default:false" json:"is_active"`
	CurrentOrg  uuid.UUID      `gorm:"column:current_org;null; type:uuid" json:"current_org"`
	Profile     Profile        `gorm:"foreignKey:Userid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
	UserRoleID  *string        `gorm:"type:varchar(100);null;index" json:"user_role_id"`
	Password    string         `gorm:"column:password; type:text; not null" json:"-"`
	CreatedAt   time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Role        int            `gorm:"column:role" json:"role"`
}

type CreateUserRequestModel struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name" `
}

