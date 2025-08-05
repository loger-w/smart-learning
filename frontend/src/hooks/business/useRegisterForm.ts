import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { z } from "zod";
import { toast } from "sonner";
import { useRegister } from "../auth/useAuth";
import { isAuthError, isValidationError } from "@/utils/errors";

const registerSchema = z
  .object({
    email: z.string().min(1, "請輸入電子郵件").email("電子郵件格式不正確"),
    username: z
      .string()
      .min(1, "請輸入用戶名")
      .min(2, "用戶名至少需要 2 個字符")
      .max(20, "用戶名最多 20 個字符")
      .regex(/^[a-zA-Z0-9_]+$/, "用戶名只能包含字母、數字和底線"),
    password: z.string().min(1, "請輸入密碼").min(8, "密碼至少需要 8 個字符"),
    confirm_password: z.string().min(1, "請確認密碼"),
  })
  .refine((data) => data.password === data.confirm_password, {
    message: "確認密碼與密碼不一致",
    path: ["confirm_password"],
  });

export type RegisterFormData = z.infer<typeof registerSchema>;

export const useRegisterForm = () => {
  const navigate = useNavigate();
  const registerMutation = useRegister();

  const form = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    mode: "onChange",
    defaultValues: {
      email: "",
      username: "",
      password: "",
      confirm_password: "",
    },
  });

  const onSubmit = async (data: RegisterFormData) => {
    try {
      await registerMutation.mutateAsync(data);
      toast.success("註冊成功！歡迎使用 Smart Learning");
      // 註冊成功後導向儀表板
      navigate({ to: "/dashboard" });
    } catch (error) {
      console.error("Register error:", error);
      
      if (isAuthError(error)) {
        toast.error(error.message);
      } else if (isValidationError(error)) {
        // 處理驗證錯誤，設置表單錯誤
        Object.entries(error.fields).forEach(([field, messages]) => {
          if (field === 'email' || field === 'username' || field === 'password' || field === 'confirm_password') {
            form.setError(field as keyof RegisterFormData, { message: messages[0] });
          }
        });
        toast.error("請檢查輸入的資料");
      } else {
        toast.error("註冊失敗，請稍後再試");
      }
    }
  };

  return {
    form,
    isLoading: registerMutation.isPending,
    onSubmit,
  };
};
