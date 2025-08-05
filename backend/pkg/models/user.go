package models

import (
	"time"
)

type User struct {
	ID            int       `json:"id" db:"id"`
	Email         string    `json:"email" db:"email"`
	Username      string    `json:"username" db:"username"`
	PasswordHash  string    `json:"-" db:"password_hash"`
	LearningLevel int       `json:"learning_level" db:"learning_level"`
	AvatarURL     *string   `json:"avatar_url" db:"avatar_url"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required,min=2,max=20,alphanum"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AuthResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}