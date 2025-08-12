package services

import (
	"errors"
	"smart-learning-backend/pkg/models"
	"smart-learning-backend/pkg/utils"
	"testing"
	"time"
)

// MockUserRepository 實現了 UserRepositoryInterface 介面用於測試  
type MockUserRepository struct {
	users          []models.User
	shouldFailNext string // 指定下一個應該失敗的方法
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make([]models.User, 0),
	}
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	if m.shouldFailNext == "CreateUser" {
		m.shouldFailNext = ""
		return errors.New("database error")
	}

	// 檢查用戶是否已存在
	for _, existingUser := range m.users {
		if existingUser.Email == user.Email || existingUser.Username == user.Username {
			return errors.New("user already exists")
		}
	}

	// 模擬數據庫自動填充的字段
	user.ID = len(m.users) + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	m.users = append(m.users, *user)
	return nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	if m.shouldFailNext == "GetUserByEmail" {
		m.shouldFailNext = ""
		return nil, errors.New("database error")
	}

	for _, user := range m.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	if m.shouldFailNext == "GetUserByID" {
		m.shouldFailNext = ""
		return nil, errors.New("database error")
	}

	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) CheckUserExists(email, username string) (bool, error) {
	if m.shouldFailNext == "CheckUserExists" {
		m.shouldFailNext = ""
		return false, errors.New("database error")
	}

	for _, user := range m.users {
		if user.Email == email || user.Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockUserRepository) SetShouldFailNext(method string) {
	m.shouldFailNext = method
}

func TestNewAuthService(t *testing.T) {
	mockRepo := NewMockUserRepository()
	authService := NewAuthService(mockRepo)

	if authService == nil {
		t.Fatal("NewAuthService() returned nil")
	}

	if authService.userRepo == nil {
		t.Fatal("NewAuthService() userRepo is nil")
	}
}

func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name          string
		request       *models.RegisterRequest
		setupMock     func(*MockUserRepository)
		wantError     bool
		errorContains string
	}{
		{
			name: "成功註冊",
			request: &models.RegisterRequest{
				Email:           "test@example.com",
				Username:        "testuser",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				// 不需要特別設置
			},
			wantError: false,
		},
		{
			name: "密碼確認不一致",
			request: &models.RegisterRequest{
				Email:           "test@example.com",
				Username:        "testuser",
				Password:        "password123",
				ConfirmPassword: "different",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
			errorContains: "passwords do not match",
		},
		{
			name: "無效的用戶名格式",
			request: &models.RegisterRequest{
				Email:           "test@example.com",
				Username:        "test@user", // 包含特殊字符
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
			errorContains: "username can only contain",
		},
		{
			name: "用戶已存在",
			request: &models.RegisterRequest{
				Email:           "existing@example.com",
				Username:        "existinguser",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				// 先添加一個用戶
				user := &models.User{
					Email:    "existing@example.com",
					Username: "existinguser",
				}
				m.CreateUser(user)
			},
			wantError: true,
			errorContains: "user already exists",
		},
		{
			name: "CheckUserExists 資料庫錯誤",
			request: &models.RegisterRequest{
				Email:           "test@example.com",
				Username:        "testuser",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				m.SetShouldFailNext("CheckUserExists")
			},
			wantError: true,
			errorContains: "failed to check user existence",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			authService := NewAuthService(mockRepo)

			result, err := authService.Register(tt.request)

			if tt.wantError {
				if err == nil {
					t.Error("Register() expected error but got nil")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Register() error = %v, expected to contain %v", err.Error(), tt.errorContains)
				}
			} else {
				if err != nil {
					t.Errorf("Register() unexpected error = %v", err)
					return
				}
				if result == nil {
					t.Error("Register() returned nil result")
					return
				}
				if result.User.Email != tt.request.Email {
					t.Errorf("Register() user email = %v, want %v", result.User.Email, tt.request.Email)
				}
				if result.User.Username != tt.request.Username {
					t.Errorf("Register() user username = %v, want %v", result.User.Username, tt.request.Username)
				}
				if result.Token == "" {
					t.Error("Register() token is empty")
				}

				// 驗證密碼被正確雜湊
				err = utils.VerifyPassword(result.User.PasswordHash, tt.request.Password)
				if err != nil {
					t.Errorf("Password not correctly hashed: %v", err)
				}
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	// 先創建一個測試用戶
	testUser := &models.User{
		Email:         "test@example.com",
		Username:      "testuser",
		PasswordHash:  func() string { hash, _ := utils.HashPassword("password123"); return hash }(),
		LearningLevel: 1,
	}

	tests := []struct {
		name          string
		request       *models.LoginRequest
		setupMock     func(*MockUserRepository)
		wantError     bool
		errorContains string
	}{
		{
			name: "成功登入",
			request: &models.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				m.CreateUser(testUser)
			},
			wantError: false,
		},
		{
			name: "用戶不存在",
			request: &models.LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
			errorContains: "invalid credentials",
		},
		{
			name: "錯誤的密碼",
			request: &models.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			setupMock: func(m *MockUserRepository) {
				m.CreateUser(testUser)
			},
			wantError: true,
			errorContains: "invalid credentials",
		},
		{
			name: "GetUserByEmail 資料庫錯誤",
			request: &models.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				m.SetShouldFailNext("GetUserByEmail")
			},
			wantError: true,
			errorContains: "invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			authService := NewAuthService(mockRepo)

			result, err := authService.Login(tt.request)

			if tt.wantError {
				if err == nil {
					t.Error("Login() expected error but got nil")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Login() error = %v, expected to contain %v", err.Error(), tt.errorContains)
				}
			} else {
				if err != nil {
					t.Errorf("Login() unexpected error = %v", err)
					return
				}
				if result == nil {
					t.Error("Login() returned nil result")
					return
				}
				if result.User.Email != tt.request.Email {
					t.Errorf("Login() user email = %v, want %v", result.User.Email, tt.request.Email)
				}
				if result.Token == "" {
					t.Error("Login() token is empty")
				}
			}
		})
	}
}

func TestAuthService_GetUserByID(t *testing.T) {
	// 創建測試用戶
	testUser := &models.User{
		Email:         "test@example.com",
		Username:      "testuser",
		LearningLevel: 1,
	}

	tests := []struct {
		name          string
		userID        int
		setupMock     func(*MockUserRepository)
		wantError     bool
		errorContains string
	}{
		{
			name:   "成功獲取用戶",
			userID: 1,
			setupMock: func(m *MockUserRepository) {
				m.CreateUser(testUser)
			},
			wantError: false,
		},
		{
			name:          "用戶不存在",
			userID:        999,
			setupMock:     func(m *MockUserRepository) {},
			wantError:     true,
			errorContains: "user not found",
		},
		{
			name:   "資料庫錯誤",
			userID: 1,
			setupMock: func(m *MockUserRepository) {
				m.SetShouldFailNext("GetUserByID")
			},
			wantError:     true,
			errorContains: "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			authService := NewAuthService(mockRepo)

			result, err := authService.GetUserByID(tt.userID)

			if tt.wantError {
				if err == nil {
					t.Error("GetUserByID() expected error but got nil")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("GetUserByID() error = %v, expected to contain %v", err.Error(), tt.errorContains)
				}
			} else {
				if err != nil {
					t.Errorf("GetUserByID() unexpected error = %v", err)
					return
				}
				if result == nil {
					t.Error("GetUserByID() returned nil result")
					return
				}
				if result.ID != tt.userID {
					t.Errorf("GetUserByID() user ID = %v, want %v", result.ID, tt.userID)
				}
			}
		})
	}
}

// 幫助函數：檢查字符串是否包含子字符串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			s[len(s)-len(substr):] == substr || 
			containsInMiddle(s, substr))))
}

func containsInMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}