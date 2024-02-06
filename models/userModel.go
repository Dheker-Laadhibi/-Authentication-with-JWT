package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           int       `json:"id"`
	FirstName    *string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string    `json:"last_name" validate:"required,min=2,max=100"`
	Password     *string    `json:"password" validate:"required,min=6,max=100"`
	Email        *string    `json:"email" validate:"required,email"`
	Phone        *string    `json:"phone" validate:"required,min=4"`
	Token        *string    `json:"token"`
	UserType     *string    `json:"user_type" validate:"required,oneof=admin,user"`
	RefreshToken *string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       int       `json:"user_id"`
}
