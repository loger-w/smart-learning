import { Button } from "@/components/ui/button";
import { toast } from "sonner";
import { cn } from "@/lib/utils";

interface ForgotPasswordButtonProps {
  className?: string;
}

export const ForgotPasswordButton = ({
  className,
}: ForgotPasswordButtonProps) => {
  const handleForgotPassword = () => {
    // TODO: 實際的忘記密碼邏輯
    toast("忘記密碼功能開發中");
  };

  return (
    <div className={cn("text-center", className)}>
      <Button
        type="button"
        onClick={handleForgotPassword}
        variant="link"
        size="sm"
      >
        忘記密碼？
      </Button>
    </div>
  );
};
