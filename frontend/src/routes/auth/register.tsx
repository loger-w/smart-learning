import { createFileRoute, redirect } from "@tanstack/react-router";
import { RegisterPage } from "@/features/auth/pages/RegisterPage";
import { useAuthStore } from "@/stores/authStore";

export const Route = createFileRoute("/auth/register")({
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState();
    if (isAuthenticated) {
      throw redirect({ to: "/dashboard" });
    }
  },
  component: RegisterPage,
});