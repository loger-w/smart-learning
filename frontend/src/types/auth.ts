import type { APIResponse } from "./api";

export interface User {
  id: number;
  email: string;
  username: string;
  learning_level: number;
  avatar_url?: string | null;
  created_at: string;
  updated_at: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
  confirm_password: string;
}

export type LoginResponse = APIResponse<AuthResponse>;
export type RegisterResponse = APIResponse<AuthResponse>;
export type MeResponse = APIResponse<{ user: User }>;
