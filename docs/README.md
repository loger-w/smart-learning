# Smart Learning 智能英語學習平台

一個結合 AI 技術的英語單字學習 web 應用，幫助使用者透過圖卡模式和 AI 助手來提升詞彙學習效率。

## 🌐 線上試用

**[立即體驗 Smart Learning](https://smart-learning-five.vercel.app/)**

> 歡迎註冊帳號體驗完整功能，或查看現有功能展示

## 📋 主要功能

- **圖卡學習模式** - 互動式翻卡介面，支援自訂學習節奏 `🚧 開發中`
- **AI 智能解釋** - 整合 Claude AI，提供詞彙解釋、同義詞和記憶法 `🚧 開發中`
- **學習進度追蹤** - 記錄學習表現，分析記憶強度 `🚧 開發中`
- **個人化清單** - 建立和管理自己的詞彙學習清單 `🚧 開發中`
- **響應式設計** - 支援桌面和行動裝置 `🚧 開發中`

> 📝 **專案狀態**: 持續開發中，部分功能正在完善，歡迎體驗現有功能並提供回饋

## 🛠 技術特色

- **前端**: React 19 + TypeScript + Vite + TailwindCSS
- **路由**: TanStack Router (type-safe routing)
- **狀態管理**: Zustand + TanStack Query
- **UI 元件**: Shadcn UI
- **後端**: Go + Gin framework
- **資料庫**: PostgreSQL (Supabase)
- **AI 整合**: Claude Haiku API
- **部署**: Vercel (前端) + Railway (後端)

## 💻 本地開發

### 環境需求
- Node.js >= 18
- Go >= 1.24
- PostgreSQL

### 快速啟動

```bash
# 複製專案
git clone <repository-url>
cd smart-learning

# 前端設定
cd frontend
npm install
cp .env.example .env.local
npm run dev

# 後端設定 (另一個終端)
cd backend
go mod tidy
cp .env.example .env
go run cmd/main.go
```

### 測試指令

```bash
# 前端測試
cd frontend
npm run test
npm run type-check
npm run lint

# 後端測試
cd backend
make test
make coverage
```

## 🔧 專案架構

```
smart-learning/
├── frontend/          # React 前端應用
│   ├── src/
│   │   ├── features/  # 功能模組 (auth, dashboard)
│   │   ├── components/# UI 元件
│   │   ├── hooks/     # 自訂 hooks
│   │   └── stores/    # Zustand 狀態管理
├── backend/           # Go 後端 API
│   ├── cmd/           # 應用進入點
│   └── pkg/           # 共用套件
└── docs/             # 專案文檔
```

## 📖 專案背景

這是一個個人 side project，主要目的是：
- 實驗現代前端技術棧的整合應用
- 探索 AI 在教育場景的實際應用
- 練習全端開發和部署流程

專案採用 monorepo 結構，前後端分離設計，注重型別安全和開發體驗。

---

**如果這個專案對你有幫助，歡迎給個 ⭐ 支持！**