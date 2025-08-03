import { Button } from "@/components/ui/button";
import { toast } from "sonner";
import { cn } from "@/lib/utils";

interface RegisterButtonProps {
  className?: string;
}

export const RegisterButton = ({ className }: RegisterButtonProps) => {
  const handleRegister = () => {
    // TODO: 導航到註冊頁面
    return toast("註冊功能開發中");
  };
  return (
    <div
      className={cn(
        "mt-8 pt-6 border-t border-gray-200 text-center",
        className
      )}
    >
      <p className="text-gray-600 text-sm">
        還沒有帳戶？
        <Button
          onClick={handleRegister}
          variant="link"
          size="sm"
          className="ml-1 p-0 h-auto"
        >
          立即註冊
        </Button>
      </p>
    </div>
  );
};
