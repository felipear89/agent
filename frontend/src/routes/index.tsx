import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuthContext } from '@/contexts/AuthContext';
import LoginPage from '@/pages/LoginPage';
import WelcomePage from '@/pages/WelcomePage';
import LoadingSpinner from '@/components/LoadingSpinner';

export default function AppRoutes() {
  const { isAuthenticated, isLoading } = useAuthContext();

  // Show loading state while checking authentication
  if (isLoading) {
    return (
      <div className="page-container">
        <LoadingSpinner size="large" message="Checking authentication..." />
      </div>
    );
  }

  return (
    <Routes>
      <Route
        path="/login"
        element={
          isAuthenticated ? <Navigate to="/welcome" replace /> : <LoginPage />
        }
      />
      <Route
        path="/welcome"
        element={
          isAuthenticated ? <WelcomePage /> : <Navigate to="/login" replace />
        }
      />
      {/* fallback */}
      <Route
        path="*"
        element={
          <Navigate to={isAuthenticated ? '/welcome' : '/login'} replace />
        }
      />
    </Routes>
  );
}
