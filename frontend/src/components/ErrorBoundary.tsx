import React from 'react'

interface Props {
  children: React.ReactNode
  fallback?: React.ComponentType<{ error: Error }>
}

export const ErrorBoundary: React.FC<Props> = ({ children, fallback: Fallback }) => {
  const [error, setError] = React.useState<Error | null>(null)
  
  React.useEffect(() => {
    const handleError = (event: ErrorEvent) => {
      setError(new Error(event.message))
    }
    
    const handlePromiseRejection = (event: PromiseRejectionEvent) => {
      setError(new Error(event.reason))
    }
    
    window.addEventListener('error', handleError)
    window.addEventListener('unhandledrejection', handlePromiseRejection)
    
    return () => {
      window.removeEventListener('error', handleError)
      window.removeEventListener('unhandledrejection', handlePromiseRejection)
    }
  }, [])
  
  if (error) {
    if (Fallback) {
      return <Fallback error={error} />
    }
    
    return (
      <div className="p-4 border border-red-300 rounded-md bg-red-50">
        <h2 className="text-lg font-semibold text-red-800">發生錯誤</h2>
        <p className="mt-2 text-red-600">{error.message}</p>
      </div>
    )
  }
  
  return <>{children}</>
}