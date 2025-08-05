import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { authService } from "@/api/services/authService";
import { useAuthStore } from "@/stores/authStore";
import type { LoginRequest, RegisterRequest } from "@/types/auth";
import { isAuthError, isValidationError } from "@/utils/errors";

export const useLogin = () => {
  const { setAuth } = useAuthStore();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: authService.login,
    onSuccess: (data) => {
      setAuth(data.user, data.token);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
    },
    onError: (error) => {
      if (isAuthError(error) || isValidationError(error)) {
        console.error("Login failed:", error.message);
      } else {
        console.error("Login failed:", error);
      }
    },
  });
};

export const useRegister = () => {
  const { setAuth } = useAuthStore();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: authService.register,
    onSuccess: (data) => {
      setAuth(data.user, data.token);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
    },
    onError: (error) => {
      if (isAuthError(error) || isValidationError(error)) {
        console.error("Registration failed:", error.message);
      } else {
        console.error("Registration failed:", error);
      }
    },
  });
};

export const useMe = () => {
  const { token } = useAuthStore();

  return useQuery({
    queryKey: ["auth", "me"],
    queryFn: authService.getMe,
    enabled: !!token,
    staleTime: 5 * 60 * 1000, // 5 分鐘
    retry: (failureCount, error: unknown) => {
      // 如果是 401 錯誤，不重試
      if (error && typeof error === "object" && "response" in error) {
        const axiosError = error as { response?: { status?: number } };
        if (axiosError.response?.status === 401) {
          return false;
        }
      }
      return failureCount < 3;
    },
  });
};

export const useLogout = () => {
  const { logout } = useAuthStore();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: authService.logout,
    onSuccess: () => {
      logout();
      queryClient.clear();
    },
    onError: () => {
      // 即使 API 調用失敗，也要清除本地狀態
      logout();
      queryClient.clear();
    },
  });
};

// useAuth 作為狀態聚合器，使用其他 hooks 而不是重複實現邏輯
export const useAuth = () => {
  const { user, token, isAuthenticated } = useAuthStore();
  const loginMutation = useLogin();
  const registerMutation = useRegister();
  const logoutMutation = useLogout();

  const login = async (credentials: LoginRequest) => {
    return loginMutation.mutateAsync(credentials);
  };

  const register = async (data: RegisterRequest) => {
    return registerMutation.mutateAsync(data);
  };

  const logout = async () => {
    return logoutMutation.mutateAsync();
  };

  return {
    user,
    token,
    isAuthenticated,
    login,
    register,
    logout,
    isLoading:
      loginMutation.isPending ||
      registerMutation.isPending ||
      logoutMutation.isPending,
    error:
      loginMutation.error || registerMutation.error || logoutMutation.error,
  };
};
