import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

interface LoginButtonProps {
  isLoading: boolean;
  disabled?: boolean;
  onClick?: () => void;
  className?: string;
}

export const LoginButton = ({
  isLoading,
  disabled,
  onClick,
  className,
}: LoginButtonProps) => {
  return (
    <Button
      type="submit"
      disabled={isLoading || disabled}
      onClick={onClick}
      className={cn("w-full", className)}
      size="lg"
    >
      {isLoading ? (
        <div className="flex items-center justify-center space-x-2">
          <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
          <span>登入中...</span>
        </div>
      ) : (
        "登入"
      )}
    </Button>
  );
};
