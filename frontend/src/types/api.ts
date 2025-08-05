export interface APIResponse<T> {
  success: boolean
  message?: string
  data?: T
  error?: APIError
  errors?: Record<string, string[]>
}

export interface APIError {
  code: string
  message: string
}

// to-do
export interface HealthResponse {
  status: string
  message: string
  db_stats?: {
    max_open_connections: number
    open_connections: number
    in_use: number
    idle: number
  }
}

// to-do
export interface PingResponse {
  message: string
}