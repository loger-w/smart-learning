package utils

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	// 設定測試環境變數
	os.Setenv("JWT_SECRET", "test-secret-key")
	
	// 重新初始化 jwtSecret
	jwtSecret = []byte("test-secret-key")

	userID := 1
	email := "test@example.com"
	username := "testuser"

	token, err := GenerateJWT(userID, email, username)
	
	if err != nil {
		t.Fatalf("GenerateJWT() error = %v", err)
	}
	
	if token == "" {
		t.Fatal("GenerateJWT() returned empty token")
	}

	// 驗證生成的 token
	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("ValidateJWT() error = %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("UserID = %v, want %v", claims.UserID, userID)
	}
	
	if claims.Email != email {
		t.Errorf("Email = %v, want %v", claims.Email, email)
	}
	
	if claims.Username != username {
		t.Errorf("Username = %v, want %v", claims.Username, username)
	}

	// 檢查過期時間是否在合理範圍內（約24小時）
	expectedExpiry := time.Now().Add(24 * time.Hour)
	if claims.ExpiresAt.Time.Before(expectedExpiry.Add(-time.Minute)) {
		t.Error("Token expires too early")
	}
	if claims.ExpiresAt.Time.After(expectedExpiry.Add(time.Minute)) {
		t.Error("Token expires too late")
	}

	// 清理環境變數
	os.Unsetenv("JWT_SECRET")
}

func TestValidateJWT(t *testing.T) {
	// 設定測試環境
	os.Setenv("JWT_SECRET", "test-secret-key")
	jwtSecret = []byte("test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	tests := []struct {
		name      string
		token     string
		wantError bool
	}{
		{
			name:      "有效的 token",
			token:     func() string {
				token, _ := GenerateJWT(1, "test@example.com", "testuser")
				return token
			}(),
			wantError: false,
		},
		{
			name:      "無效的 token",
			token:     "invalid.token.here",
			wantError: true,
		},
		{
			name:      "空的 token",
			token:     "",
			wantError: true,
		},
		{
			name:      "過期的 token",
			token: func() string {
				// 創建一個已過期的 token
				claims := &JWTClaims{
					UserID:   1,
					Email:    "test@example.com",
					Username: "testuser",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // 1小時前過期
						IssuedAt:  jwt.NewNumericDate(time.Now().Add(-25 * time.Hour)), // 25小時前發行
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString(jwtSecret)
				return tokenString
			}(),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateJWT(tt.token)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateJWT() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestExtractTokenFromHeader(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		want       string
		wantError  bool
	}{
		{
			name:       "有效的 Bearer token",
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			want:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			wantError:  false,
		},
		{
			name:       "空的 header",
			authHeader: "",
			want:       "",
			wantError:  true,
		},
		{
			name:       "沒有 Bearer 前綴",
			authHeader: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			want:       "",
			wantError:  true,
		},
		{
			name:       "錯誤的前綴",
			authHeader: "Basic eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			want:       "",
			wantError:  true,
		},
		{
			name:       "只有 Bearer 沒有 token",
			authHeader: "Bearer ",
			want:       "",
			wantError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractTokenFromHeader(tt.authHeader)
			if (err != nil) != tt.wantError {
				t.Errorf("ExtractTokenFromHeader() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractTokenFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}