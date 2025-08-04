# Backend API Specification

本文件定義了 Smart Learning 後端 API 的詳細規格，基於前端的實現需求制定。

## 基本信息

- **Base URL**: `http://localhost:8080/api/v1`
- **Authentication**: JWT Bearer Token
- **Content-Type**: `application/json`

## 認證相關 API

### 1. 用戶註冊

**Endpoint**: `POST /auth/register`

**Request Body**:
```json
{
  "email": "user@example.com",
  "username": "testuser",
  "password": "password123",
  "confirm_password": "password123"
}
```

**Request Validation**:
- `email`: 必填，有效的電子郵件格式
- `username`: 必填，2-20字符，只能包含字母、數字和底線 (`^[a-zA-Z0-9_]+$`)
- `password`: 必填，最少8個字符
- `confirm_password`: 必填，必須與 password 一致

**Success Response** (201 Created):
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

**Error Responses**:

400 Bad Request - 驗證失敗:
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "email": ["電子郵件格式不正確"],
    "username": ["用戶名只能包含字母、數字和底線"],
    "password": ["密碼至少需要 8 個字符"]
  }
}
```

409 Conflict - 用戶已存在:
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

### 2. 用戶登入

**Endpoint**: `POST /auth/login`

**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Request Validation**:
- `email`: 必填，有效的電子郵件格式
- `password`: 必填，最少8個字符

**Success Response** (200 OK):
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

**Error Responses**:

400 Bad Request - 驗證失敗:
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "email": ["請輸入有效的電子郵件"],
    "password": ["密碼至少需要 8 個字符"]
  }
}
```

401 Unauthorized - 登入失敗:
```json
{
  "success": false,
  "message": "登入失敗",
  "error": {
    "code": "INVALID_CREDENTIALS",
    "message": "電子郵件或密碼錯誤"
  }
}
```

### 3. 用戶登出

**Endpoint**: `POST /auth/logout`

**Headers**:
```
Authorization: Bearer <jwt_token>
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "登出成功"
}
```

**Error Responses**:

401 Unauthorized - Token 無效:
```json
{
  "success": false,
  "message": "未授權",
  "error": {
    "code": "INVALID_TOKEN",
    "message": "Token 無效或已過期"
  }
}
```

### 4. 取得用戶資料

**Endpoint**: `GET /auth/me`

**Headers**:
```
Authorization: Bearer <jwt_token>
```

**Success Response** (200 OK):
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

## 數據模型

### User Model

```json
{
  "id": "integer, 主鍵，自動遞增",
  "email": "string, 唯一，電子郵件地址",
  "username": "string, 唯一，用戶名",
  "password_hash": "string, 密碼雜湊值（不在API回應中返回）",
  "learning_level": "integer, 學習等級（1-10）",
  "avatar_url": "string, 可選，頭像URL",
  "created_at": "timestamp, 建立時間",
  "updated_at": "timestamp, 更新時間"
}
```

## JWT Token 規格

### Token 結構

```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "user_id": 1,
    "email": "user@example.com",
    "username": "testuser",
    "exp": 1640995200,
    "iat": 1640908800
  }
}
```

### Token 過期時間
- Access Token: 24小時
- 建議後續實作 Refresh Token 機制

## 錯誤處理

### 通用錯誤格式

```json
{
  "success": false,
  "message": "錯誤描述",
  "error": {
    "code": "ERROR_CODE",
    "message": "詳細錯誤信息"
  }
}
```

### 驗證錯誤格式

```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "field_name": ["錯誤信息1", "錯誤信息2"]
  }
}
```

### 常見錯誤代碼

| HTTP Status | Error Code | Description |
|-------------|------------|-------------|
| 400 | VALIDATION_ERROR | 請求參數驗證失敗 |
| 401 | INVALID_CREDENTIALS | 登入憑證無效 |
| 401 | INVALID_TOKEN | JWT Token 無效或過期 |
| 409 | USER_ALREADY_EXISTS | 用戶已存在 |
| 500 | INTERNAL_SERVER_ERROR | 服務器內部錯誤 |

## 資料庫設計建議

### users 表結構

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    learning_level INTEGER DEFAULT 1 CHECK (learning_level >= 1 AND learning_level <= 10),
    avatar_url VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
```

## 安全性考量

1. **密碼處理**:
   - 使用 bcrypt 進行密碼雜湊
   - 成本因子建議設為 12 以上

2. **JWT 安全**:
   - 使用強密鑰簽署 Token
   - 實作 Token 黑名單機制（登出時）
   - 設定合理的過期時間

3. **輸入驗證**:
   - 所有輸入都需要驗證和清理
   - 防止 SQL 注入攻擊
   - 限制請求頻率（Rate Limiting）

4. **CORS 設定**:
   - 正確設定 CORS 策略
   - 只允許信任的來源

## 實作建議

### Go 後端框架選擇
- Gin (已在專案中使用)
- 中介軟體：CORS、JWT、Logging、Rate Limiting

### 必要套件
```go
// JWT 處理
github.com/golang-jwt/jwt/v5

// 密碼雜湊
golang.org/x/crypto/bcrypt

// 資料庫
github.com/lib/pq

// 驗證
github.com/go-playground/validator/v10

// 環境變數
github.com/joho/godotenv
```

### 目錄結構建議
```
backend/
├── cmd/
│   └── main.go
├── pkg/
│   ├── handlers/
│   │   └── auth.go
│   ├── models/
│   │   └── user.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── cors.go
│   ├── services/
│   │   └── auth.go
│   ├── repositories/
│   │   └── user.go
│   └── utils/
│       ├── jwt.go
│       └── password.go
```

## 測試要求

### 必須測試的場景

1. **註冊測試**:
   - 成功註冊
   - 重複電子郵件/用戶名
   - 無效輸入格式
   - 密碼不一致

2. **登入測試**:
   - 成功登入
   - 錯誤憑證
   - 無效輸入格式

3. **JWT 測試**:
   - Token 生成和驗證
   - 過期 Token 處理
   - 無效 Token 處理

4. **資料庫測試**:
   - 用戶創建
   - 用戶查詢
   - 重複約束測試

這份規格基於你的前端實作需求制定，確保後端 API 能完美配合前端的功能需求。