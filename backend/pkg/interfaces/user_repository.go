package interfaces

import "smart-learning-backend/pkg/models"

// UserRepositoryInterface 定義用戶倉庫的介面
type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CheckUserExists(email, username string) (bool, error)
}

// AuthServiceInterface 定義認證服務的介面
type AuthServiceInterface interface {
	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(req *models.LoginRequest) (*models.AuthResponse, error)
	GetUserByID(id int) (*models.User, error)
}