# Smart Learning Backend API 文檔

## 概述

Smart Learning 後端 API 提供用戶認證和學習功能的 RESTful 接口。使用 Go + Gin 框架開發，支援 JWT 認證。

**Base URL**: `http://localhost:8080`

## 認證

API 使用 JWT (JSON Web Token) 進行身份驗證。需要認證的端點需要在請求頭中包含：

```
Authorization: Bearer <token>
```

## 通用響應格式

所有 API 響應都遵循統一的格式：

### 成功響應
```json
{
  "success": true,
  "message": "操作成功",
  "data": {
    // 響應數據
  }
}
```

### 錯誤響應
```json
{
  "success": false,
  "message": "錯誤訊息",
  "error": {
    "code": "ERROR_CODE",
    "message": "詳細錯誤訊息"
  }
}
```

### 驗證錯誤響應
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "field_name": ["錯誤訊息1", "錯誤訊息2"]
  }
}
```

## 系統端點

### 健康檢查

檢查 API 服務和資料庫連接狀態。

**端點**: `GET /health`

**請求參數**: 無

**響應範例**:
```json
{
  "status": "ok",
  "message": "Smart Learning API is running",
  "db_stats": {
    "max_open_connections": 25,
    "open_connections": 1,
    "in_use": 0,
    "idle": 1
  }
}
```

### Ping 測試

簡單的 API 連通性測試。

**端點**: `GET /api/v1/ping`

**請求參數**: 無

**響應範例**:
```json
{
  "message": "pong"
}
```

## 認證端點

### 用戶註冊

創建新的用戶帳戶。

**端點**: `POST /api/v1/auth/register`

**請求體**:
```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "password123",
  "confirm_password": "password123"
}
```

**請求欄位驗證**:
- `email`: 必填，有效的電子郵件格式
- `username`: 必填，2-20 字符，只能包含字母、數字和底線
- `password`: 必填，至少 8 個字符
- `confirm_password`: 必填，必須與 password 相同

**成功響應** (201 Created):
```json
{
  "success": true,
  "message": "註冊成功",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "username",
      "learning_level": 1,
      "avatar_url": null,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**錯誤響應**:

用戶已存在 (409 Conflict):
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

密碼確認不一致 (400 Bad Request):
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "confirm_password": ["密碼確認不一致"]
  }
}
```

### 用戶登入

使用電子郵件和密碼登入。

**端點**: `POST /api/v1/auth/login`

**請求體**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**請求欄位驗證**:
- `email`: 必填，有效的電子郵件格式
- `password`: 必填，至少 8 個字符

**成功響應** (200 OK):
```json
{
  "success": true,
  "message": "登入成功",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "username",
      "learning_level": 1,
      "avatar_url": null,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**錯誤響應**:

認證失敗 (401 Unauthorized):
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

### 用戶登出

登出當前用戶會話。

**端點**: `POST /api/v1/auth/logout`

**認證**: 需要 JWT Token

**請求參數**: 無

**成功響應** (200 OK):
```json
{
  "success": true,
  "message": "登出成功"
}
```

**錯誤響應**:

未授權 (401 Unauthorized):
```json
{
  "success": false,
  "message": "未授權",
  "error": {
    "code": "MISSING_TOKEN",
    "message": "Authorization header is required"
  }
}
```

### 獲取當前用戶資料

獲取當前登入用戶的詳細資料。

**端點**: `GET /api/v1/auth/me`

**認證**: 需要 JWT Token

**請求參數**: 無

**成功響應** (200 OK):
```json
{
  "success": true,
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "username": "username",
      "learning_level": 1,
      "avatar_url": null,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    }
  }
}
```

**錯誤響應**:

未授權 (401 Unauthorized):
```json
{
  "success": false,
  "message": "未授權",
  "error": {
    "code": "UNAUTHORIZED",
    "message": "無法獲取用戶資訊"
  }
}
```

用戶不存在 (404 Not Found):
```json
{
  "success": false,
  "message": "用戶不存在",
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "用戶不存在"
  }
}
```

## 資料模型

### User 用戶模型
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "username",
  "learning_level": 1,
  "avatar_url": "https://example.com/avatar.jpg",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

### AuthResponse 認證響應模型
```json
{
  "user": {
    // User 模型
  },
  "token": "JWT Token 字符串"
}
```

## 錯誤代碼

| 錯誤代碼 | HTTP 狀態碼 | 描述 |
|---------|------------|------|
| USER_ALREADY_EXISTS | 409 | 用戶已存在（電子郵件或用戶名重複） |
| INVALID_CREDENTIALS | 401 | 登入憑證無效 |
| MISSING_TOKEN | 401 | 缺少 Authorization 標頭 |
| INVALID_TOKEN_FORMAT | 401 | Authorization 標頭格式無效 |
| INVALID_TOKEN | 401 | JWT Token 無效或已過期 |
| UNAUTHORIZED | 401 | 未授權存取 |
| USER_NOT_FOUND | 404 | 用戶不存在 |
| INTERNAL_SERVER_ERROR | 500 | 伺服器內部錯誤 |

## 使用範例

### 註冊新用戶
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "confirm_password": "password123"
  }'
```

### 用戶登入
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 獲取用戶資料
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 用戶登出
```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 環境配置

### 環境變數
- `PORT`: 服務端口（預設: 8080）
- `GIN_MODE`: Gin 運行模式（development/production）
- `DATABASE_URL`: PostgreSQL 資料庫連接字符串
- `JWT_SECRET`: JWT 簽名密鑰
- `TRUSTED_PROXIES`: 信任的代理服務器 IP 列表

### 開發環境啟動
```bash
cd backend
go mod tidy
go run cmd/main.go
```

## 版本資訊

- **版本**: v1
- **Go 版本**: 1.24.5
- **框架**: Gin 1.10.1
- **資料庫**: PostgreSQL
- **認證**: JWT

## 聯絡資訊

如有問題或建議，請聯絡開發團隊。