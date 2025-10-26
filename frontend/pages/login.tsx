import { FormEvent, useState, useEffect } from 'react';
import { login } from '../lib/api';
import Header from '../components/Header';
import { useRouter } from 'next/router';
import { useAuth } from '../contexts/AuthContext';

export default function LoginPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [err, setErr] = useState('');
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  const { user, refresh } = useAuth();

  // If already logged in, bounce to "next" or home
  useEffect(() => {
    if (user) {
      const next = (router.query.next as string) || '/';
      router.replace(next);
    }
  }, [user, router]);

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setErr('');
    setLoading(true);
    try {
      await login(username, password);
      await refresh(); // refresh context
      const next = (router.query.next as string) || '/';
      router.replace(next);
    } catch (e: any) {
      setErr(e?.message || 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Header user={user} />
      <main style={{ padding: 24 }}>
        <h1>Login</h1>
        <form onSubmit={onSubmit} style={{ display: 'grid', gap: 12, maxWidth: 360 }}>
          <label>
            <div>Username</div>
            <input value={username} onChange={(e) => setUsername(e.target.value)} />
          </label>
          <label>
            <div>Password</div>
            <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" />
          </label>
          <button type="submit" disabled={loading}>
            {loading ? 'Signing in...' : 'Sign in'}
          </button>
          {err && <p style={{ color: 'crimson' }}>{err}</p>}
        </form>
      </main>
    </>
  );
}
