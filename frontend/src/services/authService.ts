import { apiClient } from "./api";
import type {
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  RegisterResponse,
  MeResponse,
  AuthResponse,
  User,
} from "@/types/auth";
import { AxiosError } from "axios";

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    try {
      const response = await apiClient.post<LoginResponse>(
        "/auth/login",
        credentials
      );
      return response.data.data!;
    } catch (error) {
      if (error instanceof AxiosError) {
        const apiError = {
          code: error.response!.data.error.code,
          message: error.response!.data.error.message,
          status: error.status,
        };
        throw apiError;
      }

      throw {
        message: "系統暫時無法使用",
        code: "SERVICE_UNAVAILABLE",
        status: 500,
      };
    }
  },

  async register(data: RegisterRequest): Promise<AuthResponse> {
    try {
      const response = await apiClient.post<RegisterResponse>(
        "/auth/register",
        data
      );

      return response.data.data!;
    } catch (error) {
      console.log(error);

      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },

  async logout(): Promise<void> {
    try {
      await apiClient.post("/auth/logout");
    } catch (error) {
      if (error) {
        console.log(error);
        // 重新拋出非 Axios 錯誤
        throw error;
      }
    }
  },

  async getMe(): Promise<User> {
    try {
      const response = await apiClient.get<MeResponse>("/auth/me");
      return response.data.data!.user;
    } catch (error) {
      console.log(error);
      // 重新拋出非 Axios 錯誤
      throw error;
    }
  },
};
