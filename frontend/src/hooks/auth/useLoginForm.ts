import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { useAuthStore } from "@/stores/authStore";
import { z } from "zod";
import { toast } from "sonner";

const loginSchema = z.object({
  email: z.email("請輸入有效的電子郵件"),
  password: z.string().min(8, "密碼至少需要 8 個字符"),
});

export type LoginFormData = z.infer<typeof loginSchema>;

export const useLoginForm = () => {
  const [isLoading, setIsLoading] = useState(false);
  const { setAuth } = useAuthStore();
  const navigate = useNavigate();

  const form = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (data: LoginFormData) => {
    setIsLoading(true);
    try {
      // TODO: 實際的登入邏輯，這裡先模擬成功登入
      console.log("Login attempt:", data);
      await new Promise((resolve) => setTimeout(resolve, 1000)); // 模擬 API 執行時間
      // 模擬登入成功，設定假的用戶資料
      const mockUser = {
        id: 1,
        email: data.email,
        username: "Tester",
        learning_level: 3,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };
      const mockToken = "mock-jwt-token";

      setAuth(mockUser, mockToken);

      // 登入成功後導向儀表板
      navigate({ to: "/dashboard" });
    } catch (error) {
      console.error("Login error:", error);
      toast.error(`登入失敗，${error}`);
    } finally {
      setIsLoading(false);
    }
  };

  return {
    form,
    isLoading,
    onSubmit,
  };
};
