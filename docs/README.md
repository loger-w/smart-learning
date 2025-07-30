# Smart Learning 專案文檔

歡迎來到 Smart Learning 英語學習系統的開發文檔。本系統結合了現代前端技術、後端 API 服務和 AI 人工智慧，為使用者提供個人化的英語學習體驗。

## 📚 文檔導航

### 🚀 [01. 專案初始化指南](./01-project-initialization.md)
- 環境需求與設定
- 專案結構建立
- 依賴安裝與配置
- 快速啟動指南

### 🏗️ [02. 技術架構文檔](./02-technical-architecture.md)
- 系統架構概覽
- 前後端技術棧詳細說明
- 資料庫設計
- AI 服務整合
- 安全性與效能考量

### 🔄 [03. 開發流程指南](./03-development-workflow.md)
- Monorepo Git 管理策略
- 程式碼品質標準
- 測試策略
- CI/CD 流程
- 程式碼審查指南

### 🧩 [04. 功能模組文檔](./04-feature-modules.md)
- 使用者認證系統
- 學習清單管理
- 圖卡學習模式
- AI 智能回應
- 等級化學習
- 學習追蹤與分析
- 搜尋與篩選

### 🚀 [05. 部署指南](./05-deployment-guide.md)
- Vercel 前端部署
- Railway 後端部署 (推薦)
- Supabase 資料庫配置
- CI/CD 自動化流程
- 監控與日誌管理

### 🎨 [06. UI 設計系統](./06-ui-design-system.md)
- 簡潔明確的設計原則
- 色彩與字體系統
- 組件設計規範
- 響應式設計
- 無障礙設計標準

## 🎯 專案概述

Smart Learning 是一個智能英語學習平台，主要功能包括：

### 基礎功能
1. **使用者認證系統** - 註冊/登入/登出
2. **學習清單管理** - 建立、編輯、分享學習清單

### 學習功能
3. **圖卡學習模式** - 翻卡式學習介面與進度追蹤
4. **AI 智能回應** - 多樣化解釋、同義詞推薦、記憶法生成
5. **等級化學習** - 根據使用者等級調整內容難度

### 進階功能
6. **學習追蹤與分析** - 正確率分析、定時學習提醒
7. **搜尋與篩選** - 清單搜尋、狀態篩選、標籤分類

## 🛠️ 技術棧

### 前端技術
- **Vite** - 快速建構工具
- **React + TypeScript** - 組件化開發
- **TailwindCSS** - 樣式框架
- **Shadcn UI** - UI 組件庫
- **TanStack Query + Zustand** - 狀態管理
- **TanStack Router** - 類型安全路由
- **React Hook Form** - 表單處理
- **Vitest + React Testing Library** - 測試框架

### 後端技術
- **Go** - 高效能後端語言
- **Gin** - 輕量級 Web 框架
- **PostgreSQL (Supabase)** - 資料庫服務
- **Claude Haiku API** - AI 服務
- **JWT** - 使用者認證
- **Go Testing + Testify** - 測試工具

## 🚀 快速開始

### 1. 環境準備
```bash
# 確保已安裝 Node.js >= 18 和 Go >= 1.21
node --version
go version
```

### 2. 專案初始化
```bash
# Clone 專案
git clone <repository-url>
cd smart-learning

# 安裝前端依賴
cd frontend
npm install

# 安裝後端依賴
cd ../backend
go mod tidy
```

### 3. 環境配置
```bash
# 複製環境變數範本
cp frontend/.env.example frontend/.env
cp backend/.env.example backend/.env

# 編輯環境變數，設定資料庫連線和 API Key
```

### 4. 啟動服務
```bash
# 啟動前端開發伺服器
cd frontend
npm run dev

# 啟動後端 API 伺服器
cd backend
go run cmd/main.go
```

## 📋 開發檢查清單

在開始開發前，請確保：

- [ ] 已閱讀 [專案初始化指南](./01-project-initialization.md)
- [ ] 了解 [技術架構設計](./02-technical-architecture.md)
- [ ] 熟悉 [開發流程規範](./03-development-workflow.md)
- [ ] 環境設定完成，專案能正常啟動
- [ ] Git hooks 配置完成
- [ ] IDE 擴充功能安裝完成

## 🔗 相關連結

- **專案 Repository**: [GitHub](https://github.com/your-org/smart-learning)
- **API 文檔**: [Swagger UI](http://localhost:8080/swagger)
- **前端預覽**: [本地開發](http://localhost:3000)
- **資料庫管理**: [Supabase Dashboard](https://supabase.com/dashboard)
- **AI API 文檔**: [Anthropic Claude](https://docs.anthropic.com)

## 📞 支援與貢獻

### 問題回報
如果您遇到任何問題，請：
1. 查看相關文檔是否有解決方案
2. 搜尋現有的 Issues
3. 建立新的 Issue 並詳細描述問題

### 貢獻指南
1. Fork 專案
2. 建立功能分支 (`git checkout -b feature/amazing-feature`)
3. 遵循 [開發流程指南](./03-development-workflow.md)
4. 提交 Pull Request

### 開發團隊
- **專案負責人**: [您的名字]
- **前端開發**: [團隊成員]
- **後端開發**: [團隊成員]
- **UI/UX 設計**: [團隊成員]

## 📝 版本記錄

### v1.0.0 (計劃中)
- ✅ 使用者認證系統
- ✅ 學習清單管理
- ✅ 圖卡學習模式
- ✅ AI 智能回應
- ✅ 等級化學習
- ✅ 學習追蹤與分析
- ✅ 搜尋與篩選功能

### 未來計劃
- 📱 行動應用程式 (React Native)
- 🎵 語音識別與發音評估
- 👥 社群學習功能
- 🎯 學習挑戰與成就系統
- 📊 教師管理後台

---

## 📄 授權條款

本專案採用 MIT 授權條款，詳細內容請參閱 [LICENSE](../LICENSE) 檔案。

---

**Happy Coding! 🎉**

如有任何疑問，請參閱相關文檔或聯繫開發團隊。