import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { useAuthContext } from '@/contexts/AuthContext';
import { useEffect } from 'react';

interface LoginFormData {
  username: string;
  password: string;
}

export default function LoginPage() {
  const navigate = useNavigate();
  const { login, error, clearError, isAuthenticated } = useAuthContext();
  
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginFormData>();

  // Redirect if already authenticated
  useEffect(() => {
    if (isAuthenticated) {
      navigate('/welcome');
    }
  }, [isAuthenticated, navigate]);

  const onSubmit = async (data: LoginFormData) => {
    // Clear any previous errors when submitting
    clearError();
    
    try {
      await login(data);
      // Navigation will be handled by the useEffect above
    } catch {
      // Error is already handled in the hook
    }
  };

  return (
    <div className="page-container">
      <form onSubmit={handleSubmit(onSubmit)} className="login-form">
        <h1>Login</h1>
        
        {error && (
          <div className="error-message" style={{ textAlign: 'center' }}>
            {error}
          </div>
        )}
        
        <div className="form-group">
          <input
            type="text"
            placeholder="Username"
            {...register('username', { required: 'Username is required' })}
            className={errors.username ? 'error' : ''}
          />
          {errors.username && (
            <span className="error-message">{errors.username.message}</span>
          )}
        </div>
        <div className="form-group">
          <input
            type="password"
            placeholder="Password"
            {...register('password', { 
              required: 'Password is required',
              minLength: {
                value: 6,
                message: 'Password must be at least 6 characters',
              },
            })}
            className={errors.password ? 'error' : ''}
          />
          {errors.password && (
            <span className="error-message">{errors.password.message}</span>
          )}
        </div>
        <button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Logging in...' : 'Log In'}
        </button>
      </form>
    </div>
  );
}
