const AUTH_KEY = 'auth';

export const setAuth = (value: boolean) => {
  localStorage.setItem(AUTH_KEY, String(value));
};

export const getAuth = (): boolean => localStorage.getItem(AUTH_KEY) === 'true';

export const clearAuth = () => {
  localStorage.removeItem(AUTH_KEY);
};
