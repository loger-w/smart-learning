import {
  createRootRoute,
  createRoute,
  Outlet,
  redirect,
} from "@tanstack/react-router";
import { LoginPage } from "@/features/auth/pages/LoginPage";
import { RegisterPage } from "@/features/auth/pages/RegisterPage";
import { DashboardPage } from "@/features/dashboard/pages/DashboardPage";
import { useAuthStore } from "@/stores/authStore";

// 驗證檢查函數
const checkAuth = () => {
  const { isAuthenticated } = useAuthStore.getState();
  return isAuthenticated;
};

// Root route
const rootRoute = createRootRoute({
  component: () => (
    <div id="app">
      <Outlet />
    </div>
  ),
});

// Root redirect route
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  beforeLoad: () => {
    const isAuthenticated = checkAuth();
    if (!isAuthenticated) {
      throw redirect({ to: "/auth/login" });
    }
    throw redirect({ to: "/dashboard" });
  },
  component: () => null,
});

// Auth layout route
const authRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/auth",
  component: () => <Outlet />,
});

// Login route
const loginRoute = createRoute({
  getParentRoute: () => authRoute,
  path: "/login",
  beforeLoad: () => {
    const isAuthenticated = checkAuth();
    if (isAuthenticated) {
      throw redirect({ to: "/dashboard" });
    }
  },
  component: LoginPage,
});

// Register route
const registerRoute = createRoute({
  getParentRoute: () => authRoute,
  path: "/register",
  beforeLoad: () => {
    const isAuthenticated = checkAuth();
    if (isAuthenticated) {
      throw redirect({ to: "/dashboard" });
    }
  },
  component: RegisterPage,
});

// Dashboard route
const dashboardRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/dashboard",
  beforeLoad: () => {
    const isAuthenticated = checkAuth();
    if (!isAuthenticated) {
      throw redirect({ to: "/login" });
    }
  },
  component: DashboardPage,
});

// Catch-all route for unmatched paths
const catchAllRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "*",
  beforeLoad: () => {
    const isAuthenticated = checkAuth();
    if (!isAuthenticated) {
      throw redirect({ to: "/auth/login" });
    }
    throw redirect({ to: "/dashboard" });
  },
  component: () => null,
});

// Route tree
export const routeTree = rootRoute.addChildren([
  indexRoute,
  authRoute.addChildren([
    loginRoute,
    registerRoute,
  ]),
  dashboardRoute,
  catchAllRoute,
]);
