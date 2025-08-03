import { Input } from "@/components/ui/input";
import { useLoginForm } from "@/hooks/auth/useLoginForm";
import { LoginButton, ForgotPasswordButton } from "./";

export const LoginForm = () => {
  const { register, handleSubmit, errors, isLoading, onSubmit } =
    useLoginForm();

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      {/* 電子郵件 */}
      <div>
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-700 mb-2"
        >
          電子郵件
        </label>
        <Input
          {...register("email")}
          type="email"
          id="email"
          placeholder="example@email.com"
          aria-invalid={errors.email ? "true" : "false"}
        />
        {errors.email && (
          <p className="text-red-500 text-sm mt-1" role="alert">
            {errors.email.message}
          </p>
        )}
      </div>

      {/* 密碼 */}
      <div>
        <label
          htmlFor="password"
          className="block text-sm font-medium text-gray-700 mb-2"
        >
          密碼
        </label>
        <div className="relative">
          <Input
            {...register("password")}
            type="password"
            id="password"
            placeholder="輸入您的密碼"
            className="pr-12"
            aria-invalid={errors.password ? "true" : "false"}
          />
        </div>
        {errors.password && (
          <p className="text-red-500 text-sm mt-1" role="alert">
            {errors.password.message}
          </p>
        )}
      </div>

      {/* 錯誤訊息 */}
      {errors.root && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <p className="text-red-600 text-sm" role="alert">
            {errors.root.message}
          </p>
        </div>
      )}

      {/* 登入按鈕 */}
      <LoginButton isLoading={isLoading} />

      {/* 忘記密碼連結 */}
      <ForgotPasswordButton />
    </form>
  );
};
