import { useState, useCallback, useEffect } from 'react';
import { getAuth, setAuth, clearAuth } from '@/lib/storage';

interface AuthState {
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

interface AuthActions {
  login: (credentials: { username: string; password: string }) => Promise<void>;
  logout: () => void;
  clearError: () => void;
}

export default function useAuth(): AuthState & AuthActions {
  const [state, setState] = useState<AuthState>({
    isAuthenticated: false,
    isLoading: true,
    error: null,
  });

  // Initialize auth state on mount
  useEffect(() => {
    const isAuth = getAuth();
    setState(prev => ({
      ...prev,
      isAuthenticated: isAuth,
      isLoading: false,
    }));
  }, []);

  const login = useCallback(
    async (credentials: { username: string; password: string }) => {
      setState(prev => ({ ...prev, isLoading: true, error: null }));

      try {
        // TODO: Replace with actual API call
        // const response = await api.login(credentials);

        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1000));

        // For demo purposes, accept any non-empty credentials
        if (credentials.username === 'a') {
          throw new Error('Invalid credentials');
        }

        setAuth(true);
        setState(prev => ({
          ...prev,
          isAuthenticated: true,
          isLoading: false,
          error: null,
        }));
      } catch (error) {
        setState(prev => ({
          ...prev,
          isAuthenticated: false,
          isLoading: false,
          error: error instanceof Error ? error.message : 'Login failed',
        }));
      }
    },
    []
  );

  const logout = useCallback(() => {
    clearAuth();
    setState(prev => ({
      ...prev,
      isAuthenticated: false,
      error: null,
    }));
  }, []);

  const clearError = useCallback(() => {
    setState(prev => ({ ...prev, error: null }));
  }, []);

  return {
    ...state,
    login,
    logout,
    clearError,
  };
}
