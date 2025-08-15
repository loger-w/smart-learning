import { toast } from "sonner";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuthStore } from "@/stores/authStore";
import { authService } from "@/services/authService";
import type { APIErrorResponse } from "@/types/index";

export const useLogin = () => {
  const { setAuth } = useAuthStore();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: authService.login,
    onSuccess: (data) => {
      setAuth(data.user, data.token);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
    },
    onError: (error: APIErrorResponse) => {
      console.error("登入失敗", JSON.stringify(error));
      toast(error.message);
    },
  });
};

export const useRegister = () => {
  const { setAuth } = useAuthStore();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: authService.register,
    onSuccess: (data) => {
      console.log(data);
      setAuth(data.user, data.token);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
    },
    onError: (error) => {
      console.log(error);
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