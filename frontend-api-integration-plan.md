# Smart Learning 前端 API 串接規劃文件

## 專案概況

Smart Learning 是一個智能英語學習平台，後端使用 Go + Gin 框架，前端使用 React + TypeScript + TanStack Query。本文件詳細規劃前端如何與後端 API 進行串接的完整實作策略。

## 後端 API 分析

### 可用端點
- **認證系統** (`/api/v1/auth`):
  - `POST /register` - 用戶註冊
  - `POST /login` - 用戶登入
  - `POST /logout` - 用戶登出 (需 JWT)
  - `GET /me` - 獲取當前用戶資料 (需 JWT)
- **系統端點**:
  - `GET /health` - 健康檢查
  - `GET /api/v1/ping` - API 測試端點

### API 回應格式
```typescript
interface APIResponse<T = any> {
  success: boolean
  message?: string
  data?: T
  error?: {
    code: string
    message: string
  }
  errors?: Record<string, string[]> // 驗證錯誤
}
```

### 用戶模型
```typescript
interface User {
  id: number
  email: string
  username: string
  learning_level: number
  avatar_url?: string
  created_at: string
  updated_at: string
}
```

## 現有前端架構分析

### 已實作部分
✅ **認證服務** (`authService.ts`):
- Axios 實例配置
- 基本 CRUD 操作方法
- 請求/回應攔截器

✅ **狀態管理** (`authStore.ts`):
- Zustand + persist 中間件
- 用戶狀態和認證 token 管理

✅ **型別定義** (`types/auth.ts`):
- 基本認證相關型別

### 需要修正的問題
❌ **API 回應格式不匹配**:
- 前端期望：`{ user, token }`
- 後端實際：`{ success, message, data: { user, token } }`

❌ **錯誤處理不完整**:
- 缺乏後端驗證錯誤格式處理
- 缺乏統一錯誤型別定義

❌ **缺少功能**:
- 無 `/me` 端點串接
- 無 refresh token 實作（後端未提供）

## 實作策略

### 階段一：修正核心認證功能
1. **更新 API 回應型別**
2. **修正 authService 方法**
3. **完善錯誤處理機制**
4. **添加 `/me` 端點串接**

### 階段二：增強功能實作
1. **創建統一 API 客戶端**
2. **實作 TanStack Query 鉤子**
3. **創建錯誤邊界組件**
4. **添加載入狀態管理**

### 階段三：測試與優化
1. **API 串接測試**
2. **錯誤情境測試**
3. **效能優化**

## 詳細實作計畫

### 1. 型別定義更新

#### 新增統一 API 回應型別
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

#### 更新認證型別
```typescript
// types/auth.ts
export interface AuthResponse {
  user: User
  token: string
}

export interface LoginResponse extends APIResponse<AuthResponse> {}
export interface RegisterResponse extends APIResponse<AuthResponse> {}
export interface MeResponse extends APIResponse<{ user: User }> {}
```

### 2. 統一 API 客戶端

#### 創建基礎 API 客戶端
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
  
  // 統一的 API 錯誤處理
  private handleResponseError(error: AxiosError) {
    if (error.response?.status === 401) {
      this.handleUnauthorized()
    }
    return Promise.reject(error)
  }
}
```

### 3. 認證服務重構

#### 更新 authService 以匹配後端格式
```typescript
// services/authService.ts
export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await authAPI.post<LoginResponse>('/login', credentials)
    
    if (!response.data.success) {
      throw new AuthError(response.data.error?.message || '登入失敗')
    }
    
    return response.data.data! // 確保有 data
  },
  
  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await authAPI.post<RegisterResponse>('/register', data)
    
    if (!response.data.success) {
      // 處理驗證錯誤
      if (response.data.errors) {
        throw new ValidationError(response.data.errors)
      }
      throw new AuthError(response.data.error?.message || '註冊失敗')
    }
    
    return response.data.data!
  },
  
  async getMe(): Promise<User> {
    const response = await authAPI.get<MeResponse>('/me')
    
    if (!response.data.success) {
      throw new AuthError(response.data.error?.message || '獲取用戶資料失敗')
    }
    
    return response.data.data!.user
  }
}
```

### 4. TanStack Query 整合

#### 創建認證相關查詢鉤子
```typescript
// hooks/useAuth.ts
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { authService } from '@/services/authService'
import { useAuthStore } from '@/stores/authStore'

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
    staleTime: 5 * 60 * 1000 // 5 分鐘
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

### 5. 錯誤處理機制

#### 創建自定義錯誤類別
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
}
```

#### 創建錯誤邊界組件
```typescript
// components/ErrorBoundary.tsx
import React from 'react'
import { ValidationError, AuthError } from '@/utils/errors'

interface Props {
  children: React.ReactNode
  fallback?: React.ComponentType<{ error: Error }>
}

export const ErrorBoundary: React.FC<Props> = ({ children, fallback: Fallback }) => {
  // 實作錯誤邊界邏輯
}
```

### 6. 環境配置

#### 環境變數設定
```env
# .env.local
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_APP_NAME=Smart Learning
```

#### 開發環境設定
```typescript
// config/api.ts
export const API_CONFIG = {
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
  timeout: 10000,
  retries: 3
}
```

## 檔案結構規劃

```
frontend/src/
├── services/
│   ├── apiClient.ts         # 統一 API 客戶端
│   ├── authService.ts       # 認證服務 (重構)
│   └── index.ts            # 服務入口點
├── hooks/
│   ├── useAuth.ts          # 認證相關鉤子
│   ├── useAPI.ts           # 通用 API 鉤子
│   └── index.ts            # 鉤子入口點
├── types/
│   ├── api.ts              # API 相關型別
│   ├── auth.ts             # 認證型別 (更新)
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

## 測試策略

### 單元測試
- API 服務方法測試
- 錯誤處理邏輯測試
- 狀態管理測試

### 整合測試
- API 端點串接測試
- 認證流程測試
- 錯誤情境測試

### 測試工具
- Vitest + React Testing Library
- MSW (Mock Service Worker) 模擬 API

## 部署考慮

### 生產環境配置
- API 基礎 URL 環境變數化
- 錯誤日誌收集
- 效能監控

### 安全性
- JWT token 安全存儲
- HTTPS 強制使用
- CSP 內容安全政策

## 開發時程

### 第一週：核心重構
- Day 1-2: 型別定義更新
- Day 3-4: authService 重構
- Day 5: 錯誤處理機制

### 第二週：功能擴展
- Day 1-2: TanStack Query 整合
- Day 3-4: 統一 API 客戶端
- Day 5: 測試撰寫

### 第三週：測試與優化
- Day 1-3: 整合測試
- Day 4-5: 效能優化與文件

## 總結

這個規劃文件提供了完整的前端 API 串接實作策略，重點包括：

1. **標準化 API 介面**：統一後端回應格式處理
2. **強化錯誤處理**：完善的錯誤型別和處理機制
3. **現代化狀態管理**：TanStack Query + Zustand 結合
4. **型別安全**：完整的 TypeScript 型別定義
5. **可維護性**：清晰的檔案結構和職責分離

實作完成後，將能提供穩定、高效、可維護的 API 串接解決方案，為後續功能開發奠定堅實基礎。