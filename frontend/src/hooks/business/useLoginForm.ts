import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { z } from "zod";
import { toast } from "sonner";
import { useLogin } from "../auth/useAuth";
import { isAuthError, isValidationError } from "@/utils/errors";

const loginSchema = z.object({
  email: z.string().min(1, "請輸入電子郵件").email("請輸入有效的電子郵件"),
  password: z.string().min(1, "請輸入密碼").min(8, "密碼至少需要 8 個字符"),
});

export type LoginFormData = z.infer<typeof loginSchema>;

export const useLoginForm = () => {
  const navigate = useNavigate();
  const loginMutation = useLogin();

  const form = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (data: LoginFormData) => {
    try {
      await loginMutation.mutateAsync(data);
      toast.success("登入成功！");
      // 登入成功後導向儀表板
      navigate({ to: "/dashboard" });
    } catch (error) {
      console.error("Login error:", error);

      if (isAuthError(error)) {
        toast.error(error.message);
      } else if (isValidationError(error)) {
        // 處理驗證錯誤，設置表單錯誤
        Object.entries(error.fields).forEach(([field, messages]) => {
          if (field === "email" || field === "password") {
            form.setError(field, { message: messages[0] });
          }
        });
        toast.error("請檢查輸入的資料");
      } else {
        console.log(error);
        toast.error("登入失敗，請稍後再試");
      }
    }
  };

  return {
    form,
    isLoading: loginMutation.isPending,
    onSubmit,
  };
};
