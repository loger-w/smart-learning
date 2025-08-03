# 登入頁面重新設計文檔

## 修改概述

本次對登入頁面進行了重新設計，主要目標是建立一個簡潔、功能完整且易於維護的登入介面。

## 主要修改項目

### 1. 路由修正 (`src/routes/index.tsx`)

**問題描述：**
- 原始程式碼錯誤使用了 `createBrowserRouter`，這是 React Router 的 API
- 但專案使用的是 TanStack Router，應該使用 `createRouter`

**修正內容：**
```typescript
// 修正前
import { createBrowserRouter } from '@tanstack/react-router'
export const router = createBrowserRouter({
  routeTree,
})

// 修正後
import { createRouter } from '@tanstack/react-router'
export const router = createRouter({
  routeTree,
})
```

**影響範圍：**
- 修正了路由初始化錯誤
- 確保 TanStack Router 正常運作

### 2. 登入頁面重新設計 (`src/pages/LoginPage.tsx`)

**設計理念：**
- 簡潔清晰的使用者介面
- 完整的表單驗證機制
- 良好的使用者體驗和無障礙設計

**主要功能：**
1. **表單驗證**
   - 使用 React Hook Form + Zod 進行表單驗證
   - 電子郵件格式驗證
   - 密碼長度驗證（至少6個字符）
   - 即時錯誤提示

2. **使用者介面元素**
   - 品牌 Logo 和標題區域
   - 電子郵件輸入欄位
   - 密碼輸入欄位（含顯示/隱藏功能）
   - 登入按鈕（含載入狀態）
   - 忘記密碼連結
   - 註冊連結

3. **互動功能**
   - 密碼顯示/隱藏切換
   - 表單提交處理
   - 載入狀態指示
   - 錯誤訊息顯示

**技術特點：**
- 使用 TypeScript 確保型別安全
- 遵循無障礙設計原則（ARIA 標籤、語義化 HTML）
- 響應式設計，適配不同螢幕尺寸
- TailwindCSS 樣式系統

**移除的功能：**
- 移除了複雜的身份驗證 Hook 依賴
- 移除了額外的功能特色展示區域
- 簡化了頁面結構，專注於核心登入功能

## 程式碼結構

### 狀態管理
```typescript
const [showPassword, setShowPassword] = useState(false)
const [isLoading, setIsLoading] = useState(false)
```

### 表單處理
```typescript
const { 
  register, 
  handleSubmit, 
  formState: { errors },
  setError 
} = useForm<LoginFormData>({
  resolver: zodResolver(loginSchema),
})
```

### 驗證規則
```typescript
const loginSchema = z.object({
  email: z.string().email('請輸入有效的電子郵件'),
  password: z.string().min(6, '密碼至少需要6個字符'),
})
```

## 後續開發建議

### 1. 身份驗證整合
- 整合實際的 API 登入端點
- 實作 JWT Token 處理
- 加入身份驗證狀態管理

### 2. 頁面導航
- 實作註冊頁面
- 建立忘記密碼流程
- 加入登入成功後的頁面導航

### 3. 使用者體驗優化
- 加入表單自動填入支援
- 實作記住登入狀態功能
- 加入社群媒體登入選項

### 4. 錯誤處理
- 細化不同類型的錯誤訊息
- 加入網路錯誤處理
- 實作重試機制

## 檔案異動清單

### 修改的檔案
1. `src/routes/index.tsx` - 修正路由初始化
2. `src/pages/LoginPage.tsx` - 重新設計登入頁面

### 相關檔案（未修改）
- `src/routes/routeTree.tsx` - 路由定義
- `src/providers/AppProviders.tsx` - 全域 Provider 設定
- `src/components/auth/LoginForm.tsx` - 原有登入表單元件（已不使用）

## 注意事項

1. **向後相容性**：新的登入頁面移除了對 `useAuth` Hook 的依賴，需要在後續開發中重新整合身份驗證邏輯。

2. **元件架構**：目前將所有登入邏輯整合在單一頁面元件中，後續可考慮拆分為更小的可重用元件。

3. **樣式系統**：完全使用 TailwindCSS，確保與專案整體設計系統一致。

4. **測試考慮**：建議為新的登入頁面撰寫對應的單元測試和整合測試。