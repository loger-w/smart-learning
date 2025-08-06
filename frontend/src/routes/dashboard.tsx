import { createFileRoute, redirect } from "@tanstack/react-router";
import { DashboardPage } from "@/features/dashboard/pages/DashboardPage";
import { useAuthStore } from "@/stores/authStore";

export const Route = createFileRoute("/dashboard")({
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState();
    if (!isAuthenticated) {
      throw redirect({ to: "/auth/login" });
    }
  },
  component: DashboardPage,
});