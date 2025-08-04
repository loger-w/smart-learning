import {
  AuthLayout,
  AuthHeader,
  AuthFormContainer,
  AuthLinkButton,
  LoginForm,
} from "@/features/auth/components";

export const LoginPage = () => {
  return (
    <AuthLayout>
      <AuthHeader title="歡迎使用 Smart Learning" description="登入您的帳戶" />

      <AuthFormContainer>
        <LoginForm />

        {/* 註冊連結 */}
        <AuthLinkButton
          text="還沒有帳戶？"
          buttonText="立即註冊"
          targetNavigate="/auth/register"
        />
      </AuthFormContainer>
    </AuthLayout>
  );
};
