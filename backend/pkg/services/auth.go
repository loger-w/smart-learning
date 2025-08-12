package services

import (
	"fmt"
	"regexp"
	"smart-learning-backend/pkg/interfaces"
	"smart-learning-backend/pkg/models"
	"smart-learning-backend/pkg/utils"
)

type AuthService struct {
	userRepo interfaces.UserRepositoryInterface
}

func NewAuthService(userRepo interfaces.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// 驗證密碼確認
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("passwords do not match")
	}
	
	// 驗證用戶名格式
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(req.Username) {
		return nil, fmt.Errorf("username can only contain letters, numbers and underscores")
	}
	
	// 檢查用戶是否已存在
	exists, err := s.userRepo.CheckUserExists(req.Email, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("user already exists")
	}
	
	// 雜湊密碼
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	
	// 建立用戶
	user := &models.User{
		Email:         req.Email,
		Username:      req.Username,
		PasswordHash:  hashedPassword,
		LearningLevel: 1, // 預設等級
	}
	
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	// 生成 JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	
	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// 根據 email 查找用戶
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	
	// 驗證密碼
	err = utils.VerifyPassword(user.PasswordHash, req.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	
	// 生成 JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	
	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (s *AuthService) GetUserByID(id int) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}