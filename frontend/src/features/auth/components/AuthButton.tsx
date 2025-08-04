import { Button } from "@/components/ui/button";

interface AuthButtonProps {
  action: string;
  disable?: boolean;
  isLoading: boolean;
  onClick?: () => void;
}

export const AuthButton = ({
  action,
  disable = false,
  isLoading,
  onClick,
}: AuthButtonProps) => {
  return (
    <Button
      type="submit"
      disabled={disable}
      onClick={onClick}
      className="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed text-white font-medium py-2 px-4 rounded-lg transition-colors"
    >
      {isLoading ? `${action}中...` : `${action}`}
    </Button>
  );
};
