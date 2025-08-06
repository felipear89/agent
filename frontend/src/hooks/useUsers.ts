import { useState, useRef, useCallback, useMemo } from 'react';
import { User } from '@/types/user';

export function useUsers() {
  const [users, setUsers] = useState<User[]>([]);
  const nextIdRef = useRef(1);

  const addUser = useCallback(() => {
    const newUser: User = {
      id: nextIdRef.current,
      name: `New User ${nextIdRef.current}`,
    };

    setUsers(prevUsers => [...prevUsers, newUser]);
    nextIdRef.current += 1;

    console.log('User added:', newUser);
    return newUser; // Return the created user
  }, []);

  const deleteUser = useCallback((id: number) => {
    setUsers(prevUsers => prevUsers.filter(user => user.id !== id));
  }, []);

  const updateUser = useCallback((id: number, updates: Partial<User>) => {
    setUsers(prevUsers =>
      prevUsers.map(user => (user.id === id ? { ...user, ...updates } : user))
    );
  }, []);

  const clearUsers = useCallback(() => {
    setUsers([]);
    nextIdRef.current = 1; // Reset ID counter
  }, []);

  const findUser = useCallback(
    (id: number) => {
      return users.find(user => user.id === id);
    },
    [users]
  );

  const getUserCount = useMemo(() => users.length, [users]);

  const hasUsers = useMemo(() => users.length > 0, [users]);

  return {
    // State
    users,
    userCount: getUserCount,
    hasUsers,

    // Actions
    addUser,
    deleteUser,
    updateUser,
    clearUsers,
    findUser,
  };
}
