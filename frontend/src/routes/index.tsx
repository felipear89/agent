import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuthContext } from '@/contexts/AuthContext';
import LoginPage from '@/pages/LoginPage';
import WelcomePage from '@/pages/WelcomePage';
import ChatPage from '@/pages/ChatPage';
import LoadingSpinner from '@/components/LoadingSpinner';
import DemoPage from '@/pages/DemoPage';

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
          isAuthenticated ? <Navigate to="/demo" replace /> : <LoginPage />
        }
      />
      <Route
        path="/welcome"
        element={
          isAuthenticated ? <WelcomePage /> : <Navigate to="/login" replace />
        }
      />
      <Route
        path="/chat"
        element={
          isAuthenticated ? <ChatPage /> : <Navigate to="/login" replace />
        }
      />
      <Route
        path="/demo"
        element={
          isAuthenticated ? <DemoPage /> : <Navigate to="/login" replace />
        }
      />
      {/* fallback */}
      <Route
        path="*"
        element={<Navigate to={isAuthenticated ? '/demo' : '/login'} replace />}
      />
    </Routes>
  );
}
