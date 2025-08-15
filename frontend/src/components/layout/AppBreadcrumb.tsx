import { useRouter } from "@tanstack/react-router";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { SidebarTrigger } from "@/components/ui/sidebar";
import { Separator } from "@/components/ui/separator";

const routeLabels: Record<string, string> = {
  "/dashboard": "首頁",
  "/flashcards": "記憶卡片",
  "/vocabulary": "單字練習",
  "/articles": "文章練習",
  "/statistics": "學習統計",
};

export const AppBreadcrumb = () => {
  const router = useRouter();
  const pathname = router.state.location.pathname;
  
  const currentLabel = routeLabels[pathname] || "頁面";

  return (
    <header className="flex h-16 shrink-0 items-center gap-2">
      <div className="flex items-center gap-2 px-4">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb>
          <BreadcrumbList>
            <BreadcrumbItem className="hidden md:block">
              <BreadcrumbLink href="/dashboard">Smart Learning</BreadcrumbLink>
            </BreadcrumbItem>
            {pathname !== "/dashboard" && (
              <>
                <BreadcrumbSeparator className="hidden md:block" />
                <BreadcrumbItem>
                  <BreadcrumbPage>{currentLabel}</BreadcrumbPage>
                </BreadcrumbItem>
              </>
            )}
          </BreadcrumbList>
        </Breadcrumb>
      </div>
    </header>
  );
};