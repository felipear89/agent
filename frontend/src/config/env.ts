export const config = {
  apiUrl: import.meta.env.VITE_API_URL || 'http://localhost:3000',
  environment: import.meta.env.MODE,
  isDevelopment: import.meta.env.DEV,
  isProduction: import.meta.env.PROD,
} as const;

export type Config = typeof config; 