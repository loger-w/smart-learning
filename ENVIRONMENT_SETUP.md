# 環境建置指南

本文件提供 Smart Learning 專案的完整環境建置步驟和待辦事項追蹤。

## 📋 建置狀態總覽

- [ ] **前端環境建置** - 未開始
- [ ] **後端環境建置** - 未開始  
- [ ] **CI/CD 配置** - 未開始
- [ ] **數據庫設置** - 未開始
- [ ] **部署配置** - 未開始

## 🔧 系統需求

### 必要軟體版本
- **Node.js** >= 18.0.0
- **Go** >= 1.21.0
- **Git** >= 2.30.0
- **PostgreSQL** >= 14.0 (或 Supabase 帳戶)

### 推薦開發工具
- **VS Code** (推薦擴展: Go, TypeScript, TailwindCSS IntelliSense)
- **Postman** 或 **Insomnia** (API 測試)
- **TablePlus** 或 **pgAdmin** (資料庫管理)

## 🎯 前端環境建置

### 狀態: ❌ 未完成

#### 1. 創建目錄結構
```bash
# 建立前端專案結構
mkdir -p frontend/src/{components,pages,hooks,services,types,stores,utils}
mkdir -p frontend/src/components/{ui,features,layout}
mkdir -p frontend/public
mkdir -p frontend/src/assets/{images,icons,fonts}
```

#### 2. 初始化 React 專案
```bash
cd frontend

# 使用 Vite 建立 React TypeScript 專案
npm create vite@latest . -- --template react-ts
npm install
```

#### 3. 安裝核心依賴
```bash
# 狀態管理和路由
npm install @tanstack/react-query @tanstack/react-router
npm install zustand react-hook-form

# UI 和樣式
npm install tailwindcss @tailwindcss/forms @tailwindcss/typography
npm install lucide-react class-variance-authority clsx tailwind-merge

# 實用工具
npm install axios date-fns
npm install @hookform/resolvers zod
```

#### 4. 安裝開發依賴
```bash
# TypeScript 支援
npm install -D @types/node

# 測試框架
npm install -D vitest @testing-library/react @testing-library/jest-dom
npm install -D @testing-library/user-event jsdom

# 程式碼品質
npm install -D eslint-config-prettier prettier
npm install -D @typescript-eslint/eslint-plugin @typescript-eslint/parser
```

#### 5. 設置 Shadcn UI
```bash
npx shadcn-ui@latest init
npx shadcn-ui@latest add button card input form
npx shadcn-ui@latest add toast dialog sheet
npx shadcn-ui@latest add table badge avatar
```

#### 6. 配置檔案設置
```bash
# 環境變數
cp .env.example .env.local

# TailwindCSS 配置
# Vite 配置
# TypeScript 配置
# ESLint 配置
# Prettier 配置
```

#### 待辦事項 (Frontend)
- [ ] 創建基本目錄結構
- [ ] 初始化 Vite + React 專案
- [ ] 安裝並配置 TailwindCSS
- [ ] 設置 Shadcn UI 組件庫
- [ ] 配置 TanStack Query 和 Router
- [ ] 設置 Zustand 狀態管理
- [ ] 配置 React Hook Form
- [ ] 建立基本 Layout 組件
- [ ] 設置測試環境 (Vitest)
- [ ] 配置 ESLint 和 Prettier
- [ ] 建立環境變數模板

## 🚀 後端環境建置

### 狀態: ❌ 未完成

#### 1. 創建目錄結構
```bash
# 建立後端專案結構
mkdir -p backend/{cmd,internal,pkg,migrations,configs,scripts}
mkdir -p backend/internal/{handlers,services,repositories,models,middleware}
mkdir -p backend/pkg/{database,auth,utils,logger}
mkdir -p backend/tests/{unit,integration}
```

#### 2. 初始化 Go 模組
```bash
cd backend

# 初始化 Go module
go mod init smart-learning-backend
```

#### 3. 安裝核心依賴
```bash
# Web 框架
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors

# 資料庫
go get github.com/lib/pq
go get github.com/golang-migrate/migrate/v4

# 身份驗證
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

# 配置和工具
go get github.com/joho/godotenv
go get github.com/google/uuid
```

#### 4. 安裝開發工具
```bash
# 測試框架
go get github.com/stretchr/testify

# 程式碼品質工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 資料庫遷移工具
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 熱重載工具
go install github.com/cosmtrek/air@latest
```

#### 5. 創建配置檔案
```bash
# 環境變數
cp .env.example .env

# Air 配置 (熱重載)
touch .air.toml

# Docker 配置
touch Dockerfile

# Makefile
touch Makefile
```

#### 待辦事項 (Backend)
- [ ] 創建基本目錄結構
- [ ] 初始化 Go module
- [ ] 安裝 Gin 框架和中間件
- [ ] 設置資料庫連接 (PostgreSQL)
- [ ] 配置 JWT 身份驗證
- [ ] 建立基本路由結構
- [ ] 實作 CRUD 操作基類
- [ ] 設置資料庫遷移
- [ ] 配置 CORS 和安全性中間件
- [ ] 建立錯誤處理機制
- [ ] 設置日誌系統
- [ ] 配置測試環境
- [ ] 建立 Makefile 和腳本

## 🔄 CI/CD 配置

### 狀態: ❌ 未完成

#### 1. GitHub Actions 工作流程
```bash
# 創建 GitHub Actions 目錄
mkdir -p .github/workflows

# 工作流程檔案
touch .github/workflows/frontend-ci.yml
touch .github/workflows/backend-ci.yml
touch .github/workflows/deploy.yml
```

#### 2. Docker 配置
```bash
# Docker 檔案
touch Dockerfile.frontend
touch Dockerfile.backend
touch docker-compose.yml
touch docker-compose.dev.yml
touch .dockerignore
```

#### 3. 部署配置
```bash
# 部署配置目錄
mkdir -p deploy/{vercel,railway}

# 配置檔案
touch vercel.json
touch railway.json
```

#### 4. 腳本和工具
```bash
# 工具腳本目錄
mkdir -p scripts

# 環境設置腳本
touch scripts/setup-env.sh
touch scripts/validate-env.sh
touch scripts/build.sh
touch scripts/test.sh

# 賦予執行權限
chmod +x scripts/*.sh
```

#### 待辦事項 (CI/CD)
- [ ] 設置前端 CI 工作流程
- [ ] 設置後端 CI 工作流程
- [ ] 配置自動化測試流程
- [ ] 建立 Docker 容器化配置
- [ ] 設置 Vercel 前端部署
- [ ] 設置 Railway 後端部署
- [ ] 配置環境變數管理
- [ ] 建立部署腳本
- [ ] 設置代碼品質檢查
- [ ] 配置自動化依賴更新

## 🗄️ 資料庫設置

### 狀態: ❌ 未完成

#### 1. 本地開發資料庫
```bash
# 使用 Docker 啟動 PostgreSQL
docker run --name smart-learning-db \
  -e POSTGRES_DB=smart_learning_dev \
  -e POSTGRES_USER=dev_user \
  -e POSTGRES_PASSWORD=dev_password \
  -p 5432:5432 \
  -d postgres:14
```

#### 2. 資料庫遷移
```bash
cd backend

# 建立遷移檔案
migrate create -ext sql -dir migrations -seq init_schema
migrate create -ext sql -dir migrations -seq create_users_table
migrate create -ext sql -dir migrations -seq create_word_lists_table

# 執行遷移
migrate -path migrations -database "postgresql://dev_user:dev_password@localhost/smart_learning_dev?sslmode=disable" up
```

#### 3. 測試資料
```bash
# 建立測試資料腳本
touch scripts/seed-db.sql
touch scripts/test-data.sql
```

#### 待辦事項 (Database)
- [ ] 設置本地 PostgreSQL 資料庫
- [ ] 建立資料庫遷移檔案
- [ ] 設計資料表結構
- [ ] 建立索引和約束
- [ ] 準備測試資料
- [ ] 配置 Supabase 生產環境
- [ ] 設置資料庫備份策略

## 🚢 部署環境設置

### 狀態: ❌ 未完成

#### 前端部署 (Vercel)
- [ ] 連接 GitHub 儲存庫
- [ ] 配置建置命令和設定
- [ ] 設置環境變數
- [ ] 配置自定義域名
- [ ] 設置分支部署策略

#### 後端部署 (Railway)
- [ ] 連接 GitHub 儲存庫
- [ ] 配置 Go 建置環境
- [ ] 設置 PostgreSQL 資料庫
- [ ] 配置環境變數
- [ ] 設置健康檢查
- [ ] 配置自動部署

## 📝 環境變數範本

### Frontend (.env.local)
```bash
# API 配置
VITE_API_BASE_URL=http://localhost:8080/api
VITE_API_TIMEOUT=10000

# 第三方服務
VITE_CLAUDE_API_KEY=your_claude_api_key_here

# 應用配置
VITE_APP_NAME=Smart Learning
VITE_APP_VERSION=1.0.0
VITE_APP_ENV=development
```

### Backend (.env)
```bash
# 伺服器配置
PORT=8080
GIN_MODE=debug

# 資料庫配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=dev_user
DB_PASSWORD=dev_password
DB_NAME=smart_learning_dev
DB_SSL_MODE=disable

# JWT 配置
JWT_SECRET=your_jwt_secret_here
JWT_EXPIRY=24h

# 第三方 API
CLAUDE_API_KEY=your_claude_api_key_here
CLAUDE_API_URL=https://api.anthropic.com

# 日誌配置
LOG_LEVEL=info
LOG_FORMAT=json
```

## 🔧 開發工具配置

### VS Code 推薦擴展
```json
{
  "recommendations": [
    "golang.Go",
    "ms-vscode.vscode-typescript-next",
    "bradlc.vscode-tailwindcss",
    "esbenp.prettier-vscode",
    "ms-vscode.vscode-eslint",
    "ms-vscode.vscode-json"
  ]
}
```

### Git Hooks 設置
```bash
# 設置 pre-commit hook
touch .githooks/pre-commit
chmod +x .githooks/pre-commit

# 配置 Git 使用自定義 hooks
git config core.hooksPath .githooks
```

## 📚 後續步驟

完成環境建置後，請參考以下文件進行開發：

1. **功能開發**: `docs/04-feature-modules.md`
2. **API 設計**: `docs/02-technical-architecture.md`
3. **UI 設計**: `docs/06-ui-design-system.md`
4. **測試策略**: `docs/03-development-workflow.md`
5. **部署流程**: `docs/05-deployment-guide.md`

## 🆘 疑難排解

### 常見問題
1. **Node.js 版本問題**: 使用 nvm 管理 Node.js 版本
2. **Go 模組問題**: 檢查 GOPROXY 設置
3. **資料庫連接問題**: 確認 PostgreSQL 服務運行狀態
4. **權限問題**: 檢查腳本執行權限設置

### 獲取幫助
- 查看專案 README.md
- 檢查 `docs/` 目錄中的詳細文件
- 提交 Issue 到專案儲存庫