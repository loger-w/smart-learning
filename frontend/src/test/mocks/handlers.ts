import { http, HttpResponse } from 'msw'
import type { AuthResponse, LoginRequest, RegisterRequest } from '@/types/auth'

const mockUser = {
  id: 1,
  email: 'test@example.com',
  username: 'testuser',
  learning_level: 1,
  avatar_url: null,
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
}

const mockAuthResponse: AuthResponse = {
  user: mockUser,
  token: 'mock-jwt-token'
}

export const handlers = [
  // 登入成功  
  http.post('http://localhost:3000/api/v1/auth/login', async ({ request }) => {
    const body = await request.json() as LoginRequest
    
    if (body.email === 'test@example.com' && body.password === 'password123') {
      return HttpResponse.json({
        success: true,
        message: '登入成功',
        data: mockAuthResponse
      })
    }
    
    return HttpResponse.json({
      success: false,
      message: '登入失敗',
      error: {
        code: 'INVALID_CREDENTIALS',
        message: '電子郵件或密碼錯誤'
      }
    }, { status: 401 })
  }),

  // 註冊成功
  http.post('http://localhost:3000/api/v1/auth/register', async ({ request }) => {
    const body = await request.json() as RegisterRequest
    
    if (body.email === 'existing@example.com') {
      return HttpResponse.json({
        success: false,
        message: '用戶已存在',
        error: {
          code: 'USER_ALREADY_EXISTS',
          message: '電子郵件或用戶名已被使用'
        }
      }, { status: 409 })
    }
    
    return HttpResponse.json({
      success: true,
      message: '註冊成功',
      data: mockAuthResponse
    }, { status: 201 })
  }),

  // 登出成功
  http.post('http://localhost:3000/api/v1/auth/logout', () => {
    return HttpResponse.json({
      success: true,
      message: '登出成功'
    })
  }),

  // 獲取用戶資訊
  http.get('http://localhost:3000/api/v1/auth/me', ({ request }) => {
    const authHeader = request.headers.get('Authorization')
    
    if (!authHeader || authHeader !== 'Bearer mock-jwt-token') {
      return HttpResponse.json({
        success: false,
        message: '未授權',
        error: {
          code: 'UNAUTHORIZED',
          message: '無法獲取用戶資訊'
        }
      }, { status: 401 })
    }
    
    return HttpResponse.json({
      success: true,
      data: { user: mockUser }
    })
  })
]