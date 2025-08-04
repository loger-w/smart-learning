import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { useAuthStore } from "@/stores/authStore";
import { z } from "zod";
import { toast } from "sonner";

const registerSchema = z
  .object({
    email: z.email("請輸入有效的電子郵件"),
    username: z
      .string()
      .min(2, "用戶名至少需要 2 個字符")
      .max(20, "用戶名最多 20 個字符"),
    password: z.string().min(8, "密碼至少需要 8 個字符"),
    confirm_password: z.string().min(8, "請確認密碼"),
  })
  .refine((data) => data.password === data.confirm_password, {
    message: "密碼確認不符",
    path: ["confirm_password"],
  });

export type RegisterFormData = z.infer<typeof registerSchema>;

export const useRegisterForm = () => {
  const [isLoading, setIsLoading] = useState(false);
  const { setAuth } = useAuthStore();
  const navigate = useNavigate();

  const form = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: "",
      username: "",
      password: "",
      confirm_password: "",
    },
  });

  const onSubmit = async (data: RegisterFormData) => {
    setIsLoading(true);
    try {
      // TODO: 實際的註冊邏輯，這裡先模擬成功註冊
      console.log("Register attempt:", data);
      await new Promise((resolve) => setTimeout(resolve, 1000)); // 模擬 API 執行時間

      // 模擬註冊成功，設定假的用戶資料
      const mockUser = {
        id: Date.now(), // 使用時間戳作為假的 ID
        email: data.email,
        username: data.username,
        learning_level: 1, // 新用戶預設為初級
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };
      const mockToken = "mock-jwt-token";

      setAuth(mockUser, mockToken);

      // 註冊成功後導向儀表板
      navigate({ to: "/dashboard" });
    } catch (error) {
      console.error("Register error:", error);
      toast.error(`註冊失敗，${error}`);
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
