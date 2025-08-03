import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useAuthStore } from '@/stores/authStore'
import { authService } from '@/services/authService'
import type { LoginRequest, RegisterRequest } from '@/types/auth'

export const useAuth = () => {
  const { user, token, isAuthenticated, setAuth, logout: clearAuth } = useAuthStore()
  const queryClient = useQueryClient()

  const loginMutation = useMutation({
    mutationFn: authService.login,
    onSuccess: (data) => {
      setAuth(data.user, data.token)
      queryClient.invalidateQueries({ queryKey: ['user'] })
    },
    onError: (error) => {
      console.error('Login failed:', error)
    },
  })

  const registerMutation = useMutation({
    mutationFn: authService.register,
    onSuccess: (data) => {
      setAuth(data.user, data.token)
      queryClient.invalidateQueries({ queryKey: ['user'] })
    },
    onError: (error) => {
      console.error('Registration failed:', error)
    },
  })

  const logoutMutation = useMutation({
    mutationFn: authService.logout,
    onSuccess: () => {
      clearAuth()
      queryClient.clear()
    },
    onError: () => {
      // 即使 API 調用失敗，也要清除本地狀態
      clearAuth()
      queryClient.clear()
    },
  })

  const login = async (credentials: LoginRequest) => {
    return loginMutation.mutateAsync(credentials)
  }

  const register = async (data: RegisterRequest) => {
    return registerMutation.mutateAsync(data)
  }

  const logout = async () => {
    return logoutMutation.mutateAsync()
  }

  return {
    user,
    token,
    isAuthenticated,
    login,
    register,
    logout,
    isLoading: loginMutation.isPending || registerMutation.isPending || logoutMutation.isPending,
    error: loginMutation.error || registerMutation.error || logoutMutation.error,
  }
}