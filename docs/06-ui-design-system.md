# Smart Learning UI 設計系統

## 設計理念

Smart Learning 的 UI 設計遵循「簡潔、直觀、高效」的核心理念，為英語學習者提供無干擾的學習體驗。

### 設計原則

1. **簡潔明確** - 每個介面元素都有明確的目的
2. **功能導向** - 設計服務於學習功能，而非炫技
3. **認知負荷最小化** - 減少使用者的心理負擔
4. **漸進式披露** - 避免資訊過載
5. **一致性** - 統一的視覺語言和互動模式

## 色彩系統

### 主色調
```css
:root {
  /* 主品牌色 - 智慧藍 */
  --primary-50: #eff6ff;
  --primary-100: #dbeafe;
  --primary-500: #3b82f6;
  --primary-600: #2563eb;
  --primary-700: #1d4ed8;
  
  /* 輔助色 - 學習綠 */
  --success-50: #f0fdf4;
  --success-100: #dcfce7;
  --success-500: #22c55e;
  --success-600: #16a34a;
  
  /* 警示色 - 注意橙 */
  --warning-50: #fffbeb;
  --warning-100: #fef3c7;
  --warning-500: #f59e0b;
  --warning-600: #d97706;
  
  /* 錯誤色 - 錯誤紅 */
  --error-50: #fef2f2;
  --error-100: #fee2e2;
  --error-500: #ef4444;
  --error-600: #dc2626;
}
```

### 中性色調
```css
:root {
  /* 灰階系統 */
  --gray-50: #f9fafb;
  --gray-100: #f3f4f6;
  --gray-200: #e5e7eb;
  --gray-300: #d1d5db;
  --gray-400: #9ca3af;
  --gray-500: #6b7280;
  --gray-600: #4b5563;
  --gray-700: #374151;
  --gray-800: #1f2937;
  --gray-900: #111827;
  
  /* 文字色彩 */
  --text-primary: var(--gray-900);
  --text-secondary: var(--gray-600);
  --text-tertiary: var(--gray-400);
  --text-inverse: #ffffff;
}
```

### 色彩使用指南

```typescript
// 色彩使用範例
const colorUsage = {
  primary: {
    main: '主要操作按鈕、連結、進度指示',
    light: '背景色、淺色容器',
    dark: 'hover 狀態、強調文字'
  },
  success: {
    main: '正確答案、完成狀態、成功提示',
    usage: '學習進度、成就徽章'
  },
  warning: {
    main: '需要注意的內容、警告提示',
    usage: '復習提醒、學習建議'
  },
  error: {
    main: '錯誤答案、錯誤提示、刪除操作',
    usage: '表單驗證、錯誤狀態'
  }
}
```

## 字體系統

### 字體族
```css
:root {
  /* 英文字體 */
  --font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  --font-mono: 'JetBrains Mono', 'Fira Code', 'Cascadia Code', monospace;
  
  /* 中文字體 */
  --font-zh: 'Noto Sans TC', 'PingFang TC', 'Microsoft JhengHei', sans-serif;
  
  /* 組合字體 */
  --font-primary: var(--font-sans), var(--font-zh);
}
```

### 字體尺寸
```css
:root {
  /* 字體大小系統 */
  --text-xs: 0.75rem;    /* 12px - 註解文字 */
  --text-sm: 0.875rem;   /* 14px - 次要文字 */
  --text-base: 1rem;     /* 16px - 基礎文字 */
  --text-lg: 1.125rem;   /* 18px - 重要文字 */
  --text-xl: 1.25rem;    /* 20px - 小標題 */
  --text-2xl: 1.5rem;    /* 24px - 中標題 */
  --text-3xl: 1.875rem;  /* 30px - 大標題 */
  --text-4xl: 2.25rem;   /* 36px - 主標題 */
  
  /* 行高 */
  --leading-tight: 1.25;
  --leading-normal: 1.5;
  --leading-relaxed: 1.75;
}
```

### 字體使用規範
```typescript
// 字體使用指南
const typography = {
  heading: {
    h1: 'text-4xl font-bold leading-tight',
    h2: 'text-3xl font-semibold leading-tight', 
    h3: 'text-2xl font-semibold leading-normal',
    h4: 'text-xl font-medium leading-normal'
  },
  body: {
    large: 'text-lg leading-relaxed',
    normal: 'text-base leading-normal',
    small: 'text-sm leading-normal'
  },
  ui: {
    button: 'text-base font-medium',
    input: 'text-base leading-normal',
    caption: 'text-xs leading-normal'
  }
}
```

## 間距系統

### 間距標準
```css
:root {
  /* 間距系統 (基於 4px grid) */
  --space-0: 0;
  --space-1: 0.25rem;  /* 4px */
  --space-2: 0.5rem;   /* 8px */
  --space-3: 0.75rem;  /* 12px */
  --space-4: 1rem;     /* 16px */
  --space-5: 1.25rem;  /* 20px */
  --space-6: 1.5rem;   /* 24px */
  --space-8: 2rem;     /* 32px */
  --space-10: 2.5rem;  /* 40px */
  --space-12: 3rem;    /* 48px */
  --space-16: 4rem;    /* 64px */
  --space-20: 5rem;    /* 80px */
}
```

### 間距使用指南
```typescript
const spacingGuide = {
  component: {
    padding: 'p-4 md:p-6',      // 組件內邊距
    margin: 'mb-4 md:mb-6',     // 組件間距
    gap: 'gap-4'                // 網格間距
  },
  layout: {
    container: 'px-4 md:px-6 lg:px-8',
    section: 'py-8 md:py-12 lg:py-16',
    card: 'p-6 md:p-8'
  },
  text: {
    paragraph: 'mb-4',
    heading: 'mb-6',
    list: 'space-y-2'
  }
}
```

## 組件設計規範

### 按鈕系統

#### 主要按鈕
```typescript
// Primary Button Component
const PrimaryButton = {
  base: 'inline-flex items-center justify-center rounded-lg font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2',
  variants: {
    size: {
      sm: 'px-3 py-2 text-sm',
      md: 'px-4 py-2 text-base',
      lg: 'px-6 py-3 text-lg'
    },
    style: {
      primary: 'bg-primary-600 text-white hover:bg-primary-700 focus:ring-primary-500',
      secondary: 'bg-gray-100 text-gray-900 hover:bg-gray-200 focus:ring-gray-500',
      success: 'bg-success-600 text-white hover:bg-success-700 focus:ring-success-500',
      error: 'bg-error-600 text-white hover:bg-error-700 focus:ring-error-500'
    }
  }
}
```

#### 按鈕實作範例
```tsx
// components/ui/Button.tsx
interface ButtonProps {
  variant?: 'primary' | 'secondary' | 'success' | 'error'
  size?: 'sm' | 'md' | 'lg'
  isLoading?: boolean
  children: React.ReactNode
  onClick?: () => void
}

export const Button: React.FC<ButtonProps> = ({
  variant = 'primary',
  size = 'md',
  isLoading = false,
  children,
  onClick,
  ...props
}) => {
  const baseClasses = 'inline-flex items-center justify-center rounded-lg font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed'
  
  const variants = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500',
    secondary: 'bg-gray-100 text-gray-900 hover:bg-gray-200 focus:ring-gray-500',
    success: 'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500',
    error: 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500'
  }
  
  const sizes = {
    sm: 'px-3 py-2 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg'
  }
  
  const classes = cn(baseClasses, variants[variant], sizes[size])
  
  return (
    <button
      className={classes}
      onClick={onClick}
      disabled={isLoading}
      {...props}
    >
      {isLoading && (
        <svg className="animate-spin -ml-1 mr-2 h-4 w-4" fill="none" viewBox="0 0 24 24">
          <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
          <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
        </svg>
      )}
      {children}
    </button>
  )
}
```

### 卡片系統

#### 卡片設計規範
```typescript
const CardDesign = {
  base: 'bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden',
  variants: {
    default: 'hover:shadow-md transition-shadow',
    interactive: 'hover:shadow-lg hover:border-primary-200 cursor-pointer transition-all',
    highlighted: 'border-primary-200 bg-primary-50'
  },
  padding: {
    sm: 'p-4',
    md: 'p-6', 
    lg: 'p-8'
  }
}
```

#### 單字卡片組件
```tsx
// components/wordlist/WordCard.tsx
interface WordCardProps {
  word: Word
  isSelected?: boolean
  onSelect?: (word: Word) => void
  showProgress?: boolean
}

export const WordCard: React.FC<WordCardProps> = ({
  word,
  isSelected = false,
  onSelect,
  showProgress = true
}) => {
  const accuracy = word.correctCount / (word.correctCount + word.incorrectCount || 1)
  
  return (
    <div
      className={cn(
        'bg-white rounded-xl shadow-sm border transition-all cursor-pointer',
        isSelected 
          ? 'border-blue-200 bg-blue-50 shadow-md' 
          : 'border-gray-200 hover:shadow-md hover:border-gray-300'
      )}
      onClick={() => onSelect?.(word)}
    >
      <div className="p-6">
        <div className="flex justify-between items-start mb-4">
          <div>
            <h3 className="text-xl font-semibold text-gray-900 mb-1">
              {word.word}
            </h3>
            {word.pronunciation && (
              <p className="text-sm text-gray-500">
                [{word.pronunciation}]
              </p>
            )}
          </div>
          
          {showProgress && (
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center">
                <span className={cn(
                  'text-xs font-medium',
                  accuracy >= 0.8 ? 'text-green-600' :
                  accuracy >= 0.6 ? 'text-yellow-600' : 'text-red-600'
                )}>
                  {Math.round(accuracy * 100)}%
                </span>
              </div>
            </div>
          )}
        </div>
        
        <p className="text-gray-700 mb-4 line-clamp-2">
          {word.definition}
        </p>
        
        {word.exampleSentence && (
          <p className="text-sm text-gray-600 italic line-clamp-2">
            "{word.exampleSentence}"
          </p>
        )}
        
        <div className="flex justify-between items-center mt-4">
          <div className="flex items-center space-x-2">
            <DifficultyBadge level={word.difficultyLevel} />
            <MasteryBadge level={word.masteryLevel} />
          </div>
          
          <div className="text-xs text-gray-500">
            {word.lastStudied ? 
              `上次學習: ${formatDate(word.lastStudied)}` : 
              '尚未學習'
            }
          </div>
        </div>
      </div>
    </div>
  )
}
```

### 表單系統

#### 輸入框設計
```tsx
// components/ui/Input.tsx
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
  helper?: string
  leftIcon?: React.ReactNode
  rightIcon?: React.ReactNode
}

export const Input: React.FC<InputProps> = ({
  label,
  error,
  helper,
  leftIcon,
  rightIcon,
  className,
  ...props
}) => {
  const baseClasses = 'block w-full rounded-lg border transition-colors focus:outline-none focus:ring-2 focus:ring-offset-1'
  const stateClasses = error 
    ? 'border-red-300 focus:border-red-500 focus:ring-red-500' 
    : 'border-gray-300 focus:border-blue-500 focus:ring-blue-500'
  const paddingClasses = leftIcon ? 'pl-10 pr-4 py-3' : rightIcon ? 'pl-4 pr-10 py-3' : 'px-4 py-3'
  
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-2">
          {label}
        </label>
      )}
      
      <div className="relative">
        {leftIcon && (
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <div className="h-5 w-5 text-gray-400">
              {leftIcon}
            </div>
          </div>
        )}
        
        <input
          className={cn(baseClasses, stateClasses, paddingClasses, className)}
          {...props}
        />
        
        {rightIcon && (
          <div className="absolute inset-y-0 right-0 pr-3 flex items-center">
            <div className="h-5 w-5 text-gray-400">
              {rightIcon}
            </div>
          </div>
        )}
      </div>
      
      {error && (
        <p className="mt-2 text-sm text-red-600 flex items-center">
          <ExclamationCircleIcon className="h-4 w-4 mr-1" />
          {error}
        </p>
      )}
      
      {helper && !error && (
        <p className="mt-2 text-sm text-gray-500">
          {helper}
        </p>
      )}
    </div>
  )
}
```

## 學習介面設計

### 翻卡學習介面

```tsx
// components/learning/FlashCardInterface.tsx
export const FlashCardInterface: React.FC = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
      <div className="w-full max-w-2xl">
        {/* 進度指示器 */}
        <div className="mb-8">
          <div className="flex justify-between items-center mb-2">
            <span className="text-sm font-medium text-gray-600">
              進度: {currentIndex + 1} / {totalWords}
            </span>
            <span className="text-sm font-medium text-gray-600">
              正確率: {Math.round(accuracy * 100)}%
            </span>
          </div>
          <div className="w-full bg-gray-200 rounded-full h-2">
            <div 
              className="bg-blue-600 h-2 rounded-full transition-all duration-300"
              style={{ width: `${(currentIndex / totalWords) * 100}%` }}
            />
          </div>
        </div>
        
        {/* 翻卡容器 */}
        <div className="perspective-1000 mb-8">
          <FlashCard
            word={currentWord}
            isFlipped={isFlipped}
            onFlip={() => setIsFlipped(!isFlipped)}
            onCorrect={handleCorrect}
            onIncorrect={handleIncorrect}
          />
        </div>
        
        {/* 控制按鈕 */}
        <div className="flex justify-center space-x-4">
          <Button
            variant="secondary"
            onClick={handlePrevious}
            disabled={currentIndex === 0}
          >
            上一個
          </Button>
          
          <Button
            variant="primary"
            onClick={handleNext}
            disabled={currentIndex === totalWords - 1}
          >
            下一個
          </Button>
        </div>
      </div>
    </div>
  )
}
```

### AI 回應介面設計

```tsx
// components/ai/AIResponseCard.tsx
export const AIResponseCard: React.FC<{ response: AIResponse }> = ({ response }) => {
  return (
    <div className="bg-gradient-to-br from-purple-50 to-blue-50 rounded-xl p-6 border border-purple-100">
      {/* AI 標識 */}
      <div className="flex items-center mb-4">
        <div className="w-8 h-8 bg-gradient-to-r from-purple-500 to-blue-500 rounded-full flex items-center justify-center">
          <SparklesIcon className="h-4 w-4 text-white" />
        </div>
        <h3 className="ml-3 text-lg font-semibold text-gray-800">AI 智能解釋</h3>
      </div>
      
      {/* 內容區塊 */}
      <div className="space-y-6">
        {/* 定義 */}
        <div>
          <h4 className="font-semibold text-gray-700 mb-2">詳細定義</h4>
          <p className="text-gray-700 leading-relaxed">{response.definition}</p>
        </div>
        
        {/* 同義詞與反義詞 */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <h4 className="font-semibold text-gray-700 mb-3">同義詞</h4>
            <div className="flex flex-wrap gap-2">
              {response.synonyms.map((synonym, index) => (
                <span key={index} className="px-3 py-1 bg-green-100 text-green-800 rounded-full text-sm font-medium">
                  {synonym}
                </span>
              ))}
            </div>
          </div>
          
          <div>
            <h4 className="font-semibold text-gray-700 mb-3">反義詞</h4>
            <div className="flex flex-wrap gap-2">
              {response.antonyms.map((antonym, index) => (
                <span key={index} className="px-3 py-1 bg-red-100 text-red-800 rounded-full text-sm font-medium">
                  {antonym}
                </span>
              ))}
            </div>
          </div>
        </div>
        
        {/* 記憶技巧 */}
        <div>
          <h4 className="font-semibold text-gray-700 mb-3 flex items-center">
            <LightBulbIcon className="h-5 w-5 mr-2 text-yellow-500" />
            記憶技巧
          </h4>
          <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 rounded-r-lg">
            <p className="text-gray-700 italic">{response.memoryTips}</p>
          </div>
        </div>
        
        {/* 例句 */}
        <div>
          <h4 className="font-semibold text-gray-700 mb-3">例句示範</h4>
          <div className="space-y-3">
            {response.examples.map((example, index) => (
              <div key={index} className="bg-white rounded-lg p-4 border border-gray-200">
                <p className="text-gray-800 mb-2">"{example.sentence}"</p>
                <p className="text-gray-600 text-sm">{example.translation}</p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
```

## 響應式設計

### 斷點系統
```css
:root {
  /* 響應式斷點 */
  --breakpoint-sm: 640px;   /* 手機橫向 */
  --breakpoint-md: 768px;   /* 平板直向 */
  --breakpoint-lg: 1024px;  /* 平板橫向/小筆電 */
  --breakpoint-xl: 1280px;  /* 桌面 */
  --breakpoint-2xl: 1536px; /* 大桌面 */
}
```

### 響應式設計原則

#### 行動優先
```typescript
// 行動優先的設計策略
const responsiveDesign = {
  mobile: {
    layout: 'single-column',
    navigation: 'bottom-tab',
    cards: 'full-width',
    spacing: 'compact'
  },
  tablet: {
    layout: 'two-column',
    navigation: 'side-drawer',
    cards: 'grid-2',
    spacing: 'normal'
  },
  desktop: {
    layout: 'three-column',
    navigation: 'top-header',
    cards: 'grid-3',
    spacing: 'spacious'
  }
}
```

#### 響應式組件範例
```tsx
// 響應式 WordList 組件
export const WordListGrid: React.FC = () => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6">
      {wordLists.map(wordList => (
        <WordListCard
          key={wordList.id}
          wordList={wordList}
          className="h-full" // 確保卡片等高
        />
      ))}
    </div>
  )
}
```

## 無障礙設計

### 可訪問性標準
```typescript
// WCAG 2.1 AA 標準實作
const accessibilityStandards = {
  colorContrast: {
    normal: '4.5:1',    // 一般文字
    large: '3:1',       // 大文字
    nonText: '3:1'      // UI 元件
  },
  focusManagement: {
    visible: 'focus:ring-2 focus:ring-blue-500',
    logical: 'tab-index routing',
    skip: 'skip-to-content links'
  },
  semantics: {
    headings: 'hierarchical h1-h6',
    landmarks: 'main, nav, aside, footer',
    labels: 'aria-label, aria-describedby'
  }
}
```

### 無障礙組件實作
```tsx
// 無障礙按鈕組件
export const AccessibleButton: React.FC<ButtonProps> = ({
  children,
  ariaLabel,
  ariaDescribedBy,
  ...props
}) => {
  return (
    <button
      aria-label={ariaLabel}
      aria-describedby={ariaDescribedBy}
      className="focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
      {...props}
    >
      {children}
    </button>
  )
}
```

## 動畫與互動

### 動畫原則
```css
/* 動畫時間標準 */
:root {
  --duration-fast: 150ms;
  --duration-normal: 250ms;
  --duration-slow: 350ms;
  
  /* 緩動函數 */
  --ease-in: cubic-bezier(0.4, 0, 1, 1);
  --ease-out: cubic-bezier(0, 0, 0.2, 1);
  --ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
}

/* 常用動畫類別 */
.transition-smooth {
  transition: all var(--duration-normal) var(--ease-in-out);
}

.fade-in {
  animation: fadeIn var(--duration-normal) var(--ease-out);
}

.slide-up {
  animation: slideUp var(--duration-normal) var(--ease-out);
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideUp {
  from { 
    opacity: 0;
    transform: translateY(20px);
  }
  to { 
    opacity: 1;
    transform: translateY(0);
  }
}
```

### 翻卡動畫實作
```css
/* 3D 翻卡效果 */
.card-container {
  perspective: 1000px;
}

.card {
  position: relative;
  width: 100%;
  height: 300px;
  transform-style: preserve-3d;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.card.flipped {
  transform: rotateY(180deg);
}

.card-front,
.card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  border-radius: 12px;
}

.card-back {
  transform: rotateY(180deg);
}
```

這份 UI 設計系統文檔確保 Smart Learning 擁有一致、美觀且易用的使用者介面，同時考慮了可訪問性和響應式設計需求。