import { createRootRoute, createRoute, Outlet, redirect } from '@tanstack/react-router'
import { LoginPage } from '@/features/auth/pages/LoginPage'
import { DashboardPage } from '@/features/dashboard/pages/DashboardPage'
import { useAuthStore } from '@/stores/authStore'

// Root route
const rootRoute = createRootRoute({
  component: () => (
    <div id="app">
      <Outlet />
    </div>
  ),
})

// Root redirect route
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState()
    if (!isAuthenticated) {
      throw redirect({ to: '/login' })
    }
    throw redirect({ to: '/dashboard' })
  },
  component: () => null,
})

// Login route
const loginRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/login',
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState()
    if (isAuthenticated) {
      throw redirect({ to: '/dashboard' })
    }
  },
  component: LoginPage,
})

// Dashboard route
const dashboardRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/dashboard',
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState()
    if (!isAuthenticated) {
      throw redirect({ to: '/login' })
    }
  },
  component: DashboardPage,
})

// Catch-all route for unmatched paths
const catchAllRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/$',
  beforeLoad: () => {
    const { isAuthenticated } = useAuthStore.getState()
    if (!isAuthenticated) {
      throw redirect({ to: '/login' })
    }
    throw redirect({ to: '/dashboard' })
  },
  component: () => null,
})

// Route tree
export const routeTree = rootRoute.addChildren([
  indexRoute, 
  loginRoute, 
  dashboardRoute, 
  catchAllRoute
])