import { useAuthContext } from '@/contexts/AuthContext';
import './WelcomePage.css';
import { useState } from 'react';

interface User {
  id: number;
  name: string;
}

export default function WelcomePage() {
  const { logout } = useAuthContext();

  const [users, setUsers] = useState<User[]>([]);

  console.log("carregando")
  let usersList = users.map((u) => (
    <li className="users-list-item" key={u.id}>
      {u.name}
    </li>
  ));

  const addUser = () => {
      setUsers([...users, { id: users.length + 1, name: `New User ${users.length}` }]);
  }

  return (
    <div className="page-container">
      <div className="welcome-container">
        <div className="welcome-header">
          <h1>Welcome!</h1>
          <p>You have successfully logged in.</p>
        </div>

        <div className="users-list">
          <h2 className="users-list-title">List of users</h2>
          <ul className="users-list-ul">
            {usersList}
          </ul>
        </div>

        <button className="add-user-button" onClick={addUser}>
          Add User
        </button>

        <button className="logout-button" onClick={logout}>
          Log Out
        </button>
      </div>
    </div>
  );
}
