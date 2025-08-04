import { Button } from "@/components/ui/button";
import { useNavigate } from "@tanstack/react-router";

interface AuthLinkButtonProps {
  text: string;
  buttonText: string;
  targetNavigate: string;
}

export const AuthLinkButton = ({
  text,
  buttonText,
  targetNavigate,
}: AuthLinkButtonProps) => {
  const navigate = useNavigate();

  const handleNavigateToLogin = () => {
    navigate({ to: targetNavigate });
  };

  return (
    <div className={"mt-8 pt-6 border-t border-gray-200 text-center"}>
      <p className="text-gray-600 text-sm">
        {text}
        <Button
          onClick={handleNavigateToLogin}
          variant="link"
          size="sm"
          className="ml-1 p-0 h-auto"
        >
          {buttonText}
        </Button>
      </p>
    </div>
  );
};
