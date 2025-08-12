import { describe, it, expect, vi, beforeEach } from 'vitest'
import { authService } from './authService'
import type { LoginRequest, RegisterRequest } from '@/types/auth'

// Mock localStorage
const localStorageMock = {
  getItem: vi.fn(),
  setItem: vi.fn(),
  removeItem: vi.fn(),
  clear: vi.fn(),
}

Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
})

describe('authService', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    localStorageMock.getItem.mockReturnValue('mock-jwt-token')
  })
  describe('login', () => {
    it('應該成功登入並返回用戶資料和 token', async () => {
      const credentials: LoginRequest = {
        email: 'test@example.com',
        password: 'password123'
      }

      const result = await authService.login(credentials)

      expect(result).toEqual({
        user: {
          id: 1,
          email: 'test@example.com',
          username: 'testuser',
          learning_level: 1,
          avatar_url: null,
          created_at: '2024-01-01T00:00:00Z',
          updated_at: '2024-01-01T00:00:00Z'
        },
        token: 'mock-jwt-token'
      })
    })

    it('應該處理無效憑證錯誤', async () => {
      const credentials: LoginRequest = {
        email: 'wrong@example.com',
        password: 'wrongpassword'
      }

      await expect(authService.login(credentials)).rejects.toMatchObject({
        code: 'INVALID_CREDENTIALS',
        message: '電子郵件或密碼錯誤'
      })
    })
  })

  describe('register', () => {
    it('應該成功註冊並返回用戶資料和 token', async () => {
      const registerData: RegisterRequest = {
        email: 'newuser@example.com',
        username: 'newuser',
        password: 'password123',
        confirm_password: 'password123'
      }

      const result = await authService.register(registerData)

      expect(result).toEqual({
        user: {
          id: 1,
          email: 'test@example.com',
          username: 'testuser',
          learning_level: 1,
          avatar_url: null,
          created_at: '2024-01-01T00:00:00Z',
          updated_at: '2024-01-01T00:00:00Z'
        },
        token: 'mock-jwt-token'
      })
    })
  })

  describe('logout', () => {
    it('應該成功登出', async () => {
      await expect(authService.logout()).resolves.toBeUndefined()
    })
  })

  describe('getMe', () => {
    it('應該成功獲取用戶資訊', async () => {
      const result = await authService.getMe()

      expect(result).toEqual({
        id: 1,
        email: 'test@example.com',
        username: 'testuser',
        learning_level: 1,
        avatar_url: null,
        created_at: '2024-01-01T00:00:00Z',
        updated_at: '2024-01-01T00:00:00Z'
      })
    })
  })
})