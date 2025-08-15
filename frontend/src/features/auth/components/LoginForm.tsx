import { Input } from "@/components/ui/input";
import { useLoginForm } from "../hooks";
import { AuthButton, ForgotPasswordButton } from "./";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

export const LoginForm = () => {
  const { form, isLoading, onSubmit } = useLoginForm();

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
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


        {/* 登入按鈕 */}
        <AuthButton
          action="登入"
          disable={isLoading}
          isLoading={isLoading}
        />

        {/* 忘記密碼連結 */}
        <ForgotPasswordButton />
      </form>
    </Form>
  );
};
