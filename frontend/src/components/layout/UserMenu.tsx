import { useAuthStore } from "@/stores/authStore";
import { useRouter } from "@tanstack/react-router";
import { useLogout } from "@/features/auth/hooks";
import { toast } from "sonner";
import { useQueryClient } from "@tanstack/react-query";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ThemeToggle } from "@/components/ui/theme-toggle";
import { User, LogOut, Palette } from "lucide-react";

export const UserMenu = () => {
  const { user, logout: clearAuth } = useAuthStore();
  const router = useRouter();
  const logoutMutation = useLogout();
  const queryClient = useQueryClient();

  const handleLogout = async () => {
    try {
      await logoutMutation.mutateAsync();
      toast.success("已成功登出");
      router.navigate({ to: "/auth/login" });
    } catch (error) {
      console.error("Logout error:", error);
      clearAuth();
      queryClient.clear();
      toast.error("登出時發生錯誤，但已清除本地資料");
      router.navigate({ to: "/auth/login" });
    }
  };

  if (!user) return null;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="flex items-center gap-3 w-full p-2 rounded-lg hover:bg-sidebar-accent hover:text-sidebar-accent-foreground transition-colors">
        <Avatar className="h-8 w-8">
          <AvatarImage src={user.avatar_url ?? undefined} alt={user.username} />
          <AvatarFallback>
            <User className="h-4 w-4" />
          </AvatarFallback>
        </Avatar>
        <div className="flex flex-col items-start text-left group-data-[collapsible=icon]:hidden">
          <span className="text-sm font-medium truncate">{user.username}</span>
          <span className="text-xs text-muted-foreground truncate">{user.email}</span>
        </div>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-56">
        <DropdownMenuLabel>我的帳戶</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem className="flex items-center justify-between">
          <div className="flex items-center">
            <Palette className="mr-2 h-4 w-4" />
            主題設定
          </div>
          <ThemeToggle />
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onClick={handleLogout}
          disabled={logoutMutation.isPending}
          className="text-red-600 focus:text-red-600 focus:bg-red-50"
        >
          <LogOut className="mr-2 h-4 w-4" />
          {logoutMutation.isPending ? "登出中..." : "登出"}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};