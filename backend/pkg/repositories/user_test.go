package repositories

import (
	"database/sql"
	"errors"
	"smart-learning-backend/pkg/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewUserRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	if repo == nil {
		t.Fatal("NewUserRepository() returned nil")
	}
	if repo.db != db {
		t.Fatal("NewUserRepository() db field not set correctly")
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	testUser := &models.User{
		Email:         "test@example.com",
		Username:      "testuser",
		PasswordHash:  "hashedpassword",
		LearningLevel: 1,
		AvatarURL:     nil,
	}

	expectedTime := time.Now()

	tests := []struct {
		name      string
		user      *models.User
		mockSetup func()
		wantError bool
		errorMsg  string
	}{
		{
			name: "成功創建用戶",
			user: testUser,
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
					AddRow(1, expectedTime, expectedTime)
				
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs(testUser.Email, testUser.Username, testUser.PasswordHash, 
						testUser.LearningLevel, testUser.AvatarURL).
					WillReturnRows(rows)
			},
			wantError: false,
		},
		{
			name: "數據庫約束錯誤",
			user: testUser,
			mockSetup: func() {
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs(testUser.Email, testUser.Username, testUser.PasswordHash, 
						testUser.LearningLevel, testUser.AvatarURL).
					WillReturnError(errors.New("constraint violation"))
			},
			wantError: true,
			errorMsg:  "failed to create user",
		},
		{
			name: "一般數據庫錯誤",
			user: testUser,
			mockSetup: func() {
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs(testUser.Email, testUser.Username, testUser.PasswordHash, 
						testUser.LearningLevel, testUser.AvatarURL).
					WillReturnError(errors.New("database connection failed"))
			},
			wantError: true,
			errorMsg:  "failed to create user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			
			// 創建用戶的副本以避免修改原始測試數據
			userCopy := *tt.user
			
			err := repo.CreateUser(&userCopy)
			
			if tt.wantError {
				if err == nil {
					t.Error("CreateUser() expected error but got nil")
					return
				}
				if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("CreateUser() error = %v, expected to contain %v", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("CreateUser() unexpected error = %v", err)
					return
				}
				if userCopy.ID == 0 {
					t.Error("CreateUser() did not set user ID")
				}
				if userCopy.CreatedAt.IsZero() {
					t.Error("CreateUser() did not set CreatedAt")
				}
				if userCopy.UpdatedAt.IsZero() {
					t.Error("CreateUser() did not set UpdatedAt")
				}
			}
			
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %v", err)
			}
		})
	}
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	expectedUser := &models.User{
		ID:            1,
		Email:         "test@example.com",
		Username:      "testuser",
		PasswordHash:  "hashedpassword",
		LearningLevel: 1,
		AvatarURL:     nil,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	tests := []struct {
		name      string
		email     string
		mockSetup func()
		wantUser  *models.User
		wantError bool
		errorMsg  string
	}{
		{
			name:  "成功獲取用戶",
			email: "test@example.com",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "email", "username", "password_hash", 
					"learning_level", "avatar_url", "created_at", "updated_at"}).
					AddRow(expectedUser.ID, expectedUser.Email, expectedUser.Username, 
						expectedUser.PasswordHash, expectedUser.LearningLevel, expectedUser.AvatarURL,
						expectedUser.CreatedAt, expectedUser.UpdatedAt)
				
				mock.ExpectQuery(`SELECT (.+) FROM users WHERE email`).
					WithArgs("test@example.com").
					WillReturnRows(rows)
			},
			wantUser:  expectedUser,
			wantError: false,
		},
		{
			name:  "用戶不存在",
			email: "nonexistent@example.com",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT (.+) FROM users WHERE email`).
					WithArgs("nonexistent@example.com").
					WillReturnError(sql.ErrNoRows)
			},
			wantUser:  nil,
			wantError: true,
			errorMsg:  "user not found",
		},
		{
			name:  "數據庫錯誤",
			email: "test@example.com",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT (.+) FROM users WHERE email`).
					WithArgs("test@example.com").
					WillReturnError(errors.New("database connection failed"))
			},
			wantUser:  nil,
			wantError: true,
			errorMsg:  "failed to get user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			
			user, err := repo.GetUserByEmail(tt.email)
			
			if tt.wantError {
				if err == nil {
					t.Error("GetUserByEmail() expected error but got nil")
					return
				}
				if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("GetUserByEmail() error = %v, expected to contain %v", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("GetUserByEmail() unexpected error = %v", err)
					return
				}
				if user == nil {
					t.Error("GetUserByEmail() returned nil user")
					return
				}
				if user.Email != tt.wantUser.Email {
					t.Errorf("GetUserByEmail() email = %v, want %v", user.Email, tt.wantUser.Email)
				}
				if user.Username != tt.wantUser.Username {
					t.Errorf("GetUserByEmail() username = %v, want %v", user.Username, tt.wantUser.Username)
				}
			}
			
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %v", err)
			}
		})
	}
}

func TestUserRepository_CheckUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	tests := []struct {
		name      string
		email     string
		username  string
		mockSetup func()
		wantExists bool
		wantError  bool
		errorMsg   string
	}{
		{
			name:     "用戶存在",
			email:    "existing@example.com",
			username: "existinguser",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
				mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users`).
					WithArgs("existing@example.com", "existinguser").
					WillReturnRows(rows)
			},
			wantExists: true,
			wantError:  false,
		},
		{
			name:     "用戶不存在",
			email:    "new@example.com",
			username: "newuser",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
				mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users`).
					WithArgs("new@example.com", "newuser").
					WillReturnRows(rows)
			},
			wantExists: false,
			wantError:  false,
		},
		{
			name:     "數據庫錯誤",
			email:    "test@example.com",
			username: "testuser",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users`).
					WithArgs("test@example.com", "testuser").
					WillReturnError(errors.New("database connection failed"))
			},
			wantExists: false,
			wantError:  true,
			errorMsg:   "failed to check user existence",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			
			exists, err := repo.CheckUserExists(tt.email, tt.username)
			
			if tt.wantError {
				if err == nil {
					t.Error("CheckUserExists() expected error but got nil")
					return
				}
				if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("CheckUserExists() error = %v, expected to contain %v", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("CheckUserExists() unexpected error = %v", err)
					return
				}
				if exists != tt.wantExists {
					t.Errorf("CheckUserExists() = %v, want %v", exists, tt.wantExists)
				}
			}
			
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %v", err)
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