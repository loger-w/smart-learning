package utils

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "正常密碼",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "空密碼",
			password: "",
			wantErr:  false, // bcrypt 可以處理空字符串
		},
		{
			name:     "長密碼（在72字節限制內）",
			password: strings.Repeat("a", 70),
			wantErr:  false,
		},
		{
			name:     "超長密碼（超過72字節限制）",
			password: strings.Repeat("a", 100),
			wantErr:  true,
		},
		{
			name:     "特殊字符密碼",
			password: "p@ssw0rd!@#$%^&*()",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if !tt.wantErr {
				// 檢查返回的雜湊值不是空的
				if hashed == "" {
					t.Error("HashPassword() returned empty hash")
				}
				
				// 檢查雜湊值與原密碼不同
				if hashed == tt.password {
					t.Error("HashPassword() returned unhashed password")
				}
				
				// 檢查雜湊值可以通過 bcrypt 驗證
				err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(tt.password))
				if err != nil {
					t.Errorf("Generated hash cannot be verified: %v", err)
				}
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	// 先生成一個測試用的雜湊密碼
	testPassword := "testpassword123"
	hashedPassword, err := HashPassword(testPassword)
	if err != nil {
		t.Fatalf("Failed to generate test hash: %v", err)
	}

	tests := []struct {
		name           string
		hashedPassword string
		password       string
		wantErr        bool
	}{
		{
			name:           "正確的密碼",
			hashedPassword: hashedPassword,
			password:       testPassword,
			wantErr:        false,
		},
		{
			name:           "錯誤的密碼",
			hashedPassword: hashedPassword,
			password:       "wrongpassword",
			wantErr:        true,
		},
		{
			name:           "空密碼",
			hashedPassword: hashedPassword,
			password:       "",
			wantErr:        true,
		},
		{
			name:           "無效的雜湊",
			hashedPassword: "invalid-hash",
			password:       testPassword,
			wantErr:        true,
		},
		{
			name:           "空雜湊",
			hashedPassword: "",
			password:       testPassword,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyPassword(tt.hashedPassword, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashPasswordConsistency(t *testing.T) {
	password := "consistency_test"
	
	// 生成兩個不同的雜湊
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)
	
	if err1 != nil || err2 != nil {
		t.Fatalf("HashPassword() failed: err1=%v, err2=%v", err1, err2)
	}
	
	// 雜湊值應該不同（因為 bcrypt 使用隨機 salt）
	if hash1 == hash2 {
		t.Error("HashPassword() should generate different hashes for same password")
	}
	
	// 但兩個雜湊都應該能驗證原密碼
	if err := VerifyPassword(hash1, password); err != nil {
		t.Errorf("First hash verification failed: %v", err)
	}
	
	if err := VerifyPassword(hash2, password); err != nil {
		t.Errorf("Second hash verification failed: %v", err)
	}
}

func TestBcryptCost(t *testing.T) {
	password := "cost_test_password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() failed: %v", err)
	}
	
	// 檢查使用的成本是否符合預期
	cost, err := bcrypt.Cost([]byte(hashedPassword))
	if err != nil {
		t.Fatalf("Failed to get bcrypt cost: %v", err)
	}
	
	expectedCost := bcryptCost // 應該是 12
	if cost != expectedCost {
		t.Errorf("Bcrypt cost = %v, want %v", cost, expectedCost)
	}
}