import React from "react";
import { cn } from "@/lib/utils";
import { Card, CardContent } from "@/components/ui/card";

interface AuthFormContainerProps {
  children: React.ReactNode;
  className?: string;
}

export const AuthFormContainer = ({
  children,
  className,
}: AuthFormContainerProps) => {
  return (
    <Card className={cn("shadow-lg", className)}>
      <CardContent className="p-8">
        {children}
      </CardContent>
    </Card>
  );
};
