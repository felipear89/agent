import './UserList.css';
import { User } from '@/types/user';

interface UserListProps {
  users: User[];
  onAddUser: () => void;
  onDeleteUser: (id: number) => void;
}

export default function UserList({ users, onAddUser, onDeleteUser }: UserListProps) {
  return (
    <div className="users-list">
      <h2 className="users-list-title">List of users</h2>
      <ul className="users-list-ul">
        {users.map((user) => (
          <li className="users-list-item" key={user.id}>
            <span className="user-name">{user.name}</span>
            <button 
              className="delete-user-button"
              onClick={() => onDeleteUser(user.id)}
              title="Delete user"
            >
              ×
            </button>
          </li>
        ))}
      </ul>
      
      <button className="add-user-button" onClick={onAddUser}>
        Add User
      </button>
    </div>
  );
} 