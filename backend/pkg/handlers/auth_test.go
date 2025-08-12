package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"smart-learning-backend/pkg/interfaces"
	"smart-learning-backend/pkg/models"
	"smart-learning-backend/pkg/utils"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// MockAuthService 實現了認證服務的 mock 版本用於測試
type MockAuthService struct {
	users          []models.User
	shouldFailNext string // 指定下一個應該失敗的方法
}

// 確保 MockAuthService 實現 AuthServiceInterface 介面
var _ interfaces.AuthServiceInterface = (*MockAuthService)(nil)

func NewMockAuthService() *MockAuthService {
	return &MockAuthService{
		users: make([]models.User, 0),
	}
}

func (m *MockAuthService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	if m.shouldFailNext == "Register" {
		m.shouldFailNext = ""
		return nil, errors.New("database error")
	}

	// 模擬用戶已存在錯誤
	for _, user := range m.users {
		if user.Email == req.Email || user.Username == req.Username {
			return nil, errors.New("user already exists")
		}
	}

	// 模擬密碼不匹配錯誤
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	// 模擬用戶名格式錯誤
	if strings.Contains(req.Username, "@") {
		return nil, errors.New("username can only contain letters, numbers and underscores")
	}

	// 創建新用戶
	hashedPassword, _ := utils.HashPassword(req.Password)
	user := models.User{
		ID:            len(m.users) + 1,
		Email:         req.Email,
		Username:      req.Username,
		PasswordHash:  hashedPassword,
		LearningLevel: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	m.users = append(m.users, user)

	// 生成 JWT
	token, _ := utils.GenerateJWT(user.ID, user.Email, user.Username)

	return &models.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (m *MockAuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	if m.shouldFailNext == "Login" {
		m.shouldFailNext = ""
		return nil, errors.New("invalid credentials")
	}

	// 查找用戶
	var foundUser *models.User
	for _, user := range m.users {
		if user.Email == req.Email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		return nil, errors.New("invalid credentials")
	}

	// 驗證密碼
	if err := utils.VerifyPassword(foundUser.PasswordHash, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 生成 JWT
	token, _ := utils.GenerateJWT(foundUser.ID, foundUser.Email, foundUser.Username)

	return &models.AuthResponse{
		User:  *foundUser,
		Token: token,
	}, nil
}

func (m *MockAuthService) GetUserByID(id int) (*models.User, error) {
	if m.shouldFailNext == "GetUserByID" {
		m.shouldFailNext = ""
		return nil, errors.New("user not found")
	}

	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockAuthService) SetShouldFailNext(method string) {
	m.shouldFailNext = method
}

func (m *MockAuthService) AddUser(user models.User) {
	m.users = append(m.users, user)
}

// 測試幫助函數
func setupGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func createAuthHandler() *AuthHandler {
	mockService := NewMockAuthService()
	return NewAuthHandler(mockService)
}

func createAuthHandlerWithService(service *MockAuthService) *AuthHandler {
	return NewAuthHandler(service)
}

func createTestUser() models.User {
	hashedPassword, _ := utils.HashPassword("password123")
	return models.User{
		ID:            1,
		Email:         "test@example.com",
		Username:      "testuser",
		PasswordHash:  hashedPassword,
		LearningLevel: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func TestNewAuthHandler(t *testing.T) {
	mockService := NewMockAuthService()
	handler := NewAuthHandler(mockService)

	if handler == nil {
		t.Fatal("NewAuthHandler() returned nil")
	}

	if handler.authService == nil {
		t.Fatal("NewAuthHandler() authService is nil")
	}

	if handler.validator == nil {
		t.Fatal("NewAuthHandler() validator is nil")
	}
}

func TestAuthHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupService   func(*MockAuthService)
		expectedStatus int
		checkResponse  func(t *testing.T, response models.APIResponse)
	}{
		{
			name: "成功註冊",
			requestBody: models.RegisterRequest{
				Email:           "new@example.com",
				Username:        "newuser",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if !response.Success {
					t.Error("Expected success to be true")
				}
				if response.Message != "註冊成功" {
					t.Errorf("Expected message '註冊成功', got '%s'", response.Message)
				}
			},
		},
		{
			name: "無效的 JSON 格式",
			requestBody: `{"email": "invalid json"`,
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
			},
		},
		{
			name: "缺少必填欄位",
			requestBody: models.RegisterRequest{
				Email: "test@example.com",
				// 缺少其他必填欄位
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "驗證失敗" {
					t.Errorf("Expected message '驗證失敗', got '%s'", response.Message)
				}
			},
		},
		{
			name: "用戶已存在",
			requestBody: models.RegisterRequest{
				Email:           "existing@example.com",
				Username:        "existinguser",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			setupService: func(m *MockAuthService) {
				user := createTestUser()
				user.Email = "existing@example.com"
				user.Username = "existinguser"
				m.AddUser(user)
			},
			expectedStatus: http.StatusConflict,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "用戶已存在" {
					t.Errorf("Expected message '用戶已存在', got '%s'", response.Message)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設置
			r := setupGin()
			mockService := NewMockAuthService()
			tt.setupService(mockService)
			handler := createAuthHandlerWithService(mockService)

			r.POST("/register", handler.Register)

			// 準備請求
			var reqBody []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// 執行
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// 驗證
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response models.APIResponse
			err = json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil && w.Code != http.StatusBadRequest {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if w.Code != http.StatusBadRequest {
				tt.checkResponse(t, response)
			}
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupService   func(*MockAuthService)
		expectedStatus int
		checkResponse  func(t *testing.T, response models.APIResponse)
	}{
		{
			name: "成功登入",
			requestBody: models.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupService: func(m *MockAuthService) {
				user := createTestUser()
				m.AddUser(user)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if !response.Success {
					t.Error("Expected success to be true")
				}
				if response.Message != "登入成功" {
					t.Errorf("Expected message '登入成功', got '%s'", response.Message)
				}
			},
		},
		{
			name: "無效的電子郵件格式",
			requestBody: models.LoginRequest{
				Email:    "invalid-email",
				Password: "password123",
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
			},
		},
		{
			name: "用戶不存在",
			requestBody: models.LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "password123",
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusUnauthorized,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "登入失敗" {
					t.Errorf("Expected message '登入失敗', got '%s'", response.Message)
				}
			},
		},
		{
			name: "錯誤的密碼",
			requestBody: models.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			setupService: func(m *MockAuthService) {
				user := createTestUser()
				m.AddUser(user)
			},
			expectedStatus: http.StatusUnauthorized,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設置
			r := setupGin()
			mockService := NewMockAuthService()
			tt.setupService(mockService)
			handler := createAuthHandlerWithService(mockService)

			r.POST("/login", handler.Login)

			// 準備請求
			reqBody, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// 執行
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// 驗證
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response models.APIResponse
			if w.Code != http.StatusBadRequest {
				err = json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}
				tt.checkResponse(t, response)
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	tests := []struct {
		name           string
		setupContext   func(*gin.Context)
		expectedStatus int
		checkResponse  func(t *testing.T, response models.APIResponse)
	}{
		{
			name: "成功登出",
			setupContext: func(c *gin.Context) {
				c.Set("user_id", 1)
				c.Set("username", "testuser")
				c.Set("email", "test@example.com")
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if !response.Success {
					t.Error("Expected success to be true")
				}
				if response.Message != "登出成功" {
					t.Errorf("Expected message '登出成功', got '%s'", response.Message)
				}
				// 檢查回應資料
				data, ok := response.Data.(map[string]interface{})
				if !ok {
					t.Error("Expected data to be a map")
					return
				}
				if data["user_id"] != float64(1) {
					t.Errorf("Expected user_id 1, got %v", data["user_id"])
				}
				if data["username"] != "testuser" {
					t.Errorf("Expected username 'testuser', got %v", data["username"])
				}
			},
		},
		{
			name: "缺少用戶資訊",
			setupContext: func(c *gin.Context) {
				// 不設置任何用戶資訊
			},
			expectedStatus: http.StatusUnauthorized,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "未授權" {
					t.Errorf("Expected message '未授權', got '%s'", response.Message)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設置
			r := setupGin()
			handler := createAuthHandler()

			r.POST("/logout", func(c *gin.Context) {
				tt.setupContext(c)
				handler.Logout(c)
			})

			// 準備請求
			req, err := http.NewRequest("POST", "/logout", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// 執行
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// 驗證
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response models.APIResponse
			err = json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			tt.checkResponse(t, response)
		})
	}
}

func TestAuthHandler_GetMe(t *testing.T) {
	tests := []struct {
		name           string
		setupContext   func(*gin.Context)
		setupService   func(*MockAuthService)
		expectedStatus int
		checkResponse  func(t *testing.T, response models.APIResponse)
	}{
		{
			name: "成功獲取用戶資料",
			setupContext: func(c *gin.Context) {
				c.Set("user_id", 1)
				c.Set("email", "test@example.com")
				c.Set("username", "testuser")
			},
			setupService: func(m *MockAuthService) {
				user := createTestUser()
				m.AddUser(user)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if !response.Success {
					t.Error("Expected success to be true")
				}
			},
		},
		{
			name: "缺少用戶資訊",
			setupContext: func(c *gin.Context) {
				// 不設置 user_id
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusUnauthorized,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "未授權" {
					t.Errorf("Expected message '未授權', got '%s'", response.Message)
				}
			},
		},
		{
			name: "用戶不存在",
			setupContext: func(c *gin.Context) {
				c.Set("user_id", 999)
			},
			setupService:   func(m *MockAuthService) {},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, response models.APIResponse) {
				if response.Success {
					t.Error("Expected success to be false")
				}
				if response.Message != "用戶不存在" {
					t.Errorf("Expected message '用戶不存在', got '%s'", response.Message)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設置
			r := setupGin()
			mockService := NewMockAuthService()
			tt.setupService(mockService)
			handler := createAuthHandlerWithService(mockService)

			r.GET("/me", func(c *gin.Context) {
				tt.setupContext(c)
				handler.GetMe(c)
			})

			// 準備請求
			req, err := http.NewRequest("GET", "/me", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// 執行
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// 驗證
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response models.APIResponse
			err = json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			tt.checkResponse(t, response)
		})
	}
}