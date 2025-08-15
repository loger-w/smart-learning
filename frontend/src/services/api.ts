import axios, { type AxiosInstance, AxiosError } from "axios";

const getStoredToken = (): string | null => {
  const authStore = localStorage.getItem("auth-store");
  if (authStore) {
    try {
      const parsed = JSON.parse(authStore);
      return parsed.state?.token || null;
    } catch {
      return null;
    }
  }
  return null;
};

const handleUnauthorized = (): void => {
  localStorage.removeItem("auth-store");
};

const handleResponseError = (error: AxiosError) => {
  if (error.response?.status === 401) {
    handleUnauthorized();
  }
  return Promise.reject(error);
};

const setupInterceptors = (client: AxiosInstance): void => {
  client.interceptors.request.use((config) => {
    const token = getStoredToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  });

  client.interceptors.response.use(
    (response) => response,
    (error: AxiosError) => handleResponseError(error)
  );
};

const createAPIClient = (baseURL: string) => {
  const client = axios.create({
    baseURL,
    headers: { "Content-Type": "application/json" },
  });

  setupInterceptors(client);

  return {
    get: <T>(url: string, config = {}) => client.get<T>(url, config),

    post: <T>(url: string, data = {}, config = {}) =>
      client.post<T>(url, data, config),

    put: <T>(url: string, data = {}, config = {}) =>
      client.put<T>(url, data, config),

    delete: <T>(url: string, config = {}) => client.delete<T>(url, config),
  };
};

export const apiClient = createAPIClient(
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/api/v1"
);
