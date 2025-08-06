import { useForm } from "react-hook-form";
import { z } from "zod";
import { toast } from "sonner";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { useLogin } from "@/hooks/auth/useAuth";

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
    await loginMutation.mutateAsync(data);
    toast.success("登入成功！");
    // 登入成功後導向儀表板
    navigate({ to: "/dashboard" });
  };

  return {
    form,
    isLoading: loginMutation.isPending,
    onSubmit,
  };
};
