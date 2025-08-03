import axios from 'axios'
import type { LoginRequest, LoginResponse, RegisterRequest } from '@/types/auth'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api'

const authAPI = axios.create({
  baseURL: `${API_BASE_URL}/auth`,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const authService = {
  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await authAPI.post<LoginResponse>('/login', credentials)
    return response.data
  },

  async register(data: RegisterRequest): Promise<LoginResponse> {
    const response = await authAPI.post<LoginResponse>('/register', data)
    return response.data
  },

  async logout(): Promise<void> {
    await authAPI.post('/logout')
  },

  async refreshToken(): Promise<{ token: string }> {
    const response = await authAPI.post<{ token: string }>('/refresh')
    return response.data
  },
}

// 設定請求攔截器，自動添加 token
authAPI.interceptors.request.use((config) => {
  const token = localStorage.getItem('auth-store')
  if (token) {
    try {
      const parsedData = JSON.parse(token)
      if (parsedData.state?.token) {
        config.headers.Authorization = `Bearer ${parsedData.state.token}`
      }
    } catch (error) {
      console.error('Error parsing auth token:', error)
    }
  }
  return config
})

// 設定回應攔截器，處理 token 過期
authAPI.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Token 過期，清除本地存儲
      localStorage.removeItem('auth-store')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)