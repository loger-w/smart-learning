import { cn } from "@/lib/utils";
import { TbBook } from "react-icons/tb";

interface AuthHeaderProps {
  title: string;
  description: string;
  showLogo?: boolean;
  logoClassName?: string;
}

export const AuthHeader = ({
  title,
  description,
  showLogo = true,
  logoClassName,
}: AuthHeaderProps) => {
  return (
    <div className="text-center mb-8">
      {showLogo && (
        <div
          className={cn(
            "mx-auto w-16 h-16 text-white text-4xl bg-blue-600 rounded-2xl flex items-center justify-center mb-4",
            logoClassName
          )}
        >
          <TbBook />
        </div>
      )}
      <h1 className="text-3xl font-bold text-gray-900 mb-2">{title}</h1>
      <p className="text-gray-600">{description}</p>
    </div>
  );
};
