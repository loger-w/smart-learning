export interface User {
  id: number
  email: string
  username: string
  learning_level: number
  avatar_url?: string
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  user: User
  token: string
}

export interface RegisterRequest {
  email: string
  username: string
  password: string
  confirm_password: string
}

export interface AuthError {
  message: string
  code?: string
}