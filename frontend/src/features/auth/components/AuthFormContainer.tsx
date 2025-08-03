import React from "react";
import { cn } from "@/lib/utils";

interface AuthFormContainerProps {
  children: React.ReactNode;
  className?: string;
}

export const AuthFormContainer = ({
  children,
  className,
}: AuthFormContainerProps) => {
  return (
    <div className={cn("bg-white rounded-xl shadow-lg p-8", className)}>
      {children}
    </div>
  );
};
