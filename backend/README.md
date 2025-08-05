# Smart Learning Backend

這是 Smart Learning 平台的後端 API 服務，使用 Go 和 Gin 框架開發。

## 功能特性

- 用戶註冊和登入
- JWT 身份驗證
- PostgreSQL 資料庫支援
- RESTful API 設計
- CORS 支援
- 密碼雜湊加密
- 資料驗證

## 快速開始

### 1. 安裝依賴

```bash
go mod tidy
```

### 2. 設定環境變數

複製環境變數範例檔案：

```bash
cp .env.example .env
```

編輯 `.env` 檔案並設定你的資料庫連接資訊：

```env
# 資料庫配置
DATABASE_URL=postgres://username:password@localhost:5432/smart_learning?sslmode=disable

# JWT 配置
JWT_SECRET=your-very-secure-jwt-secret-key
```

### 3. 設定資料庫

執行資料庫遷移：

```sql
-- 連接到你的 PostgreSQL 資料庫並執行：
-- migrations/001_create_users_table.sql
```

### 4. 啟動服務

```bash
# 開發模式
make dev

# 或直接執行
go run cmd/main.go
```

伺服器將在 `http://localhost:8080` 啟動。

## API 端點

### 認證 API

所有認證相關的 API 都在 `/api/v1/auth` 路徑下：

#### 用戶註冊
- **POST** `/api/v1/auth/register`
- **Content-Type**: `application/json`

請求體：
```json
{
  "email": "user@example.com",
  "username": "testuser",
  "password": "password123",
  "confirm_password": "password123"
}
```

成功回應 (201 Created)：
```json
{
  "success": true,
  "message": "註冊成功",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "testuser",
      "learning_level": 1,
      "avatar_url": null,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 用戶登入
- **POST** `/api/v1/auth/login`
- **Content-Type**: `application/json`

請求體：
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

成功回應 (200 OK)：
```json
{
  "success": true,
  "message": "登入成功",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "testuser",
      "learning_level": 3,
      "avatar_url": "https://example.com/avatar.jpg",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 用戶登出
- **POST** `/api/v1/auth/logout`
- **Authorization**: `Bearer <jwt_token>`

成功回應 (200 OK)：
```json
{
  "success": true,
  "message": "登出成功"
}
```

#### 取得用戶資料
- **GET** `/api/v1/auth/me`
- **Authorization**: `Bearer <jwt_token>`

成功回應 (200 OK)：
```json
{
  "success": true,
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "testuser",
      "learning_level": 3,
      "avatar_url": "https://example.com/avatar.jpg",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

### 其他端點

#### 健康檢查
- **GET** `/health`

```json
{
  "status": "ok",
  "message": "Smart Learning API is running",
  "db_stats": {
    "max_open_connections": 30,
    "open_connections": 1,
    "in_use": 0,
    "idle": 1
  }
}
```

#### 測試端點
- **GET** `/api/v1/ping`

```json
{
  "message": "pong"
}
```

## 錯誤處理

所有 API 錯誤都會回傳統一的錯誤格式：

### 驗證錯誤 (400 Bad Request)
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "email": ["電子郵件格式不正確"],
    "password": ["密碼至少需要 8 個字符"]
  }
}
```

### 業務邏輯錯誤 (401, 409 等)
```json
{
  "success": false,
  "message": "用戶已存在",
  "error": {
    "code": "USER_ALREADY_EXISTS",
    "message": "電子郵件或用戶名已被使用"
  }
}
```

## 開發指令

### 使用 Makefile

```bash
# 安裝依賴
make deps

# 啟動開發服務器 (使用 air 熱重載)
make dev

# 運行服務器
make run

# 運行測試
make test

# 運行測試並生成覆蓋率報告
make coverage

# 建構二進制檔案
make build

# 清理建構檔案
make clean
```

### 直接使用 Go 指令

```bash
# 安裝依賴
go mod tidy

# 啟動服務器
go run cmd/main.go

# 運行測試
go test ./...

# 建構
go build -o bin/smart-learning-backend cmd/main.go
```

## 專案結構

```
backend/
├── cmd/                    # 應用程式入口點
│   └── main.go
├── pkg/                    # 共享套件
│   ├── database/          # 資料庫連接
│   ├── handlers/          # HTTP 處理器
│   ├── middleware/        # 中介軟體
│   ├── models/           # 資料模型
│   ├── repositories/     # 資料存取層
│   ├── services/         # 業務邏輯層
│   └── utils/            # 工具函數
├── migrations/           # 資料庫遷移檔案
├── .env.example         # 環境變數範例
├── Dockerfile           # Docker 配置
├── Makefile            # 建構指令
├── go.mod              # Go 模組定義
└── go.sum              # 依賴鎖定檔案
```

## 技術棧

- **語言**: Go 1.24.5
- **框架**: Gin Web Framework
- **資料庫**: PostgreSQL
- **認證**: JWT (JSON Web Tokens)
- **密碼加密**: bcrypt
- **環境變數**: godotenv

## 部署

### Docker 部署

```bash
# 建構 Docker 映像
docker build -t smart-learning-backend .

# 運行容器
docker run -p 8080:8080 --env-file .env smart-learning-backend
```

### 環境變數

確保在生產環境中設定以下環境變數：

- `DATABASE_URL`: PostgreSQL 連接字串
- `JWT_SECRET`: JWT 簽署密鑰 (請使用強密鑰)
- `GIN_MODE`: 設為 `release`
- `PORT`: 伺服器端口 (預設 8080)

## 安全性

- 所有密碼使用 bcrypt 雜湊加密
- JWT Token 有效期為 24 小時
- 輸入驗證和清理
- CORS 配置
- SQL 注入防護

## 授權

MIT License