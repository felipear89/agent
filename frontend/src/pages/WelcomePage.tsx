import { useAuthContext } from '@/contexts/AuthContext';

export default function WelcomePage() {
  const { logout } = useAuthContext();

  return (
    <div className="page-container">
      <h1>Welcome!</h1>
      <p>You have successfully logged in.</p>
      <button onClick={logout} style={{ marginTop: '1rem' }}>
        Log Out
      </button>
    </div>
  );
}
