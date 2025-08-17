import { useAuthStore } from "@/stores/authStore";
import { Button } from "@/components/ui/button";
import { useRouter } from "@tanstack/react-router";
import { useLogout } from "@/features/auth/hooks";
import { toast } from "sonner";
import { useQueryClient } from "@tanstack/react-query";
import {
  BookOpen,
  CreditCard,
  FileText,
  Home,
  BarChart3,
} from "lucide-react";
import { TbBook } from "react-icons/tb";

const navigationItems = [
  {
    label: "首頁",
    to: "/dashboard",
    icon: Home,
  },
  {
    label: "記憶卡片",
    to: "/flashcards",
    icon: CreditCard,
  },
  {
    label: "單字練習",
    to: "/vocabulary",
    icon: BookOpen,
  },
  {
    label: "文章練習",
    to: "/articles",
    icon: FileText,
  },
  {
    label: "學習統計",
    to: "/statistics",
    icon: BarChart3,
  },
];

export const AppHeader = () => {
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
      // 即使 API 失敗，也要確保清除本地狀態並導向登入頁
      console.error("Logout error:", error);
      clearAuth();
      queryClient.clear();
      toast.error("登出時發生錯誤，但已清除本地資料");
      router.navigate({ to: "/auth/login" });
    }
  };

  const handleNavigation = (to: string) => {
    // 目前先用 toast 提示，未來實際實現時移除
    if (to !== "/dashboard") {
      toast.info(`${to} 功能即將推出`);
      return;
    }
    router.navigate({ to });
  };

  return (
    <header className="bg-background shadow-sm border-b border-border">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo and Title */}
          <div className="flex items-center space-x-4">
            <div className="w-8 h-8 bg-blue-600 text-white rounded-lg flex items-center justify-center">
              <TbBook />
            </div>
            <h1 className="text-xl font-semibold text-foreground">
              Smart Learning
            </h1>
          </div>

          {/* Navigation Menu */}
          <nav className="hidden md:flex items-center space-x-1">
            {navigationItems.map((item) => {
              const Icon = item.icon;
              const isActive = router.state.location.pathname === item.to;

              return (
                <Button
                  key={item.to}
                  variant={isActive ? "default" : "ghost"}
                  size="sm"
                  onClick={() => handleNavigation(item.to)}
                  className="flex items-center space-x-2"
                >
                  <Icon className="w-4 h-4" />
                  <span>{item.label}</span>
                </Button>
              );
            })}
          </nav>

          {/* User Menu */}
          <div className="flex items-center space-x-4">
            <span className="text-sm text-muted-foreground hidden sm:block">
              歡迎回來，{user?.username ?? "使用者"}
            </span>
            <Button
              onClick={handleLogout}
              variant="outline"
              size="sm"
              disabled={logoutMutation.isPending}
            >
              {logoutMutation.isPending ? "登出中..." : "登出"}
            </Button>
          </div>
        </div>

        {/* Mobile Navigation */}
        <div className="md:hidden border-t pt-2 pb-2">
          <div className="flex flex-wrap gap-2">
            {navigationItems.map((item) => {
              const Icon = item.icon;
              const isActive = router.state.location.pathname === item.to;

              return (
                <Button
                  key={item.to}
                  variant={isActive ? "default" : "ghost"}
                  size="sm"
                  onClick={() => handleNavigation(item.to)}
                  className="flex items-center space-x-1 text-xs"
                >
                  <Icon className="w-3 h-3" />
                  <span>{item.label}</span>
                </Button>
              );
            })}
          </div>
        </div>
      </div>
    </header>
  );
};
