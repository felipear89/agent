import { useAuthContext } from '@/contexts/AuthContext';
import './WelcomePage.css';

export default function DemoPage() {
  const { logout } = useAuthContext();

  return (
    <div className="page-container">
      <div className="welcome-container">
        
        <button className="logout-button" onClick={logout}>
          Log Out
        </button>
      </div>
    </div>
  );
}
