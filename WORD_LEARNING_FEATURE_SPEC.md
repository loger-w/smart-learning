# Smart Learning 記憶單字功能開發規格書

## 📋 功能需求概述

### 核心功能
1. **單字清單管理系統**
   - 建立、編輯、刪除單字清單
   - 單字新增與AI輔助資訊補充
   - CEFR等級導向的內容生成

2. **記憶卡片學習系統**
   - 互動式卡片學習介面
   - 艾賓浩斯記憶曲線複習排程
   - 學習進度追蹤與分析

## 🎯 功能需求詳細規格

### 1. 單字清單管理功能

#### 1.1 清單基本操作
- **建立清單**
  - 使用者可建立命名的單字清單
  - 支援清單描述與CEFR等級標記
  - 清單可設為公開或私人

- **編輯清單**
  - 修改清單名稱、描述
  - 調整清單的目標CEFR等級
  - 單字的新增、移除、重新排序

- **刪除清單**
  - 軟刪除機制（保留學習記錄）
  - 刪除前確認對話
  - 相關學習數據的處理

#### 1.2 單字管理功能

**手動新增模式**：
```
輸入單字 → 手動設定：
- 詞性 (名詞/動詞/形容詞等)
- 中文解釋
- 英文例句
- 同義詞/反義詞
- 相關詞彙
- CEFR適用等級
```

**AI輔助模式**：
```
輸入單字 → AI自動生成：
- 根據使用者CEFR等級調整解釋複雜度
- 生成適當難度的例句
- 推薦同級別的同義詞
- 提供記憶技巧建議
- 標注詞彙的CEFR等級
```

#### 1.3 CEFR等級整合

**等級定義**：
- A1 (初學者): 基本日常詞彙
- A2 (基礎): 常用詞彙與片語
- B1 (中級): 抽象概念詞彙
- B2 (中高級): 專業與學術詞彙
- C1 (高級): 進階與專門術語
- C2 (精通): 母語者程度詞彙

**等級適應功能**：
- 詞彙難度自動評估
- 解釋內容依等級調整
- 例句複雜度分級
- 學習建議個人化

### 2. 記憶卡片系統

#### 2.1 學習會話管理
- 使用者選擇要學習的清單
- 設定本次學習的單字數量 (5-50個)
- 支援打亂順序或按清單順序學習
- 學習中斷與續學功能

#### 2.2 卡片顯示邏輯
```
顯示流程：
1. 顯示英文單字
2. 使用者思考 → 點擊「翻開答案」
3. 顯示完整資訊：
   - 中文解釋
   - 詞性標示
   - 例句
   - 同義詞 (依CEFR等級)
4. 使用者自評：太簡單/適中/困難
5. 下一張卡片
```

#### 2.3 艾賓浩斯記憶曲線實作

**複習間隔設計**：
- 第1次：學習後1天
- 第2次：學習後3天  
- 第3次：學習後7天
- 第4次：學習後15天
- 第5次：學習後30天
- 第6次：學習後60天

**動態調整機制**：
- 答對：延長下次複習間隔
- 答錯：縮短下次複習間隔
- 連續答對3次：標記為「已掌握」
- 依CEFR等級調整基準間隔

#### 2.4 學習記錄系統
- 記錄每次學習的詳細數據
- 答題正確率統計
- 學習時間追蹤
- 困難單字標記
- 學習曲線分析

## 🏗️ 技術開發流程

### 階段一：資料庫設計與基礎架構 (1週)

**里程碑 1.1：CEFR系統重構**
- [ ] 用戶表 learning_level 欄位遷移為 cefr_level
- [ ] 建立CEFR等級轉換邏輯
- [ ] 更新現有API以支援CEFR
- [ ] 資料遷移腳本撰寫與測試

**里程碑 1.2：核心資料表建立**
- [ ] word_lists 表設計與建立
- [ ] words 表設計與建立  
- [ ] list_words 關聯表建立
- [ ] learning_records 學習記錄表
- [ ] review_schedules 複習排程表
- [ ] 索引與約束設定

**交付物**：
- 完整的資料庫遷移檔案
- 更新的API文檔
- 單元測試覆蓋率 > 80%

### 階段二：單字清單管理API (1週)

**里程碑 2.1：清單CRUD操作**
- [ ] 清單建立API (`POST /api/v1/lists`)
- [ ] 清單查詢API (`GET /api/v1/lists`)
- [ ] 清單更新API (`PUT /api/v1/lists/:id`)
- [ ] 清單刪除API (`DELETE /api/v1/lists/:id`)

**里程碑 2.2：單字管理功能**
- [ ] 單字新增API (`POST /api/v1/lists/:id/words`)
- [ ] 單字查詢API (`GET /api/v1/words/:id`)
- [ ] 單字更新API (`PUT /api/v1/words/:id`)
- [ ] 單字刪除API (`DELETE /api/v1/words/:id`)

**里程碑 2.3：AI整合功能**
- [ ] Claude API客戶端建立
- [ ] AI輔助單字資訊生成
- [ ] CEFR等級適應邏輯
- [ ] 提示工程優化

**交付物**：
- 完整的清單管理API
- Claude API整合模組
- Postman測試集合
- API文檔更新

### 階段三：記憶卡片與複習系統 (1週)

**里程碑 3.1：學習會話管理**
- [ ] 學習會話開始API (`POST /api/v1/learning/sessions`)
- [ ] 會話狀態管理API (`GET /api/v1/learning/sessions/:id`)
- [ ] 學習進度提交API (`POST /api/v1/learning/sessions/:id/progress`)
- [ ] 會話結束API (`PUT /api/v1/learning/sessions/:id/complete`)

**里程碑 3.2：艾賓浩斯演算法**
- [ ] 記憶曲線計算模組
- [ ] 複習排程生成邏輯
- [ ] 動態間隔調整機制
- [ ] 複習提醒系統

**里程碑 3.3：學習記錄系統**
- [ ] 學習數據記錄API
- [ ] 統計分析API (`GET /api/v1/analytics`)
- [ ] 進度查詢API (`GET /api/v1/learning/progress`)
- [ ] 複習排程API (`GET /api/v1/learning/reviews`)

**交付物**：
- 完整的學習系統API
- 記憶曲線演算法實作
- 學習分析功能
- 整合測試報告

### 階段四：優化與進階功能 (1週)

**里程碑 4.1：效能優化**
- [ ] 資料庫查詢優化
- [ ] API響應時間優化
- [ ] 快取機制實作
- [ ] 負載測試與調優

**里程碑 4.2：進階功能**
- [ ] 批量匯入單字功能
- [ ] 清單分享功能
- [ ] 學習統計儀表板API
- [ ] 個人化推薦系統

**里程碑 4.3：測試與部署**
- [ ] 完整的端到端測試
- [ ] 安全測試與驗證
- [ ] 部署腳本準備
- [ ] 監控與日誌系統

**交付物**：
- 生產就緒的API系統
- 完整的測試覆蓋
- 部署與監控方案
- 使用者手冊

## 📊 詳細開發需求

### 資料庫設計

#### 表格結構設計

**1. users 表更新**
```sql
ALTER TABLE users 
DROP COLUMN learning_level,
ADD COLUMN cefr_level VARCHAR(2) DEFAULT 'A1' 
CHECK (cefr_level IN ('A1', 'A2', 'B1', 'B2', 'C1', 'C2'));
```

**2. word_lists 表**
```sql
CREATE TABLE word_lists (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    target_cefr_level VARCHAR(2) DEFAULT 'A1',
    is_public BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**3. words 表**
```sql
CREATE TABLE words (
    id SERIAL PRIMARY KEY,
    word VARCHAR(100) NOT NULL,
    phonetic VARCHAR(200),
    cefr_level VARCHAR(2) NOT NULL,
    definitions JSONB NOT NULL, -- 各等級的定義
    examples JSONB, -- 各等級的例句
    synonyms JSONB, -- 同義詞
    antonyms JSONB, -- 反義詞
    related_words JSONB, -- 相關詞彙
    memory_tips TEXT, -- AI生成的記憶技巧
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(word, cefr_level)
);
```

**4. list_words 關聯表**
```sql
CREATE TABLE list_words (
    id SERIAL PRIMARY KEY,
    list_id INTEGER NOT NULL REFERENCES word_lists(id) ON DELETE CASCADE,
    word_id INTEGER NOT NULL REFERENCES words(id) ON DELETE CASCADE,
    position INTEGER NOT NULL DEFAULT 0,
    added_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(list_id, word_id)
);
```

**5. learning_records 學習記錄**
```sql
CREATE TABLE learning_records (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    word_id INTEGER NOT NULL REFERENCES words(id) ON DELETE CASCADE,
    session_id UUID NOT NULL,
    is_correct BOOLEAN NOT NULL,
    difficulty_rating INTEGER CHECK (difficulty_rating >= 1 AND difficulty_rating <= 5),
    response_time_ms INTEGER,
    learned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**6. review_schedules 複習排程**
```sql
CREATE TABLE review_schedules (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    word_id INTEGER NOT NULL REFERENCES words(id) ON DELETE CASCADE,
    current_interval_days INTEGER DEFAULT 1,
    ease_factor DECIMAL(3,2) DEFAULT 2.50,
    repetition_count INTEGER DEFAULT 0,
    next_review_date DATE NOT NULL,
    is_graduated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, word_id)
);
```

#### 索引設計
```sql
-- 效能優化索引
CREATE INDEX idx_word_lists_user_id ON word_lists(user_id);
CREATE INDEX idx_word_lists_public ON word_lists(is_public) WHERE is_public = TRUE;
CREATE INDEX idx_words_cefr_level ON words(cefr_level);
CREATE INDEX idx_words_word_lower ON words(LOWER(word));
CREATE INDEX idx_list_words_list_id ON list_words(list_id);
CREATE INDEX idx_learning_records_user_word ON learning_records(user_id, word_id);
CREATE INDEX idx_review_schedules_user_date ON review_schedules(user_id, next_review_date);
CREATE INDEX idx_review_schedules_date ON review_schedules(next_review_date) WHERE is_graduated = FALSE;
```

### API 端點設計

#### 清單管理 API

**1. 建立清單**
```
POST /api/v1/lists
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "TOEIC核心詞彙",
  "description": "TOEIC考試必備的核心詞彙清單",
  "target_cefr_level": "B2",
  "is_public": false
}

Response 201:
{
  "success": true,
  "message": "清單建立成功",
  "data": {
    "list": {
      "id": 1,
      "name": "TOEIC核心詞彙",
      "description": "TOEIC考試必備的核心詞彙清單",
      "target_cefr_level": "B2",
      "is_public": false,
      "word_count": 0,
      "created_at": "2025-01-01T00:00:00Z"
    }
  }
}
```

**2. 查詢使用者清單**
```
GET /api/v1/lists?page=1&limit=20&cefr_level=B2
Authorization: Bearer <token>

Response 200:
{
  "success": true,
  "data": {
    "lists": [
      {
        "id": 1,
        "name": "TOEIC核心詞彙",
        "description": "TOEIC考試必備的核心詞彙清單",
        "target_cefr_level": "B2",
        "is_public": false,
        "word_count": 25,
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 1,
      "total_pages": 1
    }
  }
}
```

**3. 清單詳細資訊**
```
GET /api/v1/lists/:id
Authorization: Bearer <token>

Response 200:
{
  "success": true,
  "data": {
    "list": {
      "id": 1,
      "name": "TOEIC核心詞彙",
      "description": "TOEIC考試必備的核心詞彙清單",
      "target_cefr_level": "B2",
      "is_public": false,
      "word_count": 25,
      "words": [
        {
          "id": 1,
          "word": "sophisticated",
          "cefr_level": "C1",
          "position": 1,
          "added_at": "2025-01-01T00:00:00Z"
        }
      ],
      "created_at": "2025-01-01T00:00:00Z"
    }
  }
}
```

#### 單字管理 API

**1. AI輔助新增單字**
```
POST /api/v1/lists/:id/words/ai-assist
Authorization: Bearer <token>
Content-Type: application/json

{
  "word": "sophisticated",
  "user_cefr_level": "B2"
}

Response 201:
{
  "success": true,
  "message": "單字資訊已由AI生成",
  "data": {
    "word": {
      "id": 1,
      "word": "sophisticated",
      "phonetic": "/səˈfɪstɪkeɪtɪd/",
      "cefr_level": "C1",
      "definitions": {
        "B1": "很聰明或複雜的",
        "B2": "精密複雜的；老練世故的",
        "C1": "精密複雜的；世故的；詭辯的"
      },
      "examples": {
        "B2": "This is a sophisticated system.",
        "C1": "She has a sophisticated understanding of quantum physics."
      },
      "synonyms": ["complex", "advanced", "refined"],
      "antonyms": ["simple", "basic", "naive"],
      "memory_tips": "記住 'sophisticated' = soph(智慧) + istic + ated，意思是充滿智慧的、複雜的"
    }
  }
}
```

**2. 手動新增單字**
```
POST /api/v1/lists/:id/words
Authorization: Bearer <token>
Content-Type: application/json

{
  "word": "excellent",
  "phonetic": "/ˈeksələnt/",
  "cefr_level": "B1",
  "definitions": {
    "A2": "非常好的",
    "B1": "卓越的，極好的"
  },
  "examples": {
    "A2": "This is excellent!",
    "B1": "She did an excellent job on the project."
  },
  "synonyms": ["outstanding", "superb"],
  "antonyms": ["poor", "terrible"]
}
```

#### 學習會話 API

**1. 開始學習會話**
```
POST /api/v1/learning/sessions
Authorization: Bearer <token>
Content-Type: application/json

{
  "list_id": 1,
  "word_count": 10,
  "shuffle": true,
  "review_mode": false
}

Response 201:
{
  "success": true,
  "data": {
    "session": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "list_id": 1,
      "word_count": 10,
      "current_position": 0,
      "words": [
        {
          "id": 1,
          "word": "sophisticated",
          "show_answer": false
        }
      ],
      "started_at": "2025-01-01T00:00:00Z"
    }
  }
}
```

**2. 提交學習進度**
```
POST /api/v1/learning/sessions/:id/progress
Authorization: Bearer <token>
Content-Type: application/json

{
  "word_id": 1,
  "is_correct": true,
  "difficulty_rating": 3,
  "response_time_ms": 5000
}

Response 200:
{
  "success": true,
  "data": {
    "next_word": {
      "id": 2,
      "word": "excellent",
      "show_answer": false
    },
    "progress": {
      "completed": 1,
      "total": 10,
      "percentage": 10
    }
  }
}
```

#### 複習系統 API

**1. 取得今日複習**
```
GET /api/v1/learning/reviews/today
Authorization: Bearer <token>

Response 200:
{
  "success": true,
  "data": {
    "reviews": [
      {
        "word_id": 1,
        "word": "sophisticated",
        "list_name": "TOEIC核心詞彙",
        "repetition_count": 2,
        "last_reviewed": "2025-01-01T00:00:00Z"
      }
    ],
    "total_count": 1
  }
}
```

**2. 更新複習排程**
```
PUT /api/v1/learning/reviews/:word_id
Authorization: Bearer <token>
Content-Type: application/json

{
  "is_correct": true,
  "difficulty_rating": 4
}

Response 200:
{
  "success": true,
  "data": {
    "schedule": {
      "next_review_date": "2025-01-08",
      "current_interval_days": 7,
      "repetition_count": 3
    }
  }
}
```

### Go 資料模型定義

```go
// pkg/models/wordlist.go
package models

import (
    "time"
    "database/sql/driver"
    "encoding/json"
)

type WordList struct {
    ID              int       `json:"id" db:"id"`
    UserID          int       `json:"user_id" db:"user_id"`
    Name            string    `json:"name" db:"name"`
    Description     *string   `json:"description" db:"description"`
    TargetCEFRLevel string    `json:"target_cefr_level" db:"target_cefr_level"`
    IsPublic        bool      `json:"is_public" db:"is_public"`
    IsActive        bool      `json:"is_active" db:"is_active"`
    WordCount       int       `json:"word_count"`
    CreatedAt       time.Time `json:"created_at" db:"created_at"`
    UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Word struct {
    ID           int             `json:"id" db:"id"`
    Word         string          `json:"word" db:"word"`
    Phonetic     *string         `json:"phonetic" db:"phonetic"`
    CEFRLevel    string          `json:"cefr_level" db:"cefr_level"`
    Definitions  DefinitionMap   `json:"definitions" db:"definitions"`
    Examples     ExampleMap      `json:"examples" db:"examples"`
    Synonyms     []string        `json:"synonyms" db:"synonyms"`
    Antonyms     []string        `json:"antonyms" db:"antonyms"`
    RelatedWords []string        `json:"related_words" db:"related_words"`
    MemoryTips   *string         `json:"memory_tips" db:"memory_tips"`
    CreatedAt    time.Time       `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
}

type DefinitionMap map[string]string
type ExampleMap map[string]string

// 實作 database/sql driver.Valuer 介面
func (dm DefinitionMap) Value() (driver.Value, error) {
    return json.Marshal(dm)
}

func (dm *DefinitionMap) Scan(value interface{}) error {
    if value == nil {
        *dm = make(DefinitionMap)
        return nil
    }
    
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    
    return json.Unmarshal(bytes, dm)
}

type LearningRecord struct {
    ID               int       `json:"id" db:"id"`
    UserID           int       `json:"user_id" db:"user_id"`
    WordID           int       `json:"word_id" db:"word_id"`
    SessionID        string    `json:"session_id" db:"session_id"`
    IsCorrect        bool      `json:"is_correct" db:"is_correct"`
    DifficultyRating int       `json:"difficulty_rating" db:"difficulty_rating"`
    ResponseTimeMs   int       `json:"response_time_ms" db:"response_time_ms"`
    LearnedAt        time.Time `json:"learned_at" db:"learned_at"`
}

type ReviewSchedule struct {
    ID                int       `json:"id" db:"id"`
    UserID            int       `json:"user_id" db:"user_id"`
    WordID            int       `json:"word_id" db:"word_id"`
    CurrentIntervalDays int     `json:"current_interval_days" db:"current_interval_days"`
    EaseFactor        float32   `json:"ease_factor" db:"ease_factor"`
    RepetitionCount   int       `json:"repetition_count" db:"repetition_count"`
    NextReviewDate    time.Time `json:"next_review_date" db:"next_review_date"`
    IsGraduated       bool      `json:"is_graduated" db:"is_graduated"`
    CreatedAt         time.Time `json:"created_at" db:"created_at"`
    UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// 請求模型
type CreateWordListRequest struct {
    Name            string  `json:"name" binding:"required,min=1,max=100"`
    Description     *string `json:"description" binding:"omitempty,max=500"`
    TargetCEFRLevel string  `json:"target_cefr_level" binding:"required,oneof=A1 A2 B1 B2 C1 C2"`
    IsPublic        bool    `json:"is_public"`
}

type AddWordAIRequest struct {
    Word          string `json:"word" binding:"required,min=1,max=100"`
    UserCEFRLevel string `json:"user_cefr_level" binding:"required,oneof=A1 A2 B1 B2 C1 C2"`
}

type StartLearningSessionRequest struct {
    ListID     int  `json:"list_id" binding:"required"`
    WordCount  int  `json:"word_count" binding:"required,min=1,max=50"`
    Shuffle    bool `json:"shuffle"`
    ReviewMode bool `json:"review_mode"`
}

type SubmitProgressRequest struct {
    WordID           int `json:"word_id" binding:"required"`
    IsCorrect        bool `json:"is_correct"`
    DifficultyRating int `json:"difficulty_rating" binding:"required,min=1,max=5"`
    ResponseTimeMs   int `json:"response_time_ms" binding:"min=0"`
}
```

### AI 整合規格

#### Claude API 客戶端設計

```go
// pkg/services/ai_service.go
package services

import (
    "context"
    "fmt"
    "encoding/json"
)

type AIService interface {
    GenerateWordInfo(ctx context.Context, word string, userCEFRLevel string) (*WordInfo, error)
    GenerateMemoryTip(ctx context.Context, word string, definition string) (string, error)
}

type ClaudeAIService struct {
    client    *anthropic.Client
    model     string
    maxTokens int
}

type WordInfo struct {
    Word         string            `json:"word"`
    Phonetic     string            `json:"phonetic"`
    CEFRLevel    string            `json:"cefr_level"`
    Definitions  map[string]string `json:"definitions"`
    Examples     map[string]string `json:"examples"`
    Synonyms     []string          `json:"synonyms"`
    Antonyms     []string          `json:"antonyms"`
    MemoryTips   string            `json:"memory_tips"`
}

func (s *ClaudeAIService) GenerateWordInfo(ctx context.Context, word string, userCEFRLevel string) (*WordInfo, error) {
    prompt := s.buildWordInfoPrompt(word, userCEFRLevel)
    
    response, err := s.client.Messages.New(ctx, anthropic.MessageNewParams{
        Model:     anthropic.String(s.model),
        MaxTokens: anthropic.Int(s.maxTokens),
        Messages: []anthropic.MessageParam{
            anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
        },
    })
    
    if err != nil {
        return nil, fmt.Errorf("Claude API call failed: %w", err)
    }
    
    var wordInfo WordInfo
    if err := json.Unmarshal([]byte(response.Content[0].Text), &wordInfo); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }
    
    return &wordInfo, nil
}

func (s *ClaudeAIService) buildWordInfoPrompt(word string, userCEFRLevel string) string {
    return fmt.Sprintf(`
你是一個專業的英語教學助手。請為單字 "%s" 生成適合 %s 等級學習者的完整資訊。

請以JSON格式回應，包含以下欄位：
{
  "word": "%s",
  "phonetic": "音標",
  "cefr_level": "評估此單字的CEFR等級",
  "definitions": {
    "A1": "如果適用於A1等級的簡單定義",
    "A2": "如果適用於A2等級的基礎定義",
    "B1": "如果適用於B1等級的中級定義",
    "B2": "如果適用於B2等級的中高級定義",
    "C1": "如果適用於C1等級的高級定義",
    "C2": "如果適用於C2等級的精通定義"
  },
  "examples": {
    "相應等級": "適合該等級的例句"
  },
  "synonyms": ["同義詞列表，適合用戶等級"],
  "antonyms": ["反義詞列表"],
  "memory_tips": "記憶技巧建議"
}

請根據使用者的 %s 等級調整內容複雜度，確保定義和例句適合其理解程度。
`, word, userCEFRLevel, word, userCEFRLevel)
}
```

### 艾賓浩斯記憶曲線演算法

```go
// pkg/services/spaced_repetition.go
package services

import (
    "math"
    "time"
)

type SpacedRepetitionService struct{}

type ReviewResult struct {
    NextReviewDate      time.Time
    NewIntervalDays     int
    NewEaseFactor       float32
    NewRepetitionCount  int
    IsGraduated         bool
}

// 根據SM-2演算法計算下次複習時間
func (s *SpacedRepetitionService) CalculateNextReview(
    isCorrect bool,
    difficultyRating int, // 1-5, 5是最簡單
    currentInterval int,
    easeFactor float32,
    repetitionCount int,
) ReviewResult {
    newRepetitionCount := repetitionCount
    newEaseFactor := easeFactor
    newInterval := currentInterval
    
    if isCorrect {
        newRepetitionCount++
        
        // 更新容易度因子
        newEaseFactor = s.updateEaseFactor(easeFactor, difficultyRating)
        
        // 計算新的間隔
        switch newRepetitionCount {
        case 1:
            newInterval = 1
        case 2:
            newInterval = 6
        default:
            newInterval = int(math.Round(float64(currentInterval) * float64(newEaseFactor)))
        }
        
        // 確保最小間隔為1天
        if newInterval < 1 {
            newInterval = 1
        }
        
    } else {
        // 答錯時重置
        newRepetitionCount = 0
        newInterval = 1
    }
    
    nextReviewDate := time.Now().AddDate(0, 0, newInterval)
    
    // 判斷是否已畢業（連續答對且間隔超過60天）
    isGraduated := newRepetitionCount >= 5 && newInterval >= 60
    
    return ReviewResult{
        NextReviewDate:     nextReviewDate,
        NewIntervalDays:    newInterval,
        NewEaseFactor:      newEaseFactor,
        NewRepetitionCount: newRepetitionCount,
        IsGraduated:        isGraduated,
    }
}

// 更新容易度因子 (SM-2算法)
func (s *SpacedRepetitionService) updateEaseFactor(currentEF float32, quality int) float32 {
    // quality: 5=完美, 4=正確且容易, 3=正確但困難, 2=錯誤但記得, 1=錯誤且不記得
    // 轉換為0-5的q值
    q := float32(quality)
    
    newEF := currentEF + (0.1 - (5-q)*(0.08+(5-q)*0.02))
    
    // 確保EF在合理範圍內
    if newEF < 1.3 {
        newEF = 1.3
    }
    
    return newEF
}

// 取得今日需要複習的單字
func (s *SpacedRepetitionService) GetTodaysReviews(userID int) ([]ReviewSchedule, error) {
    today := time.Now().Format("2006-01-02")
    
    // 這裡會呼叫repository查詢數據
    // 實際實作在repository層
    return nil, nil
}
```

## 🧪 測試策略

### 單元測試

**測試覆蓋率目標**: > 80%

**重點測試模組**:
1. **AI服務測試**
   - Claude API整合測試
   - 提示工程驗證
   - 錯誤處理測試

2. **記憶曲線演算法測試**
   - SM-2演算法正確性
   - 邊界條件測試
   - 複習排程生成

3. **資料模型測試**
   - 資料驗證邏輯
   - JSON序列化/反序列化
   - 資料庫映射

### 整合測試

**API端點測試**:
```go
// tests/integration/wordlist_test.go
func TestWordListAPI(t *testing.T) {
    // 測試完整的清單管理流程
    // 1. 建立清單
    // 2. 新增單字
    // 3. 查詢清單
    // 4. 更新清單
    // 5. 刪除清單
}

func TestLearningSessionAPI(t *testing.T) {
    // 測試完整的學習流程
    // 1. 開始學習會話
    // 2. 提交學習進度
    // 3. 完成會話
    // 4. 檢查複習排程
}
```

**資料庫測試**:
- 使用測試資料庫
- 交易回滾機制
- 資料一致性驗證

### 效能測試

**負載測試指標**:
- API響應時間 < 200ms
- 並發用戶支援 > 100
- AI API呼叫優化

**測試工具**:
- Go benchmark測試
- Apache JMeter負載測試
- 資料庫查詢效能分析

## 🚀 部署與監控

### 部署檢查清單

**環境變數設定**:
```bash
# 新增的環境變數
CLAUDE_API_KEY=your_claude_api_key
CLAUDE_MODEL=claude-3-haiku-20240307
MAX_AI_TOKENS=1000
REVIEW_REMINDER_ENABLED=true
DEFAULT_CEFR_LEVEL=A1
```

**資料庫遷移**:
```bash
# 執行遷移腳本
make migrate-up

# 驗證資料遷移
make migrate-status
```

**API測試**:
```bash
# 健康檢查
curl http://localhost:8080/health

# API功能測試
curl -X POST http://localhost:8080/api/v1/lists \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "測試清單", "target_cefr_level": "B1"}'
```

### 監控指標

**業務指標**:
- 每日活躍用戶數
- 單字學習完成率
- AI API呼叫成功率
- 複習準時率

**技術指標**:
- API響應時間
- 資料庫連接池狀態
- AI API延遲時間
- 錯誤率統計

**告警設定**:
- API錯誤率 > 5%
- 資料庫連接失敗
- AI API調用失敗
- 磁碟空間不足

## 📚 使用者手冊

### API使用範例

**完整學習流程範例**:
```javascript
// 1. 建立清單
const createList = async () => {
  const response = await fetch('/api/v1/lists', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      name: 'TOEIC詞彙',
      target_cefr_level: 'B2'
    })
  });
  return response.json();
};

// 2. AI輔助新增單字
const addWordWithAI = async (listId, word) => {
  const response = await fetch(`/api/v1/lists/${listId}/words/ai-assist`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      word: word,
      user_cefr_level: 'B2'
    })
  });
  return response.json();
};

// 3. 開始學習會話
const startLearning = async (listId) => {
  const response = await fetch('/api/v1/learning/sessions', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      list_id: listId,
      word_count: 10,
      shuffle: true
    })
  });
  return response.json();
};

// 4. 提交學習結果
const submitProgress = async (sessionId, wordId, isCorrect) => {
  const response = await fetch(`/api/v1/learning/sessions/${sessionId}/progress`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      word_id: wordId,
      is_correct: isCorrect,
      difficulty_rating: 3,
      response_time_ms: 5000
    })
  });
  return response.json();
};
```

### 常見問題解答

**Q: AI生成的單字資訊不準確怎麼辦？**
A: 系統支援手動編輯AI生成的內容，使用PUT API更新單字資訊。

**Q: 如何調整複習頻率？**
A: 系統會根據答題表現自動調整，也可以通過difficulty_rating參數微調。

**Q: 支援匯入現有的單字清單嗎？**
A: 目前支援逐個新增，未來版本將支援CSV批量匯入功能。

**Q: CEFR等級如何判定？**
A: AI會自動評估單字的CEFR等級，也可以手動調整。

## 📈 後續擴展計劃

### Phase 1 延伸功能
- 語音朗讀整合
- 批量匯入/匯出
- 清單分享功能
- 學習排行榜

### Phase 2 進階功能
- 個人化推薦系統
- 學習社群功能
- 多語言支援
- 離線學習模式

### Phase 3 智能化升級
- 自適應學習算法
- 語音識別練習
- 圖像關聯記憶
- VR/AR學習體驗

---

## 🔧 開發環境設定

### 必要軟體安裝
```bash
# 安裝 golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 安裝 Air 熱重載工具
go install github.com/cosmtrek/air@latest

# 安裝 golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 開發流程
1. 建立feature分支
2. 實作功能與測試
3. 運行完整測試套件
4. 提交PR並進行代碼審查
5. 合併到develop分支
6. 部署到測試環境驗證

### 代碼品質檢查
```bash
# 運行linter
make lint

# 運行測試
make test

# 檢查覆蓋率
make coverage

# 格式化代碼
go fmt ./...
```

此規格書將作為整個記憶單字功能開發的指南，確保功能的完整性和一致性。所有開發人員應該遵循此規格進行實作。