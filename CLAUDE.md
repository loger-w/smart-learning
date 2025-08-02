# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Smart Learning is an intelligent English learning platform combining modern frontend technology, backend API services, and AI integration. The project uses a monorepo structure with separate frontend (React) and backend (Go) applications.

## Architecture

### Technology Stack

**Frontend**:
- Vite + React + TypeScript
- TailwindCSS + Shadcn UI
- TanStack Query (server state) + Zustand (client state)
- TanStack Router (type-safe routing)
- React Hook Form
- Vitest + React Testing Library

**Backend**:
- Go + Gin framework
- PostgreSQL (Supabase hosted)
- JWT authentication
- Claude Haiku API for AI features
- Go Testing + Testify

### Project Structure

```
smart-learning/
├── frontend/          # React frontend application
├── backend/           # Go backend API server
├── docs/             # Comprehensive project documentation
│   ├── 01-project-initialization.md
│   ├── 02-technical-architecture.md
│   ├── 03-development-workflow.md
│   ├── 04-feature-modules.md
│   ├── 05-deployment-guide.md
│   └── 06-ui-design-system.md
└── *.md files        # Project specifications and requirements
```

## Development Commands

### Frontend Development
```bash
cd frontend
npm install          # Install dependencies
npm run dev          # Start development server
npm run build        # Build for production
npm run lint         # Run ESLint
npm run type-check   # TypeScript type checking
npm run test         # Run tests
npm run test:watch   # Run tests in watch mode
npm run test:coverage # Generate coverage report
```

### Backend Development
```bash
cd backend
go mod tidy          # Install/update dependencies
go run cmd/main.go   # Start development server
go test ./...        # Run all tests
go test -v ./...     # Run tests with verbose output
go test -cover ./... # Run tests with coverage
```

### Database Operations
```bash
# Create migration
migrate create -ext sql -dir migrations -seq migration_name

# Apply migrations
migrate -path migrations -database "postgresql://localhost/smart_learning_db?sslmode=disable" up

# Rollback migrations
migrate -path migrations -database "postgresql://localhost/smart_learning_db?sslmode=disable" down 1
```

## Core Features

1. **User Authentication** - Registration, login, JWT-based auth
2. **Word List Management** - Create, edit, share learning lists
3. **Flashcard Learning** - Interactive card-based learning with progress tracking
4. **AI Integration** - Claude AI for explanations, synonyms, and memory techniques
5. **Level-based Learning** - Content difficulty adjusted to user level
6. **Learning Analytics** - Progress tracking and performance analysis
7. **Search & Filtering** - Advanced list search and filtering capabilities

## Key Architecture Patterns

### Frontend State Management
- **Server State**: TanStack Query for API data caching and synchronization
- **Client State**: Zustand for local application state
- **Forms**: React Hook Form for performance-optimized form handling

### Backend Architecture
- **Layered Architecture**: Handler → Service → Repository pattern
- **Dependency Injection**: Interfaces for testability and modularity
- **Error Handling**: Custom error types with proper HTTP status mapping

### Database Design
- **Core Tables**: users, word_lists, words, learning_records
- **Relationships**: Proper foreign key constraints and cascading deletes
- **Indexing**: Optimized for common query patterns

## Development Standards

### Git Workflow
- **Branch Strategy**: Git Flow (main/develop/feature/release/hotfix)
- **Commit Format**: Conventional Commits (`feat:`, `fix:`, `docs:`, etc.)
- **Branch Naming**: `feature/feature-name`, `bugfix/issue-description`

### Code Quality
- **TypeScript**: Strict type checking, proper interface definitions
- **Go**: Follow Go conventions, proper error handling, interface-based design
- **Testing**: Unit tests for business logic, integration tests for API endpoints
- **Linting**: ESLint for TypeScript, golangci-lint for Go

## Environment Setup

### Prerequisites
- Node.js >= 18
- Go >= 1.21
- PostgreSQL (or Supabase account)
- Claude API key for AI features

### Quick Setup
For detailed environment setup instructions, please refer to **[ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)**

### Environment Files
```bash
# Frontend
cp frontend/.env.example frontend/.env.local

# Backend
cp backend/.env.example backend/.env
```

## Testing Strategy

### Frontend Testing
- **Unit Tests**: Component logic and utility functions
- **Integration Tests**: Hook behavior and API interactions
- **E2E Tests**: Critical user journeys (if implemented)

### Backend Testing
- **Unit Tests**: Service layer business logic
- **Integration Tests**: API endpoints with test database
- **Repository Tests**: Database operations

## Deployment

### Frontend
- **Platform**: Vercel (recommended)
- **Build Command**: `npm run build`
- **Output Directory**: `dist`

### Backend
- **Platform**: Railway (recommended)
- **Database**: Supabase PostgreSQL
- **Environment**: Docker containerization

## Important Notes

- **Current Status**: Project is in planning phase with comprehensive documentation but minimal code implementation
- **AI Integration**: Uses Claude Haiku API for generating explanations, memory techniques, and learning content
- **Internationalization**: Primarily focused on English learning for Chinese speakers
- **Security**: JWT-based authentication, input validation, SQL injection prevention
- **Performance**: Query optimization, caching strategies, code splitting planned

## Documentation References

For detailed information, refer to the comprehensive documentation in the `docs/` directory:
- Technical architecture details in `02-technical-architecture.md`
- Development workflow in `03-development-workflow.md`
- Feature specifications in `04-feature-modules.md`
- Deployment procedures in `05-deployment-guide.md`

## custum

You are a Senior Front-End Developer and an Expert in ReactJS, JavaScript, TypeScript, HTML, CSS and modern UI/UX frameworks (e.g., TailwindCSS, Shadcn). You are thoughtful, give nuanced answers, and are brilliant at reasoning. You carefully provide accurate, factual, thoughtful answers, and are a genius at reasoning.

- Follow the user’s requirements carefully & to the letter.
- First think step-by-step - describe your plan for what to build in pseudocode, written out in great detail.
- Confirm, then write code!
- Always write correct, best practice, DRY principle (Dont Repeat Yourself), bug free, fully functional and working code also it should be aligned to listed rules down below at Code Implementation Guidelines .
- Focus on easy and readability code, over being performant.
- Fully implement all requested functionality.
- Leave NO todo’s, placeholders or missing pieces.
- Ensure code is complete! Verify thoroughly finalised.
- Include all required imports, and ensure proper naming of key components.
- Be concise Minimize any other prose.
- If you think there might not be a correct answer, you say so.
- If you do not know the answer, say so, instead of guessing.

### Coding Environment
The user asks questions about the following coding languages:
- ReactJS
- JavaScript
- TypeScript
- TailwindCSS
- HTML
- CSS
- TankStack

### Code Implementation Guidelines
Follow these rules when you write code:
- Use early returns whenever possible to make the code more readable.
- Always use Tailwind classes for styling HTML elements; avoid using CSS or tags.
- Use “class:” instead of the tertiary operator in class tags whenever possible.
- Use descriptive variable and function/const names. Also, event functions should be named with a “handle” prefix, like “handleClick” for onClick and “handleKeyDown” for onKeyDown.
- Implement accessibility features on elements. For example, a tag should have a tabindex=“0”, aria-label, on:click, and on:keydown, and similar attributes.
- Use consts instead of functions, for example, “const toggle = () =>”. Also, define a type if possible.
生成 console.log 時，記得 JSON.stringify()
改 code 時，只改我這次跟你討論的部分，不要改無關的 code
只有我叫你實作的時候才要改 code，不然不要改任何的 code，只先提出討論即可
盡量使用 pure function
在做 value 的 fallback 的時候，請用 ?? 而非 ||
如果是請你做 commit 資訊、上版的問題時，格式請統一並使用正規的方法，例如: feat(fix) ...
若解決完使用者的問題，請簡述問題、把有問題的程式提出並闡述問題點、為甚麼發生以及怎麼用了甚麼方法解決，若之前有介紹過就不用再介紹一次
不要刪除註解的程式碼
思考過程跟回答總是為繁體中文
如果是討論的時候就不要修改任何程式碼 例如我會問你你認為呢或是你有任何想法嗎之類的
盡量使用 early return