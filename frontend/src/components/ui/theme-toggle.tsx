import { Moon, Sun } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useTheme } from "@/hooks/use-theme";

export const ThemeToggle = () => {
  const { setTheme } = useTheme();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="sm">
          <Sun className="h-4 w-4 rotate-0 scale-100 transition-[transform,opacity] duration-300 dark:-rotate-90 dark:scale-0" />
          <Moon className="absolute h-4 w-4 rotate-90 scale-0 transition-[transform,opacity] duration-300 dark:rotate-0 dark:scale-100" />
          <span className="sr-only">切換主題</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem onClick={() => setTheme("light")}>
          明亮模式
        </DropdownMenuItem>
        <DropdownMenuItem onClick={() => setTheme("dark")}>
          深色模式
        </DropdownMenuItem>
        <DropdownMenuItem onClick={() => setTheme("system")}>
          系統預設
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
