import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuthContext } from '@/contexts/AuthContext';
import LoginPage from '@/pages/LoginPage';
import WelcomePage from '@/pages/WelcomePage';
import ChatPage from '@/pages/ChatPage';
import LoadingSpinner from '@/components/LoadingSpinner';
import DemoPage from '@/pages/DemoPage';

export default function AppRoutes() {
  const { isAuthenticated, isLoading } = useAuthContext();

  const hideChatbot = () => {
    const script = document.querySelector<HTMLIFrameElement>('iframe[src*="agents"]');
    if (script) {
      script.style.display = 'none';
    }
  };

  // Show loading state while checking authentication
  if (isLoading) {
    return (
      <div className="page-container">
        <LoadingSpinner size="large" message="Checking authentication..." />
      </div>
    );
  }

  if (!isAuthenticated) {
    hideChatbot();
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
