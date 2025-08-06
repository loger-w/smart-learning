import { useForm } from "react-hook-form";
import { z } from "zod";
import { toast } from "sonner";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { useRegister } from "@/hooks/auth/useAuth";

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

      toast.error("註冊失敗，請稍後再試");
    }
  };

  return {
    form,
    isLoading: registerMutation.isPending,
    onSubmit,
  };
};
