import { apiClient } from "../apiClient";
import type {
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  RegisterResponse,
  MeResponse,
  AuthResponse,
  User,
} from "@/types/auth";
import { createAuthError, createValidationError } from "@/utils/errors";
import { AxiosError } from "axios";

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    try {
      const response = await apiClient.post<LoginResponse>(
        "/auth/login",
        credentials
      );
      
      if (!response.data.success) {
        // 處理驗證錯誤
        if (response.data.errors) {
          throw createValidationError(response.data.errors);
        }
        throw createAuthError(
          response.data.error?.message || "發生未知錯誤，登入失敗"
        );
      }

      return response.data.data!;
    } catch (error) {
      if (error instanceof AxiosError && error.response?.data) {
        const errorData = error.response.data;
        
        // 處理驗證錯誤
        if (errorData.errors) {
          throw createValidationError(errorData.errors);
        }
        
        // 優先使用 error.message，然後是 message，最後是默認訊息
        const errorMessage = errorData.error?.message || errorData.message || "登入失敗";
        throw createAuthError(errorMessage);
      }
      
      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },

  async register(data: RegisterRequest): Promise<AuthResponse> {
    try {
      const response = await apiClient.post<RegisterResponse>(
        "/auth/register",
        data
      );

      if (!response.data.success) {
        // 處理驗證錯誤
        if (response.data.errors) {
          throw createValidationError(response.data.errors);
        }
        throw createAuthError(
          response.data.error?.message || "發生未知錯誤，註冊失敗"
        );
      }

      return response.data.data!;
    } catch (error) {
      if (error instanceof AxiosError && error.response?.data) {
        const errorData = error.response.data;
        
        // 處理驗證錯誤
        if (errorData.errors) {
          throw createValidationError(errorData.errors);
        }
        
        // 優先使用 error.message，然後是 message，最後是默認訊息
        const errorMessage = errorData.error?.message || errorData.message || "註冊失敗";
        throw createAuthError(errorMessage);
      }
      
      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },

  async logout(): Promise<void> {
    try {
      await apiClient.post("/auth/logout");
    } catch (error) {
      if (error instanceof AxiosError && error.response?.data) {
        const errorData = error.response.data;
        const errorMessage = errorData.error?.message || errorData.message || "登出失敗";
        throw createAuthError(errorMessage);
      }
      
      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },

  async getMe(): Promise<User> {
    try {
      const response = await apiClient.get<MeResponse>("/auth/me");

      if (!response.data.success) {
        throw createAuthError(
          response.data.error?.message || "發生未知錯誤，獲取用戶資料失敗"
        );
      }

      return response.data.data!.user;
    } catch (error) {
      if (error instanceof AxiosError && error.response?.data) {
        const errorData = error.response.data;
        const errorMessage = errorData.error?.message || errorData.message || "獲取用戶資料失敗";
        throw createAuthError(errorMessage);
      }
      
      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },
};
