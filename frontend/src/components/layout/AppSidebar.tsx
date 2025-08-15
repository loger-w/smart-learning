import { useRouter } from "@tanstack/react-router";
import { toast } from "sonner";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { UserMenu } from "./UserMenu";
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

export const AppSidebar = () => {
  const router = useRouter();

  const handleNavigation = (to: string) => {
    if (to !== "/dashboard") {
      toast.info(`${to} 功能即將推出`);
      return;
    }
    router.navigate({ to });
  };

  return (
    <Sidebar variant="inset" collapsible="icon">
      <SidebarHeader>
        <div className="flex items-center gap-2 px-4 py-2">
          <div className="w-8 h-8 bg-blue-600 text-white rounded-lg flex items-center justify-center">
            <TbBook />
          </div>
          <h1 className="text-lg font-semibold text-sidebar-foreground group-data-[collapsible=icon]:hidden">
            Smart Learning
          </h1>
        </div>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {navigationItems.map((item) => {
                const Icon = item.icon;
                const isActive = router.state.location.pathname === item.to;

                return (
                  <SidebarMenuItem key={item.to}>
                    <SidebarMenuButton
                      onClick={() => handleNavigation(item.to)}
                      isActive={isActive}
                      tooltip={item.label}
                    >
                      <Icon />
                      <span>{item.label}</span>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <UserMenu />
      </SidebarFooter>
    </Sidebar>
  );
};