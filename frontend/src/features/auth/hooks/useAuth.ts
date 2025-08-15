import { useAuthStore } from "@/stores/authStore";
import { useLogin, useRegister, useLogout } from "./useAuthApi";
import type { LoginRequest, RegisterRequest } from "@/types/index";

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