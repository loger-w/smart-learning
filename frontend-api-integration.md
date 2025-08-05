# Smart Learning 前端 API 整合指南

## 專案概述

Smart Learning 是一個智能英語學習平台，後端使用 Go + Gin 框架，前端使用 React + TypeScript + TanStack Query。本文件詳細說明後端 API 規格及前端串接的完整實作策略。

## 後端 API 規格

### Base URL
```
開發環境: http://localhost:8080
生產環境: https://api.smart-learning.com
```

### 認證機制
API 使用 JWT (JSON Web Token) 進行身份驗證。需要認證的端點需要在請求頭中包含：
```
Authorization: Bearer <token>
```

### 統一響應格式

#### 成功響應
```json
{
  "success": true,
  "message": "操作成功",
  "data": {
    // 響應數據
  }
}
```

#### 錯誤響應
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

#### 驗證錯誤響應
```json
{
  "success": false,
  "message": "驗證失敗",
  "errors": {
    "field_name": ["錯誤訊息1", "錯誤訊息2"]
  }
}
```

## API 端點規格

### 系統端點

#### 健康檢查
**端點**: `GET /health`

檢查 API 服務和資料庫連接狀態。

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

#### Ping 測試
**端點**: `GET /api/v1/ping`

簡單的 API 連通性測試。

**響應範例**:
```json
{
  "message": "pong"
}
```

### 認證端點

#### 用戶註冊
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

#### 用戶登入
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

#### 用戶登出
**端點**: `POST /api/v1/auth/logout`

**認證**: 需要 JWT Token

**成功響應** (200 OK):
```json
{
  "success": true,
  "message": "登出成功"
}
```

#### 獲取當前用戶資料
**端點**: `GET /api/v1/auth/me`

**認證**: 需要 JWT Token

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

## 錯誤代碼表

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

## 前端實作策略

### 型別定義

#### 統一 API 回應型別
```typescript
// types/api.ts
export interface APIResponse<T = any> {
  success: boolean
  message?: string
  data?: T
  error?: APIError
  errors?: Record<string, string[]>
}

export interface APIError {
  code: string
  message: string
}
```

#### 認證相關型別
```typescript
// types/auth.ts
export interface User {
  id: number
  email: string
  username: string
  learning_level: number
  avatar_url?: string
  created_at: string
  updated_at: string
}

export interface AuthResponse {
  user: User
  token: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  username: string
  password: string
  confirm_password: string
}

export interface LoginResponse extends APIResponse<AuthResponse> {}
export interface RegisterResponse extends APIResponse<AuthResponse> {}
export interface MeResponse extends APIResponse<{ user: User }> {}
```

### 統一 API 客戶端

```typescript
// services/apiClient.ts
import axios, { AxiosInstance, AxiosError } from 'axios'
import type { APIResponse } from '@/types/api'

class APIClient {
  private client: AxiosInstance
  
  constructor(baseURL: string) {
    this.client = axios.create({
      baseURL,
      headers: { 'Content-Type': 'application/json' }
    })
    
    this.setupInterceptors()
  }
  
  private setupInterceptors() {
    // 請求攔截器：自動添加 JWT token
    this.client.interceptors.request.use(config => {
      const token = this.getStoredToken()
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      return config
    })
    
    // 回應攔截器：統一錯誤處理
    this.client.interceptors.response.use(
      response => response,
      (error: AxiosError) => this.handleResponseError(error)
    )
  }
  
  private getStoredToken(): string | null {
    // 從 localStorage 或 store 獲取 token
    return localStorage.getItem('auth-token')
  }
  
  private handleResponseError(error: AxiosError) {
    if (error.response?.status === 401) {
      this.handleUnauthorized()
    }
    return Promise.reject(error)
  }
  
  private handleUnauthorized() {
    // 清除 token，重定向到登入頁
    localStorage.removeItem('auth-token')
    window.location.href = '/login'
  }
  
  // HTTP 方法封裝
  get<T = any>(url: string, config = {}) {
    return this.client.get<T>(url, config)
  }
  
  post<T = any>(url: string, data = {}, config = {}) {
    return this.client.post<T>(url, data, config)
  }
  
  put<T = any>(url: string, data = {}, config = {}) {
    return this.client.put<T>(url, data, config)
  }
  
  delete<T = any>(url: string, config = {}) {
    return this.client.delete<T>(url, config)
  }
}

export const apiClient = new APIClient(
  import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1'
)
```

### 認證服務重構

```typescript
// services/authService.ts
import { apiClient } from './apiClient'
import type { 
  LoginRequest, 
  RegisterRequest, 
  LoginResponse, 
  RegisterResponse, 
  MeResponse,
  AuthResponse,
  User 
} from '@/types/auth'
import { AuthError, ValidationError } from '@/utils/errors'

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await apiClient.post<LoginResponse>('/auth/login', credentials)
    
    if (!response.data.success) {
      throw new AuthError(response.data.error?.message || '登入失敗')
    }
    
    return response.data.data!
  },
  
  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await apiClient.post<RegisterResponse>('/auth/register', data)
    
    if (!response.data.success) {
      // 處理驗證錯誤
      if (response.data.errors) {
        throw new ValidationError(response.data.errors)
      }
      throw new AuthError(response.data.error?.message || '註冊失敗')
    }
    
    return response.data.data!
  },
  
  async logout(): Promise<void> {
    await apiClient.post('/auth/logout')
  },
  
  async getMe(): Promise<User> {
    const response = await apiClient.get<MeResponse>('/auth/me')
    
    if (!response.data.success) {
      throw new AuthError(response.data.error?.message || '獲取用戶資料失敗')
    }
    
    return response.data.data!.user
  }
}
```

### TanStack Query 整合

```typescript
// hooks/useAuth.ts
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { authService } from '@/services/authService'
import { useAuthStore } from '@/stores/authStore'
import type { LoginRequest, RegisterRequest } from '@/types/auth'

export const useLogin = () => {
  const { setAuth } = useAuthStore()
  
  return useMutation({
    mutationFn: authService.login,
    onSuccess: (data) => {
      setAuth(data.user, data.token)
    },
    onError: (error) => {
      console.error('Login failed:', error)
    }
  })
}

export const useRegister = () => {
  const { setAuth } = useAuthStore()
  
  return useMutation({
    mutationFn: authService.register,
    onSuccess: (data) => {
      setAuth(data.user, data.token)
    }
  })
}

export const useMe = () => {
  const { token } = useAuthStore()
  
  return useQuery({
    queryKey: ['auth', 'me'],
    queryFn: authService.getMe,
    enabled: !!token,
    staleTime: 5 * 60 * 1000, // 5 分鐘
    retry: (failureCount, error: any) => {
      // 如果是 401 錯誤，不重試
      if (error?.response?.status === 401) {
        return false
      }
      return failureCount < 3
    }
  })
}

export const useLogout = () => {
  const { logout } = useAuthStore()
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: authService.logout,
    onSuccess: () => {
      logout()
      queryClient.clear()
    }
  })
}
```

### 錯誤處理機制

```typescript
// utils/errors.ts
export class APIError extends Error {
  constructor(
    message: string,
    public code?: string,
    public status?: number
  ) {
    super(message)
    this.name = 'APIError'
  }
}

export class AuthError extends APIError {
  constructor(message: string, code?: string) {
    super(message, code, 401)
    this.name = 'AuthError'
  }
}

export class ValidationError extends APIError {
  constructor(public fields: Record<string, string[]>) {
    super('驗證失敗')
    this.name = 'ValidationError'
  }
  
  getFieldErrors(fieldName: string): string[] {
    return this.fields[fieldName] || []
  }
  
  hasFieldError(fieldName: string): boolean {
    return Array.isArray(this.fields[fieldName]) && this.fields[fieldName].length > 0
  }
}
```

### 錯誤邊界組件

```typescript
// components/ErrorBoundary.tsx
import React from 'react'
import { ValidationError, AuthError } from '@/utils/errors'

interface Props {
  children: React.ReactNode
  fallback?: React.ComponentType<{ error: Error }>
}

export const ErrorBoundary: React.FC<Props> = ({ children, fallback: Fallback }) => {
  const [error, setError] = React.useState<Error | null>(null)
  
  React.useEffect(() => {
    const handleError = (event: ErrorEvent) => {
      setError(new Error(event.message))
    }
    
    const handlePromiseRejection = (event: PromiseRejectionEvent) => {
      setError(new Error(event.reason))
    }
    
    window.addEventListener('error', handleError)
    window.addEventListener('unhandledrejection', handlePromiseRejection)
    
    return () => {
      window.removeEventListener('error', handleError)
      window.removeEventListener('unhandledrejection', handlePromiseRejection)
    }
  }, [])
  
  if (error) {
    if (Fallback) {
      return <Fallback error={error} />
    }
    
    return (
      <div className="p-4 border border-red-300 rounded-md bg-red-50">
        <h2 className="text-lg font-semibold text-red-800">發生錯誤</h2>
        <p className="mt-2 text-red-600">{error.message}</p>
      </div>
    )
  }
  
  return <>{children}</>
}
```

## 環境配置

### 環境變數
```env
# .env.local
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_APP_NAME=Smart Learning
```

### API 配置
```typescript
// config/api.ts
export const API_CONFIG = {
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
  timeout: 10000,
  retries: 3
}
```

## 檔案結構

```
frontend/src/
├── services/
│   ├── apiClient.ts         # 統一 API 客戶端
│   ├── authService.ts       # 認證服務
│   └── index.ts            # 服務入口點
├── hooks/
│   ├── useAuth.ts          # 認證相關鉤子
│   ├── useAPI.ts           # 通用 API 鉤子
│   └── index.ts            # 鉤子入口點
├── types/
│   ├── api.ts              # API 相關型別
│   ├── auth.ts             # 認證型別
│   └── index.ts            # 型別入口點
├── utils/
│   ├── errors.ts           # 錯誤處理工具
│   └── api.ts              # API 工具函數
├── config/
│   └── api.ts              # API 配置
└── components/
    ├── ErrorBoundary.tsx    # 錯誤邊界
    └── LoadingSpinner.tsx   # 載入組件
```

## 使用範例

### 登入範例
```typescript
// pages/LoginPage.tsx
import { useLogin } from '@/hooks/useAuth'
import { useState } from 'react'

export const LoginPage = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' })
  const loginMutation = useLogin()
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    loginMutation.mutate(credentials)
  }
  
  return (
    <form onSubmit={handleSubmit}>
      <input
        type="email"
        value={credentials.email}
        onChange={(e) => setCredentials(prev => ({ ...prev, email: e.target.value }))}
        placeholder="電子郵件"
      />
      <input
        type="password"
        value={credentials.password}
        onChange={(e) => setCredentials(prev => ({ ...prev, password: e.target.value }))}
        placeholder="密碼"
      />
      <button type="submit" disabled={loginMutation.isPending}>
        {loginMutation.isPending ? '登入中...' : '登入'}
      </button>
      {loginMutation.error && (
        <p className="text-red-600">{loginMutation.error.message}</p>
      )}
    </form>
  )
}
```

## 測試策略

### 單元測試
```typescript
// services/__tests__/authService.test.ts
import { describe, it, expect, vi } from 'vitest'
import { authService } from '../authService'

describe('authService', () => {
  it('should login successfully', async () => {
    // Mock API response
    // Test login functionality
  })
  
  it('should handle login errors', async () => {
    // Test error handling
  })
})
```

### 整合測試
```typescript
// hooks/__tests__/useAuth.test.ts
import { renderHook, act } from '@testing-library/react'
import { useLogin } from '../useAuth'

describe('useLogin', () => {
  it('should handle successful login', async () => {
    // Test hook behavior
  })
})
```

## 部署注意事項

### 生產環境配置
- 設定正確的 API 基礎 URL
- 啟用 HTTPS
- 配置 CSP (Content Security Policy)
- 實作錯誤日誌收集

### 安全性考慮
- JWT token 安全存儲 (考慮使用 httpOnly cookies)
- 實作 token 刷新機制
- API 請求頻率限制
- 輸入驗證和清理

## 總結

本文件提供了完整的前端 API 整合解決方案，包括：

1. **完整的 API 規格文件**：詳細的端點說明和響應格式
2. **統一的錯誤處理**：標準化的錯誤型別和處理機制
3. **現代化狀態管理**：TanStack Query + Zustand 整合
4. **型別安全**：完整的 TypeScript 型別定義
5. **可維護的架構**：清晰的檔案結構和職責分離

實作完成後，將提供穩定、高效、可維護的 API 整合解決方案，為後續功能開發奠定堅實基礎。