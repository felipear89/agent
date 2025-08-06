import { useAuthContext } from '@/contexts/AuthContext';
import { Link } from 'react-router-dom';
import './WelcomePage.css';
import UserList from '@/components/UserList';
import { useUsers } from '@/hooks/useUsers';

export default function WelcomePage() {
  const { logout } = useAuthContext();
  const { users, addUser, deleteUser } = useUsers();

  return (
    <div className="page-container">
      <div className="welcome-container">
        <div className="welcome-header">
          <h1>Welcome!</h1>
          <p>You have successfully logged in.</p>
        </div>

        <UserList users={users} onAddUser={addUser} onDeleteUser={deleteUser} />

        <div className="navigation-buttons">
          <Link to="/chat" className="nav-button chat-button">
            💬 Go to Chat
          </Link>

          <button className="logout-button" onClick={logout}>
            Log Out
          </button>
        </div>
      </div>
    </div>
  );
}
