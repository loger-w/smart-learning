export interface APIErrorOptions {
  code?: string
  status?: number
}

export const createAPIError = (message: string, options?: APIErrorOptions): Error & APIErrorOptions => {
  const error = new Error(message) as Error & APIErrorOptions
  error.name = 'APIError'
  error.code = options?.code
  error.status = options?.status
  return error
}

export const createAuthError = (message: string, code?: string): Error & APIErrorOptions => {
  const error = new Error(message) as Error & APIErrorOptions
  error.name = 'AuthError'
  error.code = code
  error.status = 401
  return error
}

export interface ValidationErrorData {
  fields: Record<string, string[]>
}

export const createValidationError = (fields: Record<string, string[]>): Error & ValidationErrorData => {
  const error = new Error('驗證失敗') as Error & ValidationErrorData
  error.name = 'ValidationError'
  error.fields = fields
  return error
}

export const getFieldErrors = (error: ValidationErrorData, fieldName: string): string[] => {
  return error.fields[fieldName] || []
}

export const hasFieldError = (error: ValidationErrorData, fieldName: string): boolean => {
  return Array.isArray(error.fields[fieldName]) && error.fields[fieldName].length > 0
}

// Type guards for error checking
export const isAPIError = (error: unknown): error is Error & APIErrorOptions => {
  return error instanceof Error && error.name === 'APIError'
}

export const isAuthError = (error: unknown): error is Error & APIErrorOptions => {
  return error instanceof Error && error.name === 'AuthError'
}

export const isValidationError = (error: unknown): error is Error & ValidationErrorData => {
  return error instanceof Error && error.name === 'ValidationError'
}