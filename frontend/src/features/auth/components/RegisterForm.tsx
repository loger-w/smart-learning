import { Input } from "@/components/ui/input";
import { AuthButton } from "./";
import { useRegisterForm } from "@/hooks/auth/useRegisterForm";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

export const RegisterForm = () => {
  const { form, isLoading, onSubmit } = useRegisterForm();

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
        {/* 電子郵件 */}
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>電子郵件</FormLabel>
              <FormControl>
                <Input placeholder="example@email.com" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* 用戶名 */}
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>用戶名</FormLabel>
              <FormControl>
                <Input placeholder="輸入您的用戶名" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* 密碼 */}
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>密碼</FormLabel>
              <FormControl>
                <Input type="password" placeholder="輸入您的密碼" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* 確認密碼 */}
        <FormField
          control={form.control}
          name="confirm_password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>確認密碼</FormLabel>
              <FormControl>
                <Input type="password" placeholder="再次輸入密碼" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* 註冊按鈕 */}
        <AuthButton
          action="註冊"
          disable={isLoading || !form.formState.isValid}
          isLoading={isLoading}
        />
      </form>
    </Form>
  );
};
