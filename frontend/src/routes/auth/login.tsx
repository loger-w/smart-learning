import { createFileRoute, redirect } from "@tanstack/react-router";
import { LoginPage } from "@/features/auth/pages/LoginPage";
import { useAuthStore } from "@/stores/authStore";

export const Route = createFileRoute("/auth/login")({
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState();
    if (isAuthenticated) {
      throw redirect({ to: "/dashboard" });
    }
  },
  component: LoginPage,
});