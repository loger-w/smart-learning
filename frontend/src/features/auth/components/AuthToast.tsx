import { toast } from "sonner";

interface AuthToastProps {
  message: string;
}

export const AuthToast = ({ message }: AuthToastProps) => {
  return;
  toast("發生錯誤", {
    description: `錯誤訊息：${message}`,
  });
};
