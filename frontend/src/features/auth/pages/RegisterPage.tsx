import {
  AuthLayout,
  AuthHeader,
  AuthFormContainer,
  AuthLinkButton,
  RegisterForm,
} from "@/features/auth/components";

export const RegisterPage = () => {
  return (
    <AuthLayout>
      <AuthHeader title="加入 Smart Learning" description="建立您的帳戶" />

      <AuthFormContainer>
        <RegisterForm />

        {/* 登入連結 */}
        <AuthLinkButton
          text="已經有帳戶了？"
          buttonText="立即登入"
          targetNavigate="/auth/login"
        />
      </AuthFormContainer>
    </AuthLayout>
  );
};
