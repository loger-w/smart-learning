# Smart Learning 開發流程指南

## 開發環境設定

### 1. 開發工具建議

#### IDE 配置
- **VS Code** (推薦)
  - Extensions: Go, TypeScript, Tailwind CSS IntelliSense, Prettier, ESLint
  - 配置檔案：`.vscode/settings.json`

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true,
    "source.organizeImports": true
  },
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint",
  "typescript.preferences.importModuleSpecifier": "relative"
}
```

#### Git 配置
```bash
# 配置 Git hooks
# 安裝 husky 和 lint-staged
npm install --save-dev husky lint-staged

# 設定 pre-commit hook
npx husky add .husky/pre-commit "npx lint-staged"
```

### 2. 開發環境啟動

#### 前端開發
```bash
cd frontend
npm install
npm run dev
```

#### 後端開發
```bash
cd backend
go mod tidy
go run cmd/main.go
```

#### 資料庫遷移
```bash
# 建立新遷移
migrate create -ext sql -dir migrations -seq create_users_table

# 執行遷移
migrate -path migrations -database "postgresql://localhost/smart_learning_db?sslmode=disable" up

# 回滾遷移
migrate -path migrations -database "postgresql://localhost/smart_learning_db?sslmode=disable" down 1
```

## Git 倉庫管理策略

### 🏆 **推薦：Monorepo 結構**

採用統一的 `smart-learning` 倉庫管理前後端程式碼：

```
smart-learning/                 # 主倉庫
├── frontend/                   # React 前端專案
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.ts
├── backend/                    # Go 後端專案
│   ├── cmd/
│   ├── internal/
│   ├── go.mod
│   └── Dockerfile
├── docs/                       # 專案文檔
├── .github/                    # GitHub Actions
│   ├── workflows/
│   │   ├── frontend.yml
│   │   └── backend.yml
├── .gitignore                  # 統一 Git 忽略
├── README.md                   # 主要說明文檔
└── package.json                # 根目錄腳本配置
```

### Monorepo 優勢
- ✅ **統一版本管理** - 前後端同步發布
- ✅ **共享文檔配置** - 避免重複維護
- ✅ **簡化 CI/CD** - 單一流程管理
- ✅ **更好協作** - 全棧開發者友好
- ✅ **依賴關係清晰** - API 變更同步更新

### 分支策略

採用 **Git Flow** 分支模型：

```
main (生產)
├── develop (開發主分支)
│   ├── feature/user-authentication
│   ├── feature/flashcard-learning
│   └── feature/ai-integration
├── release/v1.0.0 (發布分支)
└── hotfix/critical-bug-fix (緊急修復)
```

### 分支命名規範

- **功能分支**: `feature/功能名稱`
  - `feature/user-login`
  - `feature/word-list-management`
  
- **修復分支**: `bugfix/問題描述`
  - `bugfix/login-validation-error`
  
- **緊急修復**: `hotfix/緊急問題`
  - `hotfix/security-vulnerability`
  
- **發布分支**: `release/版本號`
  - `release/v1.0.0`

### 提交訊息規範

使用 **Conventional Commits** 格式：

```
<type>(<scope>): <description>

<body>

<footer>
```

#### 提交類型
- **feat**: 新功能
- **fix**: 修復 bug
- **docs**: 文檔更新
- **style**: 程式碼格式調整
- **refactor**: 重構程式碼
- **test**: 測試相關
- **chore**: 維護性工作

#### 範例
```bash
feat(auth): add Google OAuth login functionality

- Implement Google OAuth 2.0 integration
- Add user profile creation from Google data
- Update login flow to support social authentication

Closes #123
```

### 完整開發流程

#### 1. 開始新功能開發
```bash
# 從 develop 分支建立功能分支
git checkout develop
git pull origin develop
git checkout -b feature/user-authentication

# 開發過程中定期提交
git add .
git commit -m "feat(auth): implement user registration API"

# 推送到遠端
git push -u origin feature/user-authentication
```

#### 2. 程式碼審查流程
```bash
# 建立 Pull Request
gh pr create --title "feat: Add user authentication system" \
             --body "Implements user registration, login, and JWT authentication"

# 程式碼審查通過後合併到 develop
git checkout develop
git pull origin develop
git merge --no-ff feature/user-authentication
git push origin develop

# 刪除功能分支
git branch -d feature/user-authentication
git push origin --delete feature/user-authentication
```

## 程式碼品質標準

### 1. TypeScript 編碼規範

#### 命名規範
```typescript
// 組件名稱：PascalCase
const UserProfile: React.FC = () => { }

// 變數和函數：camelCase
const userName = 'john_doe'
const getUserData = () => { }

// 常數：UPPER_SNAKE_CASE
const API_BASE_URL = 'https://api.example.com'

// 介面和類型：PascalCase
interface UserData {
  id: number
  name: string
}

type AuthStatus = 'authenticated' | 'unauthenticated'
```

#### 檔案結構規範
```typescript
// 1. 外部 imports
import React from 'react'
import { useQuery } from '@tanstack/react-query'

// 2. 內部 imports
import { Button } from '@/components/ui/Button'
import { userService } from '@/services/userService'

// 3. 類型定義
interface Props {
  userId: string
}

// 4. 組件實作
const UserProfile: React.FC<Props> = ({ userId }) => {
  // hooks
  const { data, isLoading } = useQuery({
    queryKey: ['user', userId],
    queryFn: () => userService.getUser(userId),
  })

  // 早期返回
  if (isLoading) return <div>Loading...</div>

  // 主要邏輯
  return (
    <div>
      <h1>{data?.name}</h1>
    </div>
  )
}

export default UserProfile
```

### 2. Go 編碼規範

#### 專案結構
```go
// 包名：小寫，簡潔
package userservice

// 介面：行為描述，-er 結尾
type UserRepository interface {
    Create(user *User) error
    GetByID(id int) (*User, error)
}

// 結構體：公開使用 PascalCase
type User struct {
    ID       int    `json:"id" db:"id"`
    Username string `json:"username" db:"username"`
    Email    string `json:"email" db:"email"`
}

// 方法：PascalCase (公開) / camelCase (私有)
func (s *UserService) CreateUser(req CreateUserRequest) (*User, error) {
    if err := s.validateUserRequest(req); err != nil {
        return nil, err
    }
    
    return s.userRepo.Create(&User{
        Username: req.Username,
        Email:    req.Email,
    })
}

func (s *UserService) validateUserRequest(req CreateUserRequest) error {
    // 私有方法實作
}
```

#### 錯誤處理
```go
// 自定義錯誤類型
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

// 錯誤處理模式
func (s *UserService) GetUser(id int) (*User, error) {
    user, err := s.userRepo.GetByID(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrUserNotFound
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}
```

## 測試策略

### 1. 前端測試

#### 單元測試 (Vitest + Testing Library)
```typescript
// components/__tests__/Button.test.tsx
import { render, screen, fireEvent } from '@testing-library/react'
import { describe, it, expect, vi } from 'vitest'
import Button from '../Button'

describe('Button', () => {
  it('renders with correct text', () => {
    render(<Button>Click me</Button>)
    expect(screen.getByRole('button')).toHaveTextContent('Click me')
  })

  it('calls onClick when clicked', () => {
    const handleClick = vi.fn()
    render(<Button onClick={handleClick}>Click me</Button>)
    
    fireEvent.click(screen.getByRole('button'))
    expect(handleClick).toHaveBeenCalledTimes(1)
  })
})
```

#### 整合測試
```typescript
// hooks/__tests__/useAuth.test.tsx
import { renderHook, waitFor } from '@testing-library/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { useAuth } from '../useAuth'

const createWrapper = () => {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } }
  })
  
  return ({ children }: { children: React.ReactNode }) => (
    <QueryClientProvider client={queryClient}>
      {children}
    </QueryClientProvider>
  )
}

describe('useAuth', () => {
  it('should login successfully', async () => {
    const { result } = renderHook(() => useAuth(), {
      wrapper: createWrapper(),
    })

    result.current.login('test@example.com', 'password')

    await waitFor(() => {
      expect(result.current.isAuthenticated).toBe(true)
    })
  })
})
```

### 2. 後端測試

#### 單元測試
```go
// services/user_service_test.go
package services

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) Create(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    service := NewUserService(mockRepo)
    
    req := CreateUserRequest{
        Username: "testuser",
        Email:    "test@example.com",
    }
    
    mockRepo.On("Create", mock.AnythingOfType("*User")).Return(nil)
    
    // Act
    user, err := service.CreateUser(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "testuser", user.Username)
    mockRepo.AssertExpectations(t)
}
```

#### 整合測試
```go
// handlers/user_handler_integration_test.go
func TestUserHandler_Integration(t *testing.T) {
    // 設定測試資料庫
    testDB := setupTestDatabase(t)
    defer cleanupTestDatabase(t, testDB)
    
    // 建立測試伺服器
    router := setupTestRouter(testDB)
    
    // 測試使用者註冊
    req := CreateUserRequest{
        Username: "testuser",
        Email:    "test@example.com",
    }
    
    resp := httptest.NewRecorder()
    reqBody, _ := json.Marshal(req)
    
    httpReq := httptest.NewRequest("POST", "/api/users", strings.NewReader(string(reqBody)))
    httpReq.Header.Set("Content-Type", "application/json")
    
    router.ServeHTTP(resp, httpReq)
    
    assert.Equal(t, http.StatusCreated, resp.Code)
}
```

### 3. 測試執行

#### 前端測試指令
```bash
# 執行所有測試
npm run test

# 監視模式
npm run test:watch

# 覆蓋率報告
npm run test:coverage

# E2E 測試 (如果有設定 Playwright)
npm run test:e2e
```

#### 後端測試指令
```bash
# 執行所有測試
go test ./...

# 詳細輸出
go test -v ./...

# 覆蓋率報告
go test -cover ./...

# 特定包測試
go test ./internal/services/...
```

## CI/CD 流程

### GitHub Actions 配置

#### 前端 CI (.github/workflows/frontend.yml)
```yaml
name: Frontend CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup Node.js
      uses: actions/setup-node@v4
      with:
        node-version: '18'
        cache: 'npm'
        cache-dependency-path: frontend/package-lock.json
    
    - name: Install dependencies
      run: npm ci
      working-directory: frontend
    
    - name: Run linter
      run: npm run lint
      working-directory: frontend
    
    - name: Run type check
      run: npm run type-check
      working-directory: frontend
    
    - name: Run tests
      run: npm run test:coverage
      working-directory: frontend
    
    - name: Build
      run: npm run build
      working-directory: frontend
```

#### 後端 CI (.github/workflows/backend.yml)
```yaml
name: Backend CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: smart_learning_test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    
    - name: Install dependencies
      run: go mod download
      working-directory: backend
    
    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        working-directory: backend
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
      working-directory: backend
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost/smart_learning_test?sslmode=disable
    
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: backend/coverage.out
```

## 程式碼審查指南

### 審查檢查清單

#### 功能性
- [ ] 功能是否符合需求規格
- [ ] 邊界條件處理是否完整
- [ ] 錯誤處理是否適當
- [ ] 測試覆蓋率是否足夠

#### 程式碼品質
- [ ] 程式碼是否易讀易懂
- [ ] 變數和函數命名是否有意義
- [ ] 是否遵循專案編碼規範
- [ ] 是否有不必要的程式碼重複

#### 效能與安全
- [ ] 是否有效能瓶頸
- [ ] 記憶體使用是否合理
- [ ] 是否存在安全漏洞
- [ ] 敏感資訊是否正確處理

#### 架構設計
- [ ] 是否符合系統架構設計
- [ ] 模組耦合度是否合理
- [ ] 是否便於未來擴展
- [ ] 依賴關係是否清晰

### 審查流程

1. **自我審查**：開發者提交前自行檢查
2. **同儕審查**：至少一位其他開發者審查
3. **技術主管審查**：複雜功能需技術主管參與
4. **自動化檢查**：CI/CD 流程自動驗證

## 版本發布流程

### 語義化版本控制

採用 [Semantic Versioning](https://semver.org/) 規範：

- **MAJOR**: 不相容的 API 變更 (1.0.0 → 2.0.0)
- **MINOR**: 向後相容的功能新增 (1.0.0 → 1.1.0)
- **PATCH**: 向後相容的錯誤修復 (1.0.0 → 1.0.1)

### 發布步驟

#### 1. 準備發布
```bash
# 建立發布分支
git checkout develop
git pull origin develop
git checkout -b release/v1.1.0

# 更新版本號
# 前端: package.json
# 後端: version.go 或環境變數

# 更新 CHANGELOG.md
# 執行最終測試
npm run test
go test ./...
```

#### 2. 發布到生產
```bash
# 合併到 main
git checkout main
git merge --no-ff release/v1.1.0

# 建立版本標籤
git tag -a v1.1.0 -m "Release version 1.1.0"

# 推送變更
git push origin main
git push origin v1.1.0

# 合併回 develop
git checkout develop
git merge --no-ff release/v1.1.0
git push origin develop

# 清理發布分支
git branch -d release/v1.1.0
```

#### 3. 部署確認
- 確認生產環境正常運作
- 監控錯誤日誌和效能指標
- 準備回滾計畫（如有需要）

## 故障排除指南

### 常見開發問題

#### 前端問題
```bash
# 依賴問題
rm -rf node_modules package-lock.json
npm install

# 類型錯誤
npm run type-check

# 建構問題
npm run build -- --mode development
```

#### 後端問題
```bash
# 模組問題
go mod tidy
go mod download

# 資料庫連線問題
go run cmd/main.go --check-db

# 測試失敗
go test -v ./... -run TestSpecificFunction
```

### 效能監控

#### 前端效能
- 使用 Chrome DevTools 分析
- 監控 Core Web Vitals
- 檢查 Bundle 大小

#### 後端效能
- API 響應時間監控
- 資料庫查詢優化
- 記憶體使用分析

這份開發流程指南確保團隊成員能夠高效協作，維持程式碼品質，並持續交付高品質的產品功能。