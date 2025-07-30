# Smart Learning 專案初始化指南

## 專案概述

Smart Learning 是一個協助使用者學習英文的智能學習系統，結合了現代前端技術和 AI 人工智慧，提供個人化的英語學習體驗。

## 系統架構

```
smart-learning/
├── frontend/                 # React + TypeScript 前端應用
│   ├── src/
│   │   ├── components/       # UI 組件
│   │   ├── pages/           # 頁面組件
│   │   ├── store/           # 狀態管理
│   │   ├── services/        # API 服務
│   │   ├── hooks/           # 自定義 Hooks
│   │   ├── utils/           # 工具函數
│   │   └── types/           # TypeScript 類型定義
│   ├── public/
│   └── tests/
├── backend/                  # Go + Gin 後端 API
│   ├── cmd/                 # 應用程式入口點
│   ├── internal/            # 內部包
│   │   ├── handlers/        # HTTP 處理器
│   │   ├── services/        # 業務邏輯
│   │   ├── models/          # 資料模型
│   │   ├── repository/      # 資料存取層
│   │   ├── middleware/      # 中間件
│   │   └── config/          # 配置管理
│   ├── migrations/          # 資料庫遷移
│   └── tests/
├── database/                # 資料庫相關
│   ├── migrations/          # SQL 遷移文件
│   └── seeds/               # 種子資料
└── docs/                    # 專案文檔
    ├── api/                 # API 文檔
    ├── frontend/            # 前端文檔
    ├── backend/             # 後端文檔
    └── deployment/          # 部署文檔
```

## 環境需求

### 前端開發環境
- Node.js >= 18.0.0
- npm >= 9.0.0 或 pnpm >= 8.0.0
- VS Code (推薦)

### 後端開發環境
- Go >= 1.21
- PostgreSQL >= 15

### 外部服務
- Supabase (PostgreSQL 資料庫服務)
- Claude Haiku API (AI 服務)

## 專案初始化步驟

### 1. 建立專案結構

```bash
# 建立主要目錄
mkdir smart-learning
cd smart-learning

# 建立前端專案
mkdir frontend
cd frontend
npm create vite@latest . -- --template react-ts
npm install

# 安裝前端依賴
npm install @tanstack/react-query @tanstack/react-router zustand
npm install tailwindcss @tailwindcss/forms @tailwindcss/typography
npm install @hookform/resolvers react-hook-form zod
npm install lucide-react @radix-ui/react-dialog @radix-ui/react-dropdown-menu
npm install -D @types/node @vitejs/plugin-react
npm install -D vitest @testing-library/react @testing-library/jest-dom

# 建立後端專案
cd ../
mkdir backend
cd backend
go mod init smart-learning-backend

# 安裝後端依賴
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/golang-migrate/migrate/v4
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get github.com/stretchr/testify
```

### 2. 配置開發環境

#### 前端配置

建立 `vite.config.ts`:
```typescript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

建立 `tailwind.config.js`:
```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

#### 後端配置

建立 `cmd/main.go`:
```go
package main

import (
    "log"
    "smart-learning-backend/internal/config"
    "smart-learning-backend/internal/server"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    srv := server.New(cfg)
    if err := srv.Start(); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
```

### 3. 環境變數設定

#### 前端 `.env`
```env
VITE_API_URL=http://localhost:8080/api
```

#### 後端 `.env`
```env
PORT=8080
DATABASE_URL=postgresql://username:password@localhost/smart_learning_db
JWT_SECRET=your_jwt_secret_key
CLAUDE_API_KEY=your_claude_api_key
```

### 4. 資料庫設定

#### Supabase 設定
1. 註冊 Supabase 帳號
2. 建立新專案
3. 設定 Row Level Security (RLS)

#### 本地開發資料庫
```sql
-- 建立資料庫
CREATE DATABASE smart_learning_db;

-- 切換到資料庫
\c smart_learning_db;

-- 建立用戶表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255),
    google_id VARCHAR(255),
    learning_level INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. 開發工具配置

#### VS Code 設定 (`.vscode/settings.json`)
```json
{
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint"
}
```

#### Git 設定
```bash
# 初始化 Git
git init
git add .
git commit -m "feat: initial project setup"

# 建立 .gitignore
echo "node_modules/" >> .gitignore
echo ".env" >> .gitignore
echo "dist/" >> .gitignore
echo "*.log" >> .gitignore
```

## 快速啟動

### 開發模式
```bash
# 啟動前端
cd frontend
npm run dev

# 啟動後端
cd backend
go run cmd/main.go
```

### 測試
```bash
# 前端測試
cd frontend
npm run test

# 後端測試
cd backend
go test ./...
```

## 下一步

1. 閱讀 [技術架構文檔](./02-technical-architecture.md)
2. 參考 [開發流程指南](./03-development-workflow.md)
3. 查看 [功能模組文檔](./04-feature-modules.md)

## 常見問題

### Q: 如何設定 Claude API？
A: 前往 Anthropic 官網註冊帳號，取得 API Key 後設定在環境變數中。

### Q: Supabase 連線失敗怎麼辦？
A: 檢查網路連線和 DATABASE_URL 設定，確保防火牆允許連線。

### Q: 前端代理設定無效？
A: 確保後端服務正在運行，並檢查 vite.config.ts 中的代理設定。