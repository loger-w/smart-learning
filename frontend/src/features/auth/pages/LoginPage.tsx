import {
  AuthLayout,
  AuthHeader,
  AuthFormContainer,
  LoginForm,
  RegisterButton,
} from "@/features/auth/components";

export const LoginPage = () => {
  return (
    <AuthLayout>
      <AuthHeader title="歡迎使用 Smart Learning" description="登入您的帳戶" />

      <AuthFormContainer>
        <LoginForm />

        {/* 註冊連結 */}
        <RegisterButton />
      </AuthFormContainer>
    </AuthLayout>
  );
};
