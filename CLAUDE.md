# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Smart Learning is an intelligent English learning platform combining modern frontend technology, backend API services, and AI integration. The project uses a monorepo structure with separate frontend (React) and backend (Go) applications.

## Architecture

### Technology Stack

**Frontend**:
- Vite 7.0.4 + React 19.1.0 + TypeScript 5.8.3
- TailwindCSS 4.1.11 + Shadcn UI 0.9.5
- TanStack Query 5.84.1 (server state) + Zustand 5.0.7 (client state)
- TanStack Router 1.130.12 (type-safe routing)
- React Hook Form 7.62.0 + Zod 4.0.14 (validation)
- Vitest 3.2.4 + React Testing Library 16.3.0

**Backend**:
- Go 1.24.5 + Gin 1.10.1 framework
- PostgreSQL (Supabase hosted) with lib/pq 1.10.9 driver
- JWT authentication
- Claude Haiku API for AI features
- Go Testing framework + Air for hot reload

### Project Structure

```
smart-learning/
├── frontend/          # React frontend application
│   └── src/
│       ├── components/     # Reusable components
│       │   └── ui/        # Shadcn UI components
│       ├── features/      # Feature-based modules
│       │   ├── auth/      # Authentication module
│       │   └── dashboard/ # Dashboard module
│       ├── hooks/         # Custom hooks (grouped by feature)
│       ├── lib/           # Shared utilities
│       ├── providers/     # Context providers
│       ├── routes/        # TanStack Router configuration
│       ├── services/      # API services
│       ├── stores/        # Zustand state management
│       └── types/         # TypeScript type definitions
├── backend/           # Go backend API server
│   ├── cmd/           # Application entrypoints
│   └── pkg/           # Shared packages
├── docs/             # Comprehensive project documentation
└── *.md files        # Project specifications
```

## Development Commands

### Frontend Development
```bash
cd frontend
npm install          # Install dependencies
npm run dev          # Start development server (Vite)
npm run build        # Build for production (TypeScript + Vite)
npm run preview      # Preview production build locally
npm run lint         # Run ESLint
npm run type-check   # TypeScript type checking
npm run test         # Run tests (Vitest)
npm run test:watch   # Run tests in watch mode
npm run test:coverage # Generate coverage report
```

### Backend Development
```bash
cd backend
# Using Go commands directly
go mod tidy          # Install/update dependencies
go run cmd/main.go   # Start development server
go test ./...        # Run all tests
go test -v ./...     # Run tests with verbose output
go test -cover ./... # Run tests with coverage

# Using Makefile (recommended)
make deps            # Install/update dependencies (calls go mod tidy)
make run             # Build and start server
make dev             # Start with air hot reload (requires air installed)
make test            # Run all tests
make coverage        # Run tests with coverage
make lint            # Run golangci-lint (requires golangci-lint installed)
make build           # Build binary
make clean           # Clean build artifacts
make docker-build    # Build Docker image
make docker-run      # Run Docker container
```

### Database Operations
```bash
cd backend
# Using golang-migrate
make migrate-create name=migration_name  # Create new migration
make migrate-up                          # Apply all pending migrations
make migrate-down                        # Rollback last migration
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

## Project Configuration

### Frontend Configuration
- **Path Alias**: `@` maps to `./src` directory for imports
- **Testing**: Vitest with jsdom environment, setup file at `./src/test/setup.ts`
- **ESLint**: Modern flat config with TypeScript, React Hooks, and React Refresh plugins
- **TailwindCSS**: V4 with Vite plugin integration
- **Vite Config**: React plugin + TailwindCSS plugin with path alias support

### Backend Configuration
- **Main Entry**: `cmd/main.go`
- **Internal Structure**: handlers, middleware, models, repositories, services
- **Database Package**: `pkg/database` with PostgreSQL connection
- **Environment**: Uses `.env` file with godotenv for local development

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
- Go >= 1.24.5
- PostgreSQL (or Supabase account)
- Claude API key for AI features
- Air (for Go hot reload): `go install github.com/cosmtrek/air@latest`
- golang-migrate (for database migrations): `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- golangci-lint (for Go linting): `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

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

## Architecture Highlights

### Frontend Architecture
- **Feature-Based Organization**: Code organized by features (`auth/`, `dashboard/`) rather than technical layers
- **Component Co-location**: Components, pages, and hooks grouped within feature modules
- **Type-Safe Routing**: TanStack Router with generated route tree for compile-time safety
- **Modern React**: React 19 with functional components and hooks
- **Path Aliases**: `@` alias configured for clean imports from `src/`

### Backend Architecture
- **Minimal Setup**: Go backend with Gin framework, PostgreSQL database
- **Package Structure**: Clean separation with `cmd/` for entrypoints and `pkg/` for shared code
- **Development Tools**: Makefile for common tasks, Air for hot reload

### Key Patterns
- **State Management**: Zustand for client state, TanStack Query for server state
- **Form Handling**: React Hook Form with Zod validation
- **Styling**: TailwindCSS v4 with Shadcn UI components
- **Testing**: Vitest with React Testing Library

## Important Notes

- **Current Status**: Active development with authentication flow implemented (registration/login/logout)
- **AI Integration**: Planned Claude Haiku API integration for learning features
- **Internationalization**: Primarily focused on English learning for Chinese speakers
- **Security**: JWT-based authentication implemented
- **Performance**: Modern build tools with Vite for fast development

## Documentation References

For detailed information, refer to the comprehensive documentation in the `docs/` directory:
- Technical architecture details in `02-technical-architecture.md`
- Development workflow in `03-development-workflow.md`
- Feature specifications in `04-feature-modules.md`
- Deployment procedures in `05-deployment-guide.md`

## Development Philosophy and Code Standards

### Core Development Principles
- **Expertise Focus**: Senior-level React, TypeScript, and modern web development
- **Code Quality**: DRY principle, bug-free, fully functional implementations
- **Readability First**: Prioritize code clarity and maintainability over performance
- **Complete Implementation**: No TODOs, placeholders, or missing pieces
- **Accessibility**: Implement proper ARIA labels, keyboard navigation, and semantic HTML

### Code Implementation Rules

#### General Guidelines
- Use early returns for improved code readability
- Prefer `const` arrow functions over function declarations
- Use descriptive variable and function names with proper prefixes (e.g., `handleClick`)
- Implement accessibility features on interactive elements
- Define TypeScript types whenever possible

#### Frontend-Specific Rules
- **Styling**: Always use TailwindCSS classes; avoid inline CSS
- **State Management**: Use `??` for value fallback instead of `||`
- **Functions**: Prefer pure functions when possible
- **Logging**: Use `JSON.stringify()` for complex object logging
- **Discussion vs Implementation**: Only modify code when explicitly asked to implement

#### Communication
- **Language**: All explanations and thinking process in Traditional Chinese
- **Problem Solving**: When resolving issues, provide:
  - Brief problem description
  - Problematic code identification
  - Root cause analysis
  - Solution explanation

## Current API Endpoints

The backend currently implements the following API endpoints:

### Authentication (`/api/v1/auth`)
- **POST** `/register` - User registration with email, username, password validation
- **POST** `/login` - User login with JWT token generation
- **POST** `/logout` - User logout (requires JWT token)
- **GET** `/me` - Get current user profile (requires JWT token)

### System (`/`)
- **GET** `/health` - Health check with database status
- **GET** `/api/v1/ping` - Simple ping endpoint

## Database Schema

Current database tables implemented:
- **users** - User accounts with authentication data
  - Fields: id, email, username, password_hash, learning_level, avatar_url, created_at, updated_at
  - Constraints: unique email, unique username, email validation, password requirements

## Testing Configuration

### Frontend Testing
- **Test Runner**: Vitest with jsdom environment
- **Testing Library**: React Testing Library for component testing
- **Setup File**: `./src/test/setup.ts` for global test configuration
- **Coverage**: Available via `npm run test:coverage`

### Backend Testing
- **Test Framework**: Go's built-in testing package
- **Test Coverage**: Available via `make coverage`
- **Test Structure**: Tests organized alongside source files with `_test.go` suffix

所有程式碼請都使用 functional 的寫法