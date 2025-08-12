import React from 'react'
import { render as rtlRender, type RenderOptions } from '@testing-library/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const createTestQueryClient = (): QueryClient => new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
      staleTime: Infinity,
      gcTime: Infinity,
    },
    mutations: {
      retry: false,
    },
  },
})

interface CustomRenderOptions extends Omit<RenderOptions, 'wrapper'> {
  queryClient?: QueryClient
}

interface WrapperProps {
  children: React.ReactNode
}

const customRender = (
  ui: React.ReactElement,
  options: CustomRenderOptions = {}
): ReturnType<typeof rtlRender> => {
  const { queryClient = createTestQueryClient(), ...renderOptions } = options

  const Wrapper: React.FC<WrapperProps> = ({ children }) => (
    <QueryClientProvider client={queryClient}>
      {children}
    </QueryClientProvider>
  )

  return rtlRender(ui, { wrapper: Wrapper, ...renderOptions })
}

// 重新導出常用的測試工具函數
export {
  screen,
  fireEvent,
  waitFor,
  waitForElementToBeRemoved,
  within,
} from '@testing-library/react'

// 導出自定義的測試工具
export { customRender as render, createTestQueryClient }